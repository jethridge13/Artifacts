package main

import (
	"bufio"
	"os"
)

func GetFileScanner(path string) *bufio.Scanner {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func loadToken() string {
	scanner := GetFileScanner("token.txt")
	var token string
	for scanner.Scan() {
		token = scanner.Text()
	}
	return token
}

func main() {
	token := loadToken()
	a := Runner{Token: token, Character: "LegDay"}
	RoutineChickenFarming(a)
}
