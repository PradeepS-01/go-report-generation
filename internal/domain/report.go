package domain

type ExamReportRequest struct {
	Student Student `json:"student"`
	Exam    Exam    `json:"exam"`
}

type Student struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RollNumber string `json:"rollNumber"`
	Class      string `json:"class"`
}

type Exam struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Year     int       `json:"year"`
	Subjects []Subject `json:"subjects"`
}

type Subject struct {
	Name     string  `json:"name"`
	Marks    float64 `json:"marks"`
	MaxMarks float64 `json:"maxMarks"`
}

type ReportResult struct {
	TotalMarks float64
	MaxMarks   float64
	Percentage float64
	Result     string
}