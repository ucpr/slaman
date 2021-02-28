package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	//"strings"
	firebase "firebase.google.com/go"
	"github.com/ucpr/slaman/models"
	"google.golang.org/api/option"
)

func main() {
	// initialize firebase configuration
	ctx := context.Background()
	opt := option.WithCredentialsFile("./secret.json")
	config := &firebase.Config{DatabaseURL: "https://slaman-dev-default-rtdb.firebaseio.com/"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ref := client.NewRef("server")
	usersRef := ref.Child("users")

	p := "/home/ucpr/Downloads/hoge/users.json"
	f, err := os.Open(p)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)

	var users []models.User
	if err := json.Unmarshal(b, &users); err != nil {
		log.Println(err)
		return
	}

	data := map[string]models.User{}
	for _, v := range users {
		data[v.ID] = v
	}

	// set to firebase realtime database
	err = usersRef.Set(ctx, data)
	if err != nil {
		log.Println(err)
		return
	}

	for _, p := range users {
		log.Println(p)
	}
}
