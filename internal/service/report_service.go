package service

import (
	"fmt"

	"github.com/PradeepS-01/go-report-generation/internal/domain"
)

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) CalculateReport(
	exam domain.Exam,
) domain.ReportResult {

	var totalMarks float64
	var maxMarks float64

	for _, subject := range exam.Subjects {
		totalMarks += subject.Marks
		maxMarks += subject.MaxMarks
	}

	var percentage float64

	if maxMarks > 0 {
		percentage = (totalMarks / maxMarks) * 100
	}

	result := "PASS"

	if percentage < 35 {
		result = "FAIL"
	}

	return domain.ReportResult{
		TotalMarks: totalMarks,
		MaxMarks:   maxMarks,
		Percentage: percentage,
		Result:     result,
	}
}

func ValidateReportRequest(
	request domain.ExamReportRequest,
) error {

	if request.Student.ID == "" {
		return fmt.Errorf("student id is required")
	}

	if request.Student.Name == "" {
		return fmt.Errorf("student name is required")
	}

	if request.Exam.ID == "" {
		return fmt.Errorf("exam id is required")
	}

	if len(request.Exam.Subjects) == 0 {
		return fmt.Errorf("at least one subject is required")
	}

	for _, subject := range request.Exam.Subjects {

		if subject.Name == "" {
			return fmt.Errorf("subject name is required")
		}

		if subject.MaxMarks <= 0 {
			return fmt.Errorf(
				"max marks must be greater than zero",
			)
		}

		if subject.Marks < 0 {
			return fmt.Errorf(
				"marks cannot be negative",
			)
		}

		if subject.Marks > subject.MaxMarks {
			return fmt.Errorf(
				"marks cannot exceed maximum marks",
			)
		}
	}

	return nil
}