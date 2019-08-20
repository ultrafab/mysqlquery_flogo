package mysqlquery_flogo

import (
	"io/ioutil"
	"testing"
	"encoding/json"
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

	m := make(map[string]interface{})
	m["pressione_perdita"] = 51.6
	m["pressione_ingresso"] = 5.1
	data, _ := json.Marshal(m)
	tc.SetInput("message", string(data))
	tc.SetInput("username", "ukeau")
	tc.SetInput("password", "kaplan")
	tc.SetInput("query", "INSERT INTO data (date, $MESSAGE_KEYS) VALUES (UTC_TIMESTAMP(),$MESSAGE_VALUES)")
	tc.SetInput("db", "rub")
	act.Eval(tc)

	value := tc.GetOutput("result")
	fmt.Println(value)
}
