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
	case 478:
		fmt.Println("Missing item or insufficient quantity")
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
		fmt.Printf("Unknown code %d\n", code)
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

func FightLoop() {
	loop := true
	lossCount := 0
	for loop {
		res, status := Fight()
		if status != 200 {
			loop = false
			PrintStatus(status)
		} else {
			WaitOnCooldown(res)
		}
		var response CharacterFightDataSchema
		json.Unmarshal(res, &response)
		if response.Data.Fight.Result == "lose" {
			lossCount += 1
		} else {
			lossCount = 0
		}
		if lossCount == 5 {
			loop = false
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

func RoutineChickenFarming() {
	cookedChicken := "cooked_chicken"
	egg := "egg"
	feather := "feather"
	for {
		// Move to chickens
		c := Coordinate{X: 0, Y: 1}
		res, status := Move(c)
		if status != 200 && status != 490 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Fight 'til death!
		FightLoop()
		// Move to kitchen
		c.X = 1
		res, status = Move(c)
		if status != 200 && status != 490 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Cook chicken
		CraftLoop(cookedChicken)
		// Move to bank
		c.X = 4
		res, status = Move(c)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Get character inventory
		var response GenericSchema
		json.Unmarshal(res, &response)
		chickenQuantity := 0
		eggQuantity := 0
		featherQuantity := 0
		for _, s := range response.Data.Character.Inventory {
			if s.Code == cookedChicken {
				chickenQuantity = s.Quantity
			} else if s.Code == egg {
				eggQuantity = s.Quantity
			} else if s.Code == feather {
				featherQuantity = s.Quantity
			}
		}
		// Deposit items into bank
		res, status = BankDeposit(cookedChicken, chickenQuantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		res, status = BankDeposit(egg, eggQuantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		res, status = BankDeposit(feather, featherQuantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}
