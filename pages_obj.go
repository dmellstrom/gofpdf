package gofpdf

import (
	"fmt"
	"io"
)

const pagesType = "Pages"

//PagesObj pdf pages object
type PagesObj struct { //impl IObj
	PageCount int
	Kids      string
	getRoot   func() *Fpdf
}

func (p *PagesObj) init(funcGetRoot func() *Fpdf) {
	p.PageCount = 0
	p.getRoot = funcGetRoot
}

func (p *PagesObj) write(w io.Writer, objID int) error {

	io.WriteString(w, "<<\n")
	fmt.Fprintf(w, "  /Type /%s\n", p.getType())
	p.getRoot().curr.pageOption.writePageBoundaries(w)
	fmt.Fprintf(w, "  /Count %d\n", p.PageCount)
	fmt.Fprintf(w, "  /Kids [ %s ]\n", p.Kids) //sample Kids [ 3 0 R ]
	io.WriteString(w, ">>\n")
	return nil
}

func (p *PagesObj) getType() string {
	return pagesType
}

func (p *PagesObj) test() {
	fmt.Print(p.getType() + "\n")
}

func (p *PagesObj) addPage(page int) {
	p.PageCount++
	p.Kids = fmt.Sprintf("%s %d 0 R ", p.Kids, page+1)
}
