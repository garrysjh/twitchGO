package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"twitchGO/helpers"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the username you want to see the details of: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Print("What do you want to see about this user? 1.profilepic 2.username 3.created \n")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)
	switch option {
	case "profilepic", fmt.Sprint(1):
		fmt.Println(helpers.GetUserDp(username))
	case "username", fmt.Sprint(2):
		fmt.Println(helpers.GetLoginName(username))
	case "created", fmt.Sprint(3):
		fmt.Println(helpers.GetUserCreatedAt(username))
	default:
		fmt.Printf("This is the option, %s", option)
		fmt.Println("Invalid option")
	}

}
