// parsing a JSON configuration file 

package main 

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path string 
}

func main() {
	// open configuration file
	file, _ := os.Open("conf.json")
	defer file.Close()

	// docoding JSON file into an instance of the configuration struct 
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(conf.Path)
}