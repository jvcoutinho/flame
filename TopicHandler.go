package flame

import "./message"

type TopicHandler struct {
	topicName string
	channel   *Channel
}

func newTopicHandler(name string, channel *Channel) *TopicHandler {
	return &TopicHandler{name, channel}
}

func (topicHandler *TopicHandler) Publish(item interface{}) error {
	request := message.NewRequest("", message.Publish, topicHandler.topicName, item)
	return topicHandler.channel.call(request)
}

func (topicHandler *TopicHandler) Subscribe() (<-chan interface{}, error) {
	request := message.NewRequest("", message.Subscribe, topicHandler.topicName, "host")
	if err := topicHandler.channel.call(request); err != nil {
		return nil, err
	}

	return make(chan interface{}), nil
}
