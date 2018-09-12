package zmqpub

import (
"io/ioutil"
"log"
"testing"


"github.com/TIBCOSoftware/flogo-lib/core/activity"
"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"fmt"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
	log.Println("Test create successful")
	}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
    tc.SetInput("URI","tcp:localhost:5555")
	tc.SetInput("Topic", "zmq.REP")
	tc.SetInput("Message", "testing zmq server for flogo")
	act.Eval(tc)
	log.Printf("TestEval successful output [%d]", tc.GetOutput("output") )

	result := tc.GetOutput("result")
	fmt.Printf("[%s]", result)
	//check result attr
}
