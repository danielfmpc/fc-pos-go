package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Cotacao struct {
	UsdBrl USDBRLResponse `json:"USDBRL"`
}

type USDBRLResponse struct {
	gorm.Model
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	TimeStamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

var (
	db *gorm.DB
)

func main() {
	var err error
	db, err = initConfig()
	if err != nil {
		log.Panic("Erro ao inicializar a configuração:", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", HandleCotacao)

	log.Println("Servidor iniciado na porta 8080")
	log.Panic(http.ListenAndServe(":8080", mux))
}

func HandleCotacao(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	cotacao, err := buscaCotacao(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	insertCotacao(&cotacao.UsdBrl, r)

	data := map[string]string{"bid": cotacao.UsdBrl.Bid}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func initConfig() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("cotacao.db"), &gorm.Config{})
	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}

	err = db.AutoMigrate(&USDBRLResponse{})
	if err != nil {
		log.Println("Erro ao realizar migration:", err)
		return nil, err
	}

	return db, nil
}

func insertCotacao(cotacao *USDBRLResponse, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Millisecond)
	defer cancel()

	err := db.WithContext(ctx).Create(cotacao).Error
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout ao inserir cotação no banco de dados: tempo limite de 10ms excedido")
		} else {
			log.Println("Erro ao inserir cotação no banco de dados:", err)
		}
	}
}

func buscaCotacao(r *http.Request) (*Cotacao, error) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Erro ao criar a requisição:", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout ao buscar cotação na API: tempo limite de 200ms excedido")
		} else {
			log.Println("Erro ao fazer a requisição:", err)
		}
		return nil, err
	}
	defer res.Body.Close()

	var cotacao Cotacao
	err = json.NewDecoder(res.Body).Decode(&cotacao)
	if err != nil {
		log.Println("Erro ao decodificar a resposta:", err)
		return nil, err
	}

	log.Printf("Cotação recebida com sucesso - Bid: %s", cotacao.UsdBrl.Bid)
	return &cotacao, nil
}
