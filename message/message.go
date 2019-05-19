package message

type Message struct {
	Operation string
	QueueName string
	Body      interface{}
}

func (msg *Message) GetOperation() string {
	return msg.Operation
}

func (msg *Message) GetQueueName() string {
	return msg.QueueName
}

func (msg *Message) GetBody() interface{} {
	return msg.Body
}

func New(operation string, queueName string, body interface{}) *Message {
	return &Message{operation, queueName, body}
}
