package model

type FlashType string

const (
	Success FlashType = "success"
	Info    FlashType = "info"
	Error   FlashType = "error"
)

type Flash struct {
	Message string
	Type    FlashType
}

func (flash *Flash) GetType() string {
	return string(flash.Type)
}

func NewFlash(message string, flashType FlashType) Flash {
	return Flash{
		Message: message,
		Type:    flashType,
	}
}
