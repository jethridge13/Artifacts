package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Slot struct {
	Slot string `json:"slot"`
}

type Item struct {
	Code  string `json:"code"`
	Slot  string `json:"slot"`
	Count string `json:"count"`
}

func sendActionRequest(action string, body []byte) ([]byte, int) {
	character := "LegDay"
	endpoint := fmt.Sprintf("/my/%s/action/%s", character, action)
	return sendRequest(body, endpoint)
}

func sendCharacterRequest(name string) ([]byte, int) {
	endpoint := fmt.Sprintf("/characters/%s", name)
	return sendRequest([]byte{}, endpoint)
}

func sendRequest(body []byte, endpoint string) ([]byte, int) {
	server := "https://api.artifactsmmo.com"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6IkpvaGFuU2t1bGxjcnVzaGVyIiwicGFzc3dvcmRfY2hhbmdlZCI6IiJ9.XSij4JbWgWhHyExSkV8aIt6373cNr6HXzGQEP4xn2Ks"
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

func Move(c Coordinate) ([]byte, int) {
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return sendActionRequest("move", b)
}

func Fight() ([]byte, int) {
	return sendActionRequest("fight", []byte{})
}

func Gathering() ([]byte, int) {
	fmt.Println("Gathering at current location")
	return sendActionRequest("gathering", []byte{})
}

func Crafting(item Item) ([]byte, int) {
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return sendActionRequest("crafting", b)
}

func Recycling() {

}

func Equip(item Item) ([]byte, int) {
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return sendActionRequest("equip", b)
}

func Unequip(slot Slot) ([]byte, int) {
	b, err := json.Marshal(slot)
	if err != nil {
		panic(err)
	}
	return sendActionRequest("unequip", b)
}

func Delete() {

}

func BankDeposit() {

}

func BankWithdraw() {

}

func BankDepositGold() {

}

func BankWithdrawGold() {

}

func GeBuy() {

}

func GeSell() {

}

func TaskAccept() {

}

func TaskComplete() {

}

func GetInventory() []InventorySlot {
	res, code := sendCharacterRequest("LegDay")
	if code != 200 {
		panic(code)
	}
	var response CharacterSchema
	json.Unmarshal(res, &response)
	return response.Data.Character.Inventory
}
