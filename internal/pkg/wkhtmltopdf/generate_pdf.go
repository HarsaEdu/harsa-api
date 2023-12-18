package pdf

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// PDFGenerator represents a PDF generator.
type PDFGenerator interface {
	GenerateCertificate(certificate *domain.Certificate) ([]byte, error)
}

// wkhtmltopdfGenerator implements the PDFGenerator interface using wkhtmltopdf.
type wkhtmltopdfGenerator struct {
	htmlTemplate embed.FS
}

// NewWKHTMLToPDFGenerator creates a new instance of wkhtmltopdfGenerator.
func NewWKHTMLToPDFGenerator(htmlTemplate embed.FS) PDFGenerator {
	return &wkhtmltopdfGenerator{htmlTemplate: htmlTemplate}
}

// GenerateCertificate implements the PDFGenerator interface for wkhtmltopdf.
func (g *wkhtmltopdfGenerator) GenerateCertificate(certificate *domain.Certificate) ([]byte, error) {
	var htmlBuffer bytes.Buffer
	certificateTemplate, err := template.ParseFS(g.htmlTemplate, "certificate.html")
	if err != nil {
		return nil, err
	}

	err = certificateTemplate.Execute(&htmlBuffer, certificate)
	if err != nil {
		return nil, err
	}

	// Initialize the converter
	pdfGen, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlBuffer.Bytes()))
	pdfGen.AddPage(page)
	pdfGen.MarginBottom.Set(0)
	pdfGen.MarginLeft.Set(0)
	pdfGen.MarginRight.Set(0)
	pdfGen.MarginTop.Set(0)
	pdfGen.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfGen.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// Generate the PDF
	err = pdfGen.Create()
	if err != nil {
		return nil, err
	}
	return pdfGen.Bytes(), nil
}
