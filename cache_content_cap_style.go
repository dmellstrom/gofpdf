package gofpdf

import (
	"fmt"
	"io"
)

type cacheContentCapStyle struct {
	style int
}

func (c *cacheContentCapStyle) write(w io.Writer, protection *PDFProtection) error {
	fmt.Fprintf(w, "%d J\n", c.style)
	return nil
}
