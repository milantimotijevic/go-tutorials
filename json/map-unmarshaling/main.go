package main

import (
	"encoding/json"
	"fmt"
)

var print = fmt.Println

func main() {
	print("-- Additional Map Unmarshaling Examples --")

	jsonString := `{"name":"Pera Peric","age":30,"favoriteFoods":["apples","steaks"],"address":{"streetName":"Beach Street","streetNumber":1}}`
	var personMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &personMap)

	// grab favoriteFoods first, then assert it as an array of interfaces,
	// then grab index, then assert as string
	print(personMap["favoriteFoods"].([]interface{})[0].(string))

	// grab address first, then assert it as a map of K:V string:interface,
	// then grab specific key, then assert it as string
	print(personMap["address"].(map[string]interface{})["streetName"].(string))
}
