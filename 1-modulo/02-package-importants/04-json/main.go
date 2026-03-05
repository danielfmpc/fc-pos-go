package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero string  `json:"n"`
	Saldo  float64 `json:"s"`
}

func main() {
	conta := Conta{
		Numero: "12345",
		Saldo:  1000.0,
	}

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonData := []byte(`{"Numero":"67890","Saldo":2500.5}`)
	// jsonData := []byte(`{"n":"67890","s":2500.5}`)
	var contaX Conta
	err = json.Unmarshal(jsonData, &contaX)
	if err != nil {
		panic(err)
	}

}
