package main

import "fmt"
import "internal/account"

func main() {
	fmt.Println("Main")
	if account.CheckConfigFileExist() {
		fmt.Println("File Exists")
	} else {
		fmt.Println("File dosn't exists")
	}
}
