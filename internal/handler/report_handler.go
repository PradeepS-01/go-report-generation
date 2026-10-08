package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PradeepS-01/go-report-generation/internal/domain"
	"github.com/PradeepS-01/go-report-generation/internal/generator"
	"github.com/PradeepS-01/go-report-generation/internal/service"
)

type ReportHandler struct {
	reportService *service.ReportService
	pdfGenerator  *generator.PDFGenerator
}

func NewReportHandler(
	reportService *service.ReportService,
	pdfGenerator *generator.PDFGenerator,
) *ReportHandler {

	return &ReportHandler{
		reportService: reportService,
		pdfGenerator:  pdfGenerator,
	}
}

func (h *ReportHandler) GenerateExamReport(
	c *gin.Context,
) {

	var request domain.ExamReportRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if err := service.ValidateReportRequest(request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	result := h.reportService.CalculateReport(
		request.Exam,
	)

	pdf, err := h.pdfGenerator.Generate(
		request,
		result,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate report",
		})

		return
	}

	c.Data(
		http.StatusOK,
		"application/pdf",
		pdf,
	)
}