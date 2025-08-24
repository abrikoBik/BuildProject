package l10n

import "errors"

const (
	GenericAttachment  = "Ошибка чертежа"
	AttachmentSize     = "Файл слишком большой"
	AttachmentNotFound = "Файл не найден"
)

var (
	GenericAttachmentError  = errors.New(GenericAttachment)
	AttachmentSizeError     = errors.New(AttachmentSize)
	AttachmentNotFoundError = errors.New(AttachmentNotFound)
)
