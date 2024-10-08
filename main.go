package main

import (
	"bufio"
	"fmt"
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
	m := LoadEntireMap(token)
	fmt.Printf("Loaded map. Total tiles: %d\n", len(m))
	a := NewRunner(token, "LegDay")
	go RoutineTaskSolver(a, m)
	b := NewRunner(token, "LegBot")
	go RoutineIronBars(b)
	c := NewRunner(token, "LegElf")
	go RoutineAshPlanks(c)
	d := NewRunner(token, "LegRock")
	RoutineCopperBars(d)
}
