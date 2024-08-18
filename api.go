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

func sendRequest(action string, body []byte) ([]byte, int) {
	server := "https://api.artifactsmmo.com"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6IkpvaGFuU2t1bGxjcnVzaGVyIiwicGFzc3dvcmRfY2hhbmdlZCI6IiJ9.XSij4JbWgWhHyExSkV8aIt6373cNr6HXzGQEP4xn2Ks"
	character := "LegDay"
	url := fmt.Sprintf("%s/my/%s/action/%s", server, character, action)
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
	return sendRequest("move", b)
}

func Fight() ([]byte, int) {
	return sendRequest("fight", []byte{})
}

func Gathering() ([]byte, int) {
	fmt.Println("Gathering at current location")
	return sendRequest("gathering", []byte{})
}

func Crafting(item Item) ([]byte, int) {
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return sendRequest("crafting", b)
}

func Recycling() {

}

func Equip(item Item) ([]byte, int) {
	b, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return sendRequest("equip", b)
}

func Unequip(slot Slot) ([]byte, int) {
	b, err := json.Marshal(slot)
	if err != nil {
		panic(err)
	}
	return sendRequest("unequip", b)
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
