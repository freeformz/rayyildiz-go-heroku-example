package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Payload struct {
	Stuffs Data
}

type Data struct {
	Fruits      Fruit
	Vegenatbles Vegetable
}

type Fruit map[string]int

type Vegetable map[string]int

func getJsonResult() ([]byte, error) {
	f := make(map[string]int)
	f["Apple"] = 3
	f["Banana"] = 21

	v := make(map[string]int)
	v["Domatoes"] = 6
	v["Cucumber"] = 13

	d := Data{f, v}

	p := Payload{d}

	return json.MarshalIndent(p, "", " ")
}

func main() {
	log.SetFlags(0)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		response, err := getJsonResult()
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("Error: %q", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	log.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
