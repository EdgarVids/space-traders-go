package account

import (
	"fmt"
	"os"
	"strings"
)

//import "fmt"
//import "net/http"
//import "time"
//import "io"

func ImportAccount() {
	//check if config file exist
	//check if token is outdated
	//if token outdated request new one
	//if config file dosnt exist ask about Username and Faction CreateAccount()
	//request new token
	//if token valid import token from EnvVariable
}

// CheckConfigFileExist returns bool if config file exists or not
func CheckConfigFileExist() bool {
	file, err := os.Open("config.conf")
	if err != nil {
		return false
	}
	defer file.Close()
	return true
}

// CreateConfigFile creates new config file contains Username and Faction
func CreateConfigFile() {
	if CheckConfigFileExist() {
		fmt.Println("Config file exists")
		return
	} else {
		text := []byte(createConfigData())
		err := os.WriteFile("config.conf", text, 0644)
		if err != nil {
			fmt.Println("Error creating a file")
		} else {
			fmt.Println("Config file created")
		}
	}
}

// CreateAccount
func CreateAccount() {
}

// getUsername gets input from user and return as string
func getUsername() string {
	fmt.Print("Username: ")
	var input string
	fmt.Scanln(&input)
	return input
	// Add check if username is not empty
}

// getFaction gets input from user and return as a string
func getFaction() string {
	fmt.Print("Faction: ")
	var input string
	fmt.Scanln(&input)
	if input == "" {
		input = "COSMIC" //default faction
	}
	return input
}

// createConfigData returns combination of Username and Faction
func createConfigData() string {
  username := getUsername()
  faction := getFaction()
  return strings.Join([]string{username, faction}, "-")
}

// checkToken cheks if env variable is set and is valid
func checkToken() bool{
  _, val := os.LookupEnv("SPA_Token")
  return val
}
