/*
Example of parsing arbitrary json when we don't know the data structure 

Passing the JSON into an interface{} and use switch to print out various types 

*/


package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{
"firstName": "Jean",
"lastName": "Bartik",
"age": 86,
"education": [
	{
		"institution": "Northwest Missouri State Teachers College",
		"degree": "Bachelor of Science in Mathematics"
	},
	{
		"institution": "University of Pennsylvania",
		"degree": "Masters in English"
	}
],
"spouse": "William Bartik",
"children": [
	"Timothy John Bartik",
	"Jane Helen Bartik",
	"Mary Ruth Bartik"
]
}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f)

	m := f.(map[string]interface{})
	fmt.Println(m["firstName"])

	fmt.Print("interface{} ")
	printJSON(f)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}


// When the JSON is unmarshaled, the
// values in JSON are converted into the following Go types:
//  bool for JSON Boolean
//  float64 for JSON numbers
//  []interface{} for JSON arrays
//  map[string]interface{} for JSON objects
//  nil for JSON null
//  string for JSON strings