package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Runner struct {
	Token     string
	Character string
}

func (a Runner) sendActionRequest(action string, body []byte) ([]byte, int) {
	character := a.Character
	endpoint := fmt.Sprintf("/my/%s/action/%s", character, action)
	return a.sendRequest(body, endpoint)
}

func (a Runner) sendCharacterRequest() ([]byte, int) {
	endpoint := fmt.Sprintf("/characters/%s", a.Character)
	return a.sendRequest([]byte{}, endpoint)
}

func (a Runner) sendRequest(body []byte, endpoint string) ([]byte, int) {
	server := "https://api.artifactsmmo.com"
	token := a.Token
	url := fmt.Sprintf("%s%s", server, endpoint)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
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
	fmt.Printf("Moving to %d, %d\n", c.X, c.Y)
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return a.sendActionRequest("move", b)
}

func (a Runner) Fight() ([]byte, int) {
	fmt.Println("Fight!")
	return a.sendActionRequest("fight", []byte{})
}

func (a Runner) Gathering() ([]byte, int) {
	fmt.Println("Gathering at current location")
	return a.sendActionRequest("gathering", []byte{})
}

func (a Runner) Crafting(item Item) ([]byte, int) {
	fmt.Printf("Crafting %s\n", item.Code)
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
	fmt.Printf("Depositing %d %s into bank\n", quantity, code)
	item := Item{Code: code, Quantity: quantity}
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return a.sendActionRequest("bank/deposit", b)
}

func (a Runner) BankWithdraw(code string, quantity int) ([]byte, int) {
	fmt.Printf("Requesting %d %s from bank\n", quantity, code)
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
	fmt.Printf("Accepting task\n")
	return a.sendActionRequest("task/new", []byte{})
}

func (a Runner) TaskComplete() ([]byte, int) {
	fmt.Println("Completing task")
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
