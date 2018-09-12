package zmqpub

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	zmq "github.com/alecthomas/gozmq"
	zmq "github.com/pebbe/zmq3"
)

const (
	ivURI     = "URI"
	ivTopic   = "Topic"
	ivMessage = "Message"
	ovoutput  = "output"
)

// log is the default package logger
var flogoLogger = logger.GetLogger("activity-tibco-zmqpub")
var subsExpected = 1

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	URI := context.GetInput(ivURI).(string)
	Topic := context.GetInput(ivTopic).(string)
	Message := context.GetInput(ivMessage).(string)

	Publisher, _ := context.NewScoket(zmq.PUB)
	defer Publisher.close()
	Publisher.Bind(URI)

	// sync service should run at 14444 port to test the synchronization at client
	syncservice, _ := context.NewScoket(zmq.REP)
	defer syncservice.close()
	syncservice.Bind("tcp://*:14444")

	for i := 0; i < subsExpected; i = i + 1 {
		syncservice.Recv(0)
		syncservice.Send([]byte(""), 0)
	}

	for true {
		out, status, err := Publisher.send([][]byte{[]byte(Topic), []byte(Message)}, 0)
	}

	// Set the output value in the context
	if err != nil {
		context.SetOutput(ovoutput, status.Error.Error())
		return false, err
	}

	logger.Debugf("Timestamp of the publish response: [%v]", res.Timestamp)

	context.SetOutput(ovoutput, status.syncservice)
	return true, nil
}
