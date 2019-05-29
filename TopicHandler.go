package flame

import "./message"

type TopicHandler struct {
	topicName string
	channel   *Channel
}

func newTopicHandler(name string, channel *Channel) *TopicHandler {
	return &TopicHandler{name, channel}
}

func (topicHandler *TopicHandler) Publish(item interface{}, priority int) error {
	request := message.NewRequest(topicHandler.channel.userID, message.Publish, topicHandler.topicName, item, priority)
	return topicHandler.channel.call(request)
}

func (topicHandler *TopicHandler) Subscribe() (<-chan interface{}, error) {
	request := message.NewRequest(topicHandler.channel.userID, message.Subscribe, topicHandler.topicName, "", 0)
	if err := topicHandler.channel.call(request); err != nil && err.Error() != "You are already a subscriber of the topic '"+topicHandler.topicName+"'" {
		return nil, err
	}

	channel := make(chan interface{})
	topicHandler.channel.addChannel(topicHandler.topicName, &channel)
	return channel, nil
}
