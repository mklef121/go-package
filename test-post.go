package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mklef121/go-package/user"
)

var MIN = 0
var MAX = 26

func main() {
	user.Hostname = "localhost"
	user.Port = 5432
	user.Username = "miracool"
	user.Password = "pass"
	user.Database = "test_db"

	data, err := user.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	//generate random string for username
	random_username := getString(5)

	t := user.Userdata{
		Username:    random_username,
		Name:        "Mihalis",
		Surname:     "Tsoukalos",
		Description: "This is me!",
	}

	id := user.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	fmt.Println("User successfully '" + random_username + "' added to database")

	err = user.DeleteUser(id)
	if err != nil {
		fmt.Println("deleting error", err)
	}

	//Creating a new user
	random_username = getString(5)

	t = user.Userdata{
		Username:    random_username,
		Name:        "Mihalis",
		Surname:     "Tsoukalos",
		Description: "This is me!",
	}

	id = user.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	t = user.Userdata{
		Username:    random_username,
		Name:        "Mihalis",
		Surname:     "Tsoukalos",
		Description: "This might not be me!",
	}

	err = user.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}
}

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := Random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}
