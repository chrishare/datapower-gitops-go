package dpobjects

import (
	"encoding/json"
	"fmt"
)

type LogTarget struct {
	// This inherits DPObjects's functions etc as top level LogCat functions
	DPObjectCommon
}

func NewLogTarget() (*LogTarget, error) {
	return &LogTarget{
		DPObjectCommon: DPObjectCommon{InternalState: map[string]interface{}{}, ObjectType: "LogTarget"},
	}, nil
}

func NewLogTargetFromState(initialState map[string]interface{}) (*LogTarget, error) {
	return &LogTarget{
		DPObjectCommon: DPObjectCommon{InternalState: initialState, ObjectType: "LogTarget"},
	}, nil
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
