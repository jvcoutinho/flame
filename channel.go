package flame

import (
	"./handler"
	"./marshaller"
	"./message"
)

type Channel struct {
	handler *handler.ClientRequestHandler
}

func (channel *Channel) InitializeTopic(topicName string) (*TopicHandler, error) {
	request := message.NewRequest("", message.Initialize, topicName, nil)
	if err := channel.call(request); err != nil {
		return nil, err
	}

	return newTopicHandler(topicName, channel), nil
}

func (channel *Channel) AccessTopic(topicName string) (*TopicHandler, error) {
	request := message.NewRequest("", message.CheckExistence, topicName, nil)
	if err := channel.call(request); err != nil {
		return nil, err
	}

	return newTopicHandler(topicName, channel), nil
}

func (channel *Channel) call(request *message.Request) error {
	marshaller := &marshaller.JSONMarshaller{}

	marshalledRequest := marshaller.MarshalRequest(*request)
	channel.handler.Send(marshalledRequest)

	response := marshaller.UnmarshalResponse(channel.handler.Receive())
	if response.HasError() {
		return response.GetError()
	}
	return nil
}
