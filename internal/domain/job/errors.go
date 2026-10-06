package job

import "errors"

// Domain errors - infra adapters translate driver errors into these
var (
	ErrNotFound       = errors.New("job not found")
	ErrInvalidType    = errors.New("invalid job type")
	ErrInvalidPayload = errors.New("payload must be a JSON object")
)
