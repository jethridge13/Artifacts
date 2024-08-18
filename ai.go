package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func WaitOnCooldown(res []byte) {
	var response SkillDataSchema
	json.Unmarshal(res, &response)
	d := time.Duration(response.Data.Cooldown.Remaining_Seconds) * time.Second
	fmt.Println(fmt.Sprintf("Sleeping for %s seconds", d))
	time.Sleep(d)
}

func PrintStatus(code int) {
	if code == 486 {
		fmt.Println("Action already in progress")
	} else if code == 493 {
		fmt.Println("Not skill level required")
	} else if code == 497 {
		fmt.Println("Inventory full")
	} else if code == 498 {
		fmt.Println("Character not found")
	} else if code == 499 {
		fmt.Println("Character in cooldown")
	} else if code == 598 {
		fmt.Println("Resource not found on map")
	}
}

func GatherLoop() {
	loop := true
	for loop {
		res, status := Gathering()
		if status != 200 {
			loop = false
			PrintStatus(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}

func CopperLoop() {

}

func CraftMax(item string) {

}
