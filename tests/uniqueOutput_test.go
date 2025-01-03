package tests

import (
	"encoding/json"
	"fmt"
	"lem-in/utils"
	"lem-in/utils/parser"
	"testing"
)

func TestFindPaths(t *testing.T) {
	antFarm, err := parser.ParseFile("../examples/test.txt")
	if err != nil {
		t.Fatalf("Failed to parse file: %v", err)
	}

	numCalls := 1000
	outputs := make(map[string]struct{})

	for i := 0; i < numCalls; i++ {
		paths := utils.FindPaths(antFarm)

		outputStr, err := slice2Json(utils.ConvertPaths(antFarm, paths))
		if err != nil {
			t.Errorf("Error converting paths to JSON: %v", err)
			continue
		}

		outputs[outputStr] = struct{}{}
	}

	t.Log("Unique outputs:")
	i := 0
	for output := range outputs {
		var data []any
		err = json.Unmarshal([]byte(output), &data)
		if err != nil {
			t.Errorf("Error converting JSON to slice: %v", err)
			continue
		}
		t.Log("result N°", i)
		i++
		for _, v := range data {
			fmt.Println(v)
		}
	}
}

func slice2Json(s [][]string) (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
