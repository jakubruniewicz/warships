package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GamePayload struct {
	Coords      []string
	Desc        string
	Nick        string
	Target_nick string
	Wpbot       bool
}

func main() {
	url := "https://go-pjatk-server.fly.dev/api/game"

	payload := GamePayload{
		Coords: []string{"A1",
			"A3",
			"B9",
			"C7",
			"D1",
			"D2",
			"D3",
			"D4",
			"D7",
			"E7",
			"F1",
			"F2",
			"F3",
			"F5",
			"G5",
			"G8",
			"G9",
			"I4",
			"J4",
			"J8"},
		Desc:        "My_first_game",
		Nick:        "Kuba_R",
		Target_nick: "",
		Wpbot:       true,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Błąd podczas konwersji danych na format JSON:", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln("Błąd podczas tworzenia żądania HTTP:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Błąd podczas wysyłania żądania HTTP:", err)
	}
	resp.Body.Close()

	header := resp.Header.Get("x-auth-token")

	fmt.Println(header)

}
