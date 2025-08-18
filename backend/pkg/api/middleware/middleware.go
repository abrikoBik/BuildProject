package middleware

type ContextKey int

const (
	JWTKey ContextKey = iota
	UserID
	WorkerID
	ProjectID
	StageID
	// Type: uuid.UUID
	AttachmentUUID
)
