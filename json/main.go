package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var print = fmt.Println

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	Specie string   `json:"specie"`
	Sound  string   `json:"sound"`
	Age    int      `json:"age"`
	Food   []string `json:"food"`
	name   string
}

func main() {
	print("-- Go JSON --")
	boolRes, _ := json.Marshal(true)
	print(string(boolRes))

	intRes, _ := json.Marshal(1)
	print(string(intRes))

	floatRes, _ := json.Marshal(3.33)
	print(string(floatRes))

	stringRes, _ := json.Marshal("cat")
	print(string(stringRes))
	// have to convert them into a string, otherwise they end up being a slice of bytes

	intSliceRes, _ := json.Marshal([]int{1, 2, 3})
	print(string(intSliceRes))

	stringSliceRes, _ := json.Marshal([]string{"pera", "mika"})
	print(string(stringSliceRes))

	mapRes, _ := json.Marshal(map[int]string{1: "Pera", 2: "Mika"})
	print(string(mapRes))

	animal1 := &Animal{ // works either with a pointer or without
		Specie: "Dog",
		Sound:  "Woof",
		Age:    2,
		name:   "Barkie", // won't be marshalled
	}

	animal1Res, _ := json.Marshal(animal1)
	print(string(animal1Res))

	print("-- Unmarshaling --")

	// CUSTOM STRUCT EXAMPLE
	var animal2 Animal = Animal{}
	// unmarshaling always expects 1) a byte slice and 2) a pointer
	json.Unmarshal([]byte(`{"specie":"Cat","sound":"Meow","age":1,"food":["tuna","catnip"]}`), &animal2)
	print("animal2:", animal2)
	print("animal2.Food", animal2.Food)
	// can easily access array items with a custom struct
	print("animal2.Food[1]", animal2.Food[1])

	// MAP EXAMPLE
	var unmarshaledData1 map[string]interface{}
	json.Unmarshal([]byte(`{"specie":"Cat","sound":"Meow","age":1,"food":["tuna","catnip"],"address":{"streetName":"Beach Street","streetNumber":1}}`), &unmarshaledData1)
	print(unmarshaledData1)
	// no need to assert it as a string just to print out, but if I wanted to assign it to a string
	// variable then I would need to
	// print(unmarshaledData1["specie"].(string))
	var animal1Specie string = unmarshaledData1["specie"].(string)
	print("animal1Specie string value is", animal1Specie)
	print(unmarshaledData1["specie"])
	print(unmarshaledData1["food"])

	// print(unmarshaledData1["food"][0]) cannot access it like this,
	// need to grab the food slice first and assert it as an interface slice
	food, _ := unmarshaledData1["food"].([]interface{})
	// it's still just an interface type, not very useful
	print("food", food)
	// I cannot even directly assert "food" as a slice of strings
	// like this // food, _ := unmarshaledData1["food"].([]string)
	// I would have to first grab the food slice, then assert each element

	// I could assert the item as a string
	var assertedFoodItemString string = food[1].(string)
	print("assertedFoodItemString:", assertedFoodItemString)

	// same thing with nested objects
	// print(unmarshaledData1["address"]["streetName"]) // cannot do this
	// have to cast the embedded structure as a map
	address, _ := unmarshaledData1["address"].(map[string]interface{})
	streetName, _ := address["streetName"].(string)
	print("streetName is", streetName)

	// can also encode directly to stdout
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(map[string]string{"pera": "peric"})
}
