package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Panic("Erro ao criar a requisição:", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Panic("Timeout ao buscar cotação: tempo limite de 300ms excedido")
		}
		log.Panic("Erro ao fazer a requisição:", err)
	}
	defer res.Body.Close()

	var cotacao Cotacao
	err = json.NewDecoder(res.Body).Decode(&cotacao)
	if err != nil {
		log.Panic("Erro ao decodificar a resposta:", err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Panic("Erro ao criar o arquivo:", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Dólar: %s\n", cotacao.Bid)
	if err != nil {
		log.Panic("Erro ao escrever no arquivo:", err)
	}

	log.Println("Resposta recebida com status:", res.StatusCode)
}
