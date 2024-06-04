package account

import "os"
//import "fmt"
//import "net/http"
//import "time"
//import "io"

func ImportAccount(){
  //check if config file exist
  //check if token is outdated
  //if token outdated request new one
  //if config file dosnt exist ask about Username and Faction CreateAccount()
  //request new token
  //if token valid import token from EnvVariable
}

//CheckConfigFileExist returns bool if config file exists or not
func CheckConfigFileExist() bool{
  file, err := os.Open("config.conf")
  if err != nil {
    return false
  }
  defer file.Close()
  return true
}

//CreateConfigFile creates new config file contains Username and Faction
func CreateConfigFile(){
}

//CreateAccount
func CreateAccount() {
}
