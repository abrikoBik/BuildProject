package middleware

type ContextKey int

const (
	JWTKey ContextKey = iota

	UserID
	WorkerID
	ProjectID

	StageGroupID
	EstimateGroupID

	StageID
	EstimateID

	// Type: uuid.UUID
	AttachmentUUID
)
