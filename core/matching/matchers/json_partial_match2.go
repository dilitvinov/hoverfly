package matchers

import (
	"encoding/json"
)

func JsonPartialMatch2(match interface{}, toMatch string) bool {

	var expected, actual map[string]interface{}
	matchString, ok := match.(string)
	if !ok {
		return false
	}
	err0 := json.Unmarshal([]byte(matchString), &expected)
	err1 := json.Unmarshal([]byte(toMatch), &actual)
	if err0 != nil || err1 != nil {
		return false
	}

	return true
}

func getAllNodesFromMap(current map[string]interface{}) []map[string]interface{} {

	var allNodes = make([]map[string]interface{}, 0, 0)

	allNodes = append(allNodes, current)

	for _, val := range current {
		if innerMap, ok := val.(map[string]interface{}); ok {
			allNodes = append(allNodes, getAllNodesFromMap(innerMap)...)
		} else if innerArray, ok := val.([]interface{}); ok {
			allNodes = append(allNodes, getAllNodesFromArray(innerArray)...)
		}
	}
	return allNodes
}

func getAllNodesFromArray(current []interface{}) []map[string]interface{} {
	var allNodes = make([]map[string]interface{}, 0, 0)

	for _, val := range current {
		if innerMap, ok := val.(map[string]interface{}); ok {
			allNodes = append(allNodes, getAllNodesFromMap(innerMap)...)
		} else if innerArray, ok := val.([]interface{}); ok {
			allNodes = append(allNodes, getAllNodesFromArray(innerArray)...)
		}
	}
	return allNodes
}
