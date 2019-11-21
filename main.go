package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os/exec"
)

// Potato = sample structure
type Potato struct {
	FirstName string `json:"firstname"`
	LastName string	`json:"lastname"`
}

//function to export a json
func export(fname, lname string) {
	potato := &Potato {
		FirstName: fname,
		LastName: lname }
	potatoByte, _ := json.Marshal(potato)

	err := ioutil.WriteFile("output.json", potatoByte, 0644)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	fmt.Println("Hello World!")

	export("david", "chang")

	gitadd , _ := exec.Command("ls").Output()
	fmt.Println(string(gitadd))
	
}

