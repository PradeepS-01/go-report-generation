package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PradeepS-01/go-report-generation/internal/config"
	"github.com/PradeepS-01/go-report-generation/internal/generator"
	"github.com/PradeepS-01/go-report-generation/internal/handler"
	"github.com/PradeepS-01/go-report-generation/internal/service"
)

func main() {
	cfg := config.Load()

	reportService := service.NewReportService()
	pdfGenerator := generator.NewPDFGenerator()

	reportHandler := handler.NewReportHandler(
		reportService,
		pdfGenerator,
	)

	router := gin.Default()

	router.GET("/start", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "student-report-service",
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/reports/exam", reportHandler.GenerateExamReport)
	}

	log.Printf(
		"student report service running on :%s",
		cfg.Port,
	)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}