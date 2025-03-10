package dpobjects

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type LogTarget struct {
	internalState map[string]interface{}
}

func NewLogTarget(state map[string]interface{}) LogTarget {
	return LogTarget{
		internalState: state,
	}
}

func ToPrettyJSON(logTarget LogTarget) string {
	json, err := json.MarshalIndent(logTarget.internalState, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(json[:])
}

func ToPrettyYAML(logTarget LogTarget) string {
	yaml, err := yaml.Marshal(logTarget.internalState)
	if err != nil {
		panic(err)
	}
	return string(yaml[:])
}

func Test(l LogTarget) {
	stateJson := `{
        "name": "test_policy",
        "mAdminState": "enabled",
        "Type": "file",
        "Size": 500,
        "LocalFile": "logtemp:///temp.log",
        "LogEvents": [{
            "Class": {
                "value": "all"
            },
            "Priority": "error"
        },{
            "Class": {
                "value": "apiconnect"
            },
            "Priority": "error"
        }]
    }`

	var state map[string]interface{}
	err := json.Unmarshal([]byte(stateJson), &state)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json map: %v\n", state)
}
