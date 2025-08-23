package util

const (
	MinPasswordLength = 8

	// Should be around 128MB
	MaxAttachmentSize = 1 << 27
	StoragePath       = "/var/lib/attachments"
	FormKey           = "attachment"
)
