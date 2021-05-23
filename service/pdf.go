package service

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePdf() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println("PDF生成失败")
	}

}
