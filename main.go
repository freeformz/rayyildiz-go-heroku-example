package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"os"
)

type Payload struct {
    Stuffs Data
}

type Data struct {
    Fruits Fruit
    Vegenatbles Vegetable
}

type Fruit map[string] int

type Vegetable map[string] int

func getJsonReslt() ([]byte, error) {
    f := make(map[string] int)
    f["Apple"] = 3
    f["Banana"] = 21


    v := make(map[string] int)
    v["Domatoes"] = 6
    v["Cucumber"] = 13


    d := Data{f,v}

    p := Payload{d}

    return json.MarshalIndent(p,"", " ")
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

	    response, err := getJsonReslt()
	    if ( err != nil){
	        panic(err)
	    }

	    fmt.Fprintf(w,string(response))
	})

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
