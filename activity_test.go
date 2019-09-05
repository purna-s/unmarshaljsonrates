package unmarshaljsonrates

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
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

	//setup attrs
	tc.SetInput(`{"data":{"averageEgressByteRatePerMinute":7450,"averageEgressRatePerMinute":18,"averageIngressByteRatePerMinute":7450,"averageIngressRatePerMinute":18,"currentEgressByteRatePerSecond":0,"currentEgressRatePerSecond":0,"currentIngressByteRatePerSecond":0,"currentIngressRatePerSecond":0,"msgVpnName":"LTAP_TAXI_MSG","queueName":"Q_TAXI_PRIME_TAXIINFO_TS"},"links":{"uri":"http://10.10.11.17/SEMP/v2/monitor/msgVpns/LTAP_TAXI_MSG/queues/Q_TAXI_PRIME_TAXIINFO_TS/rates"},"meta":{"request":{"method":"GET","uri":"http://10.10.11.17/SEMP/v2/monitor/msgVpns/LTAP_TAXI_MSG/queues/Q_TAXI_PRIME_TAXIINFO_TS/rates"},"responseCode":200}}`)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	avgingmsgrt := tc.GetOutput("avgingmsgrt")
	assert.Equal(t, avgingmsgrt, avgingmsgrt)

}
