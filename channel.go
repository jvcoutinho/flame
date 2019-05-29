package flame

import (
	"errors"
	"os"

	"github.com/golang-collections/go-datastructures/queue"

	"./handler"
	"./marshaller"
	"./message"
	"./stream"
)

type Channel struct {
	userID         string
	requestHandler *handler.RequestHandler
	itemHandler    *handler.RequestHandler
	topicReceivers map[string]*chan interface{}
}

func (channel *Channel) InitializeTopic(topicName string) (*TopicHandler, error) {
	// Cria um request, envia e checa se houve erro.
	request := message.NewRequest(channel.userID, message.Initialize, topicName, nil, 0) 
	if err := channel.call(request); err != nil {
		return nil, err
	}

	return newTopicHandler(topicName, channel), nil
}

func (channel *Channel) AccessTopic(topicName string) (*TopicHandler, error) {
	// Cria um request, envia e checa se houve erro.
	request := message.NewRequest("", message.CheckExistence, topicName, nil, 0)
	if err := channel.call(request); err != nil {
		return nil, err
	}

	return newTopicHandler(topicName, channel), nil
}

func (channel *Channel) call(request *message.Request) error {
	marshaller := &marshaller.JSONMarshaller{}

	// Serializa o request e envia.
	marshalledRequest := marshaller.MarshalRequest(*request)
	channel.requestHandler.Send(marshalledRequest)

	// Espera por uma resposta, a serializa e retorna o erro do response (possivelmente nil).
	marshalledResponse, _ := channel.requestHandler.Receive()
	response := marshaller.UnmarshalResponse(marshalledResponse)
	if response.HasError() {
		return response.GetError()
	}
	return nil
}

func (channel *Channel) receiveItem() {
	marshaller := &marshaller.JSONMarshaller{}
	for {
		marshalledRequest, err := channel.itemHandler.Receive()
		if err != nil {
			// disconnection
			break
		}
		request, err := marshaller.UnmarshalRequest(marshalledRequest)
		if err != nil {
			continue
		}
		topicName := request.GetQueueName()
		body := request.GetBody()
		*channel.topicReceivers[topicName] <- body
	}
}

func (channel *Channel) addChannel(topicName string, topicChannel *chan interface{}) {
	channel.topicReceivers[topicName] = topicChannel
}

/*
 * SPECIFIC
 */

func (channel *Channel) Stream(name string, filePath string) (*stream.StreamHandler, error) {
	file, stats, err := channel.openFile(filePath)
	if err != nil {
		return nil, errors.New("Could not open file.")
	}

	request := message.NewRequest(channel.userID, message.Stream, name, stats.Size(), 0)
	if err := channel.call(request); err != nil {
		return nil, err
	}

	err = channel.transmitFile(name, file, stats)
	file.Close()
	return &stream.StreamHandler{}, err // TODO: fazer o stream handler.
}

func (channel *Channel) Play(name string) (*queue.RingBuffer, error) {
	request := message.NewRequest(channel.userID, message.Play, name, nil, 0)
	if err := channel.call(request); err != nil {
		return nil, err
	}
}

func (channel *Channel) openFile(path string) (*os.File, os.FileInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	stats, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}
	return file, stats, nil
}

func (channel *Channel) transmitFile(streamName string, file *os.File, stats os.FileInfo) error {
	fileSize := stats.Size()
	buffer := make([]byte, 1024)

	var sentBytes int64
	for sentBytes < fileSize {

		read, err := file.Read(buffer)
		if err != nil {
			return errors.New("Read unsuccessful") // TODO: destroy queue.
		}

		request := message.NewRequest(channel.userID, message.Publish, streamName, buffer[:read], 1)
		if err := channel.call(request); err != nil {
			return err
		}

		sentBytes = sentBytes + int64(read)
	}

	return nil
}
