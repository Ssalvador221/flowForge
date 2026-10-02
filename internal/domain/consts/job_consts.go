package consts

type JobStatus string
type JobType string

const (
	JobStatusPending    JobStatus = "Pending"
	JobStatusProcessing JobStatus = "Processing"
	JobStatusCompleted  JobStatus = "Completed"
	JobStatusCancelled  JobStatus = "Cancelled"
	JobStatusFailed     JobStatus = "Failed"
)

const (
	JobTypeDocumentAnalysis JobType = "document_analysis"
	JobTypeSendEmail        JobType = "send_email"
	JobTypeGenerateReport   JobType = "generate_report"
)
