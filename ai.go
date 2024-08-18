package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func WaitOnCooldown(res []byte) {
	var response GenericSchema
	json.Unmarshal(res, &response)
	d := time.Duration(response.Data.Cooldown.Remaining_Seconds) * time.Second
	fmt.Println(fmt.Sprintf("Sleeping for %s seconds", d))
	time.Sleep(d)
}

func PrintStatus(code int) {
	switch code {
	case 486:
		fmt.Println("Action already in progress")
	case 493:
		fmt.Println("Not skill level required")
	case 497:
		fmt.Println("Inventory full")
	case 498:
		fmt.Println("Character not found")
	case 499:
		fmt.Println("Character in cooldown")
	case 598:
		fmt.Println("Resource not found on map")
	default:
		fmt.Printf("Unknown code %d", code)
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

func CraftLoop(code string) {
	item := Item{Code: code, Quantity: 1}
	loop := true
	for loop {
		res, status := Crafting(item)
		if status != 200 {
			loop = false
			PrintStatus(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}

func RoutineCopperBars() {
	copper := "copper"
	for {
		// Move to copper mine
		c := Coordinate{X: 2, Y: 0}
		res, status := Move(c)
		if status != 200 && status != 490 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Gather until inventory full
		GatherLoop()
		// Move to forge
		c.X = 1
		c.Y = 5
		res, status = Move(c)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Craft copper bars
		CraftLoop(copper)
		// Move to bank
		c.X = 4
		c.Y = 1
		res, status = Move(c)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Get character inventory
		var response GenericSchema
		json.Unmarshal(res, &response)
		quantity := 0
		for _, s := range response.Data.Character.Inventory {
			if s.Code == copper {
				quantity = s.Quantity
				break
			}
		}
		// Deposit copper bars into bank
		res, status = BankDeposit(copper, quantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}
