package flame

import (
	"./handler"
	"./marshaller"
	"./message"
)

type Channel struct {
	handler *handler.ClientRequestHandler
}

func (channel *Channel) InitializeTopic(topicName string) *TopicHandler {
	message := message.New("init", topicName, nil)
	channel.send(message)
	// TODO: receber resposta e checar se tá tudo ok.
	return newTopicHandler(topicName, channel)
}

func (channel *Channel) send(message *message.Message) {
	marshaller := marshaller.JSONMarshaller{}
	data := marshaller.Marshal(*message)
	channel.handler.Send(data)
}
