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

// IsValid reports whether t is one of the job types the workers know how to run
func (t JobType) IsValid() bool {
	switch t {
	case JobTypeDocumentAnalysis, JobTypeSendEmail, JobTypeGenerateReport:
		return true
	}
	return false
}
