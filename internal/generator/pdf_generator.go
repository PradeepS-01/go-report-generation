package generator

import (
	"bytes"
	"fmt"

	"github.com/go-pdf/fpdf"

	"github.com/PradeepS-01/go-report-generation/internal/domain"
)

type PDFGenerator struct{}

func NewPDFGenerator() *PDFGenerator {
	return &PDFGenerator{}
}

func (g *PDFGenerator) Generate(
	request domain.ExamReportRequest,
	result domain.ReportResult,
) ([]byte, error) {

	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	// Title
	pdf.SetFont("Arial", "B", 18)

	pdf.CellFormat(
		0,
		12,
		"STUDENT EXAMINATION REPORT",
		"",
		1,
		"C",
		false,
		0,
		"",
	)

	pdf.Ln(5)

	// Student details
	pdf.SetFont("Arial", "", 12)

	pdf.Cell(
		0,
		8,
		fmt.Sprintf("Student: %s", request.Student.Name),
	)

	pdf.Ln(7)

	pdf.Cell(
		0,
		8,
		fmt.Sprintf(
			"Roll Number: %s",
			request.Student.RollNumber,
		),
	)

	pdf.Ln(7)

	pdf.Cell(
		0,
		8,
		fmt.Sprintf(
			"Class: %s",
			request.Student.Class,
		),
	)

	pdf.Ln(10)

	// Table header
	pdf.SetFont("Arial", "B", 11)

	pdf.CellFormat(
		80,
		8,
		"Subject",
		"1",
		0,
		"C",
		false,
		0,
		"",
	)

	pdf.CellFormat(
		40,
		8,
		"Marks",
		"1",
		0,
		"C",
		false,
		0,
		"",
	)

	pdf.CellFormat(
		40,
		8,
		"Maximum",
		"1",
		1,
		"C",
		false,
		0,
		"",
	)

	// Subjects
	pdf.SetFont("Arial", "", 11)

	for _, subject := range request.Exam.Subjects {

		pdf.CellFormat(
			80,
			8,
			subject.Name,
			"1",
			0,
			"L",
			false,
			0,
			"",
		)

		pdf.CellFormat(
			40,
			8,
			fmt.Sprintf("%.0f", subject.Marks),
			"1",
			0,
			"C",
			false,
			0,
			"",
		)

		pdf.CellFormat(
			40,
			8,
			fmt.Sprintf("%.0f", subject.MaxMarks),
			"1",
			1,
			"C",
			false,
			0,
			"",
		)
	}

	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)

	pdf.Cell(
		0,
		8,
		fmt.Sprintf(
			"Total: %.0f / %.0f",
			result.TotalMarks,
			result.MaxMarks,
		),
	)

	pdf.Ln(8)

	pdf.Cell(
		0,
		8,
		fmt.Sprintf(
			"Percentage: %.2f%%",
			result.Percentage,
		),
	)

	pdf.Ln(8)

	pdf.Cell(
		0,
		8,
		fmt.Sprintf(
			"Result: %s",
			result.Result,
		),
	)

	var buffer bytes.Buffer

	if err := pdf.Output(&buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}