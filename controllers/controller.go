package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CreateEventSearch(w http.ResponseWriter, r *http.Request) {
	url := "https://api.checks.truora.com/v1/checks"

	payload := strings.NewReader("national_id=42388604&country=PE&type=person&user_authorized=true")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Truora-API-Key", os.Getenv("api-key-truora"))

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	url := "https://api.checks.truora.com/v1/checks/CHKaaf64fa231c28b88f3b3f1010bd90b0c"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Truora-API-Key", os.Getenv("api-key-truora"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la wallet!!")
}
