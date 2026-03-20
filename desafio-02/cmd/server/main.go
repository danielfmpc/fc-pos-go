package main

import (
	"log"
	"net/http"

	"github.com/danielfmpc/pos-go-expert-desafio02/internal/handler"
	"github.com/danielfmpc/pos-go-expert-desafio02/internal/infra/rest"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cepHandler := handler.NewCepHandler(&rest.ViaCepClient{}, &rest.BrasilApiClient{})

	r.Get("/cep/{cep}", cepHandler.ConsultaCepHandler)
	r.Get("/brasilapi/{cep}", cepHandler.ConsultaCepBrasilAPIHandler)
	r.Get("/viacep/{cep}", cepHandler.ConsultaCepViaCepHandler)
	r.Get("/timeout/{cep}", cepHandler.ConsultaCepTimeOutHandler)

	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	log.Println("Servidor iniciado na porta :8000")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
