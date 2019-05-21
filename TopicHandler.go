package flame

import "./message"

type TopicHandler struct {
	topicName string
	channel   *Channel
}

func newTopicHandler(name string, channel *Channel) *TopicHandler {
	return &TopicHandler{name, channel}
}

func (topicHandler *TopicHandler) Publish(item interface{}) {
	message := message.New("push", topicHandler.topicName, item)
	topicHandler.channel.send(message)
	// TODO: checar se tá tudo ok.
}
