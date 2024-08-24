package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Runner struct {
	Token     string
	Name      string
	Character Character
}

func NewRunner(token string, name string) Runner {
	a := new(Runner)
	a.Token = token
	a.Name = name
	a.Character = a.GetCharacter()
	return *a
}

func (a Runner) sendActionRequest(action string, body []byte) ([]byte, int) {
	character := a.Name
	endpoint := fmt.Sprintf("/my/%s/action/%s", character, action)
	return a.sendRequest(body, endpoint, "POST")
}

func (a Runner) sendCharacterRequest() ([]byte, int) {
	fmt.Println(a.Name)
	endpoint := fmt.Sprintf("/characters/%s", a.Name)
	return a.sendRequest([]byte{}, endpoint, "GET")
}

func (a *Runner) sendRequest(body []byte, endpoint string, method string) ([]byte, int) {
	server := "https://api.artifactsmmo.com"
	token := a.Token
	url := fmt.Sprintf("%s%s", server, endpoint)
	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var responseBody GenericSchema
	err = json.Unmarshal(response, &responseBody)
	if err == nil {
		a.Character = responseBody.Data.Character
	}
	return response, res.StatusCode
}

func sendRequest(body []byte, endpoint string, method string, token string) ([]byte, int) {
	server := "https://api.artifactsmmo.com"
	url := fmt.Sprintf("%s%s", server, endpoint)
	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return response, res.StatusCode
}

func (a Runner) Move(c Coordinate) ([]byte, int) {
	fmt.Printf("%v: %s: Moving to %d, %d\n", time.Now(), a.Name, c.X, c.Y)
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return a.sendActionRequest("move", b)
}

func (a Runner) Fight() ([]byte, int) {
	fmt.Printf("%v: %s: Fight!\n", time.Now(), a.Name)
	return a.sendActionRequest("fight", []byte{})
}

func (a Runner) Gathering() ([]byte, int) {
	fmt.Printf("%v: %s: Gathering at current location\n", time.Now(), a.Name)
	return a.sendActionRequest("gathering", []byte{})
}

func (a Runner) Crafting(item Item) ([]byte, int) {
	fmt.Printf("%v: %s: Crafting %s\n", time.Now(), a.Name, item.Code)
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return a.sendActionRequest("crafting", b)
}

func Recycling() {

}

func (a Runner) Equip(item Item) ([]byte, int) {
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return a.sendActionRequest("equip", b)
}

func (a Runner) Unequip(slot Slot) ([]byte, int) {
	b, err := json.Marshal(slot)
	if err != nil {
		panic(err)
	}
	return a.sendActionRequest("unequip", b)
}

func Delete() {

}

func (a Runner) BankDeposit(code string, quantity int) ([]byte, int) {
	fmt.Printf("%v: %s: Depositing %d %s into bank\n", time.Now(), a.Name, quantity, code)
	item := Item{Code: code, Quantity: quantity}
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return a.sendActionRequest("bank/deposit", b)
}

func (a Runner) BankWithdraw(code string, quantity int) ([]byte, int) {
	fmt.Printf("%v: %s: Requesting %d %s from bank\n", time.Now(), a.Name, quantity, code)
	item := Item{Code: code, Quantity: quantity}
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return a.sendActionRequest("bank/withdraw", b)
}

func BankDepositGold() {

}

func BankWithdrawGold() {

}

func GeBuy() {

}

func GeSell() {

}

func (a Runner) TaskAccept() ([]byte, int) {
	fmt.Printf("%v: %s: Accepting task\n", time.Now(), a.Name)
	return a.sendActionRequest("task/new", []byte{})
}

func (a Runner) TaskComplete() ([]byte, int) {
	fmt.Printf("%v: %s: Completing task\n", time.Now(), a.Name)
	return a.sendActionRequest("task/complete", []byte{})
}

func (a Runner) GetInventory() []InventorySlot {
	res, code := a.sendCharacterRequest()
	if code != 200 {
		panic(code)
	}
	var response CharacterSchema
	json.Unmarshal(res, &response)
	return response.Data.Character.Inventory
}

func (a Runner) GetCharacter() Character {
	res, code := a.sendCharacterRequest()
	if code != 200 {
		panic(code)
	}
	var response GetCharacterSchema
	json.Unmarshal(res, &response)
	return response.Data.Character
}

func (a Runner) FindNearestEntity(entity string, maps map[Coordinate]MapSchema) (Coordinate, bool) {
	visited := make(map[Coordinate]bool)
	queue := make([]Coordinate, 0)
	start := Coordinate{X: a.Character.X, Y: a.Character.Y}
	fmt.Printf("Searching for %s starting at (%d, %d)\n", entity, start.X, start.Y)
	queue = append(queue, start)
	for len(queue) > 0 {
		// Check if position contains searching entity
		var c Coordinate
		c, queue = queue[0], queue[1:]
		if maps[c].Content.Code == entity {
			fmt.Printf("Found %s at (%d, %d)\n", entity, c.X, c.Y)
			return c, true
		}
		// Search neighboring squares
		visited[c] = true
		north := Coordinate{X: c.X, Y: c.Y - 1}
		east := Coordinate{X: c.X + 1, Y: c.Y}
		south := Coordinate{X: c.X, Y: c.Y + 1}
		west := Coordinate{X: c.X - 1, Y: c.Y}
		_, ok := visited[north]
		_, exists := maps[north]
		if !ok && exists {
			queue = append(queue, north)
			visited[north] = true
		}
		_, ok = visited[east]
		_, exists = maps[east]
		if !ok && exists {
			queue = append(queue, east)
			visited[east] = true
		}
		_, ok = visited[south]
		_, exists = maps[south]
		if !ok && exists {
			queue = append(queue, south)
			visited[south] = true
		}
		_, ok = visited[west]
		_, exists = maps[west]
		if !ok && exists {
			queue = append(queue, west)
			visited[west] = true
		}
	}
	fmt.Printf("Could not find %s on map\n", entity)
	return Coordinate{X: 0, Y: 0}, false
}

func LoadEntireMap(token string) map[Coordinate]MapSchema {
	m := make(map[Coordinate]MapSchema)
	res, err := sendRequest([]byte{}, "/maps/", "GET", token)
	if err != 200 {
		panic(err)
	}
	var response MapResponseSchema
	json.Unmarshal(res, &response)
	for _, s := range response.Data {
		c := Coordinate{X: s.X, Y: s.Y}
		m[c] = s
	}
	page := response.Page + 1
	pages := response.Pages
	for page <= pages {
		res, err := sendRequest([]byte{}, fmt.Sprintf("/maps/?page=%d", page), "GET", token)
		if err != 200 {
			panic(err)
		}
		var response MapResponseSchema
		json.Unmarshal(res, &response)
		for _, s := range response.Data {
			c := Coordinate{X: s.X, Y: s.Y}
			m[c] = s
		}
		page = response.Page + 1
	}
	return m
}
