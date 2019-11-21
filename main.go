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

//This is the part that does all the real work

func main() {
	fmt.Println("Hello World!")

	export("david", "chang")

	gitadd , _ := exec.Command("git", "add", "*").Output()
	fmt.Println(string(gitadd))

	gitcommit , _ := exec.Command("git", "commit", "-m", "automatically committed").Output()
	fmt.Println(string(gitcommit))

	gitpush , _ := exec.Command("git", "push").Output()
	fmt.Println(string(gitpush))
	
}

