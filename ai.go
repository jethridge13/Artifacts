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

func GatherLoop(a Runner) {
	loop := true
	for loop {
		res, status := a.Gathering()
		if status != 200 {
			loop = false
			PrintStatus(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}

func FightLoop(a Runner) {
	loop := true
	lossCount := 0
	for loop {
		res, status := a.Fight()
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

func CraftLoop(code string, a Runner) {
	item := Item{Code: code, Quantity: 1}
	loop := true
	for loop {
		res, status := a.Crafting(item)
		if status != 200 {
			loop = false
			PrintStatus(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}

func DepositAllInBank(a Runner) {
	// Move to bank
	c := Coordinate{X: 4, Y: 1}
	res, status := a.Move(c)
	if status != 200 {
		panic(status)
	} else {
		WaitOnCooldown(res)
	}
	// Get character inventory
	inventory := a.GetInventory()
	// Deposity EVERYTHING!
	for _, s := range inventory {
		res, status = a.BankDeposit(s.Code, s.Quantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}

func RoutineCopperBars(a Runner) {
	copper := "copper"
	for {
		// Move to copper mine
		c := Coordinate{X: 2, Y: 0}
		res, status := a.Move(c)
		if status != 200 && status != 490 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Gather until inventory full
		GatherLoop(a)
		// Move to forge
		c.X = 1
		c.Y = 5
		res, status = a.Move(c)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Craft copper bars
		CraftLoop(copper, a)
		// Move to bank
		c.X = 4
		c.Y = 1
		res, status = a.Move(c)
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
		res, status = a.BankDeposit(copper, quantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}

func RoutineAshGather(a Runner) {
	for {
		// Move to ash tree
		c := Coordinate{X: -1, Y: 0}
		res, status := a.Move(c)
		if status != 200 && status != 490 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Gather ash until inventory full
		GatherLoop(a)
		// Deposit everything in inventory
		DepositAllInBank(a)
	}
}

func RoutineAshPlanks(a Runner) {
	plank := "ash_plank"
	for {
		// Move to ash tree
		c := Coordinate{X: -1, Y: 0}
		res, status := a.Move(c)
		if status != 200 && status != 490 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Gather until inventory full
		GatherLoop(a)
		// Move to weaponsmith
		c.X = 2
		c.Y = 1
		res, status = a.Move(c)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Craft ash planks
		CraftLoop(plank, a)
		// Move to bank
		c.X = 4
		c.Y = 1
		res, status = a.Move(c)
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
			if s.Code == plank {
				quantity = s.Quantity
				break
			}
		}
		// Deposit ash planks into bank
		res, status = a.BankDeposit(plank, quantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}

func RoutineChickenFarming(a Runner) {
	cookedChicken := "cooked_chicken"
	egg := "egg"
	feather := "feather"
	for {
		// Move to chickens
		c := Coordinate{X: 0, Y: 1}
		res, status := a.Move(c)
		if status != 200 && status != 490 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Fight 'til death!
		FightLoop(a)
		// Move to kitchen
		c.X = 1
		res, status = a.Move(c)
		if status != 200 && status != 490 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		// Cook chicken
		CraftLoop(cookedChicken, a)
		// Move to bank
		c.X = 4
		res, status = a.Move(c)
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
		res, status = a.BankDeposit(cookedChicken, chickenQuantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		res, status = a.BankDeposit(egg, eggQuantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
		res, status = a.BankDeposit(feather, featherQuantity)
		if status != 200 {
			panic(status)
		} else {
			WaitOnCooldown(res)
		}
	}
}
