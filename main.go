package main

import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
	//"strings"
	"github.com/ucpr/slaman/models"
)


func main() {
	p := ""
	f, err := os.Open(p)
	if err != nil {
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)

	var users []models.User
	if err := json.Unmarshal(b, &users); err != nil {
		fmt.Println(err)
	}

	for _, p := range users {
		fmt.Println(p)
	}
}