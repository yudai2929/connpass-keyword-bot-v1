package message

type MessageRepository interface {
	SendMessage(messages []Message) error
}
