// config using yaml file

package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {

	// reads yaml files into a struct parser 
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// print the values from the yaml file 
	fmt.Println(config.Get("path")) // obtain value of a string using Get
	fmt.Println(config.GetBool("enabled")) // obtain value of a bool using GetBool

}