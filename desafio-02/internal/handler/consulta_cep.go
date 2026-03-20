package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/danielfmpc/pos-go-expert-desafio02/internal/dto"
	"github.com/danielfmpc/pos-go-expert-desafio02/internal/infra/rest"
	"github.com/go-chi/chi/v5"
)

var cepRegex = regexp.MustCompile(`^\d{5}-?\d{3}$`)

type CepHandler struct {
	viaCepClient    rest.CepInterface
	brasilApiClient rest.CepInterface
}

func NewCepHandler(viaCepClient rest.CepInterface, brasilApiClient rest.CepInterface) *CepHandler {
	return &CepHandler{viaCepClient: viaCepClient, brasilApiClient: brasilApiClient}
}

func (h *CepHandler) ConsultaCepHandler(w http.ResponseWriter, r *http.Request) {
	h.consultaCep(w, r, 0, 0)
}

func (h *CepHandler) ConsultaCepBrasilAPIHandler(w http.ResponseWriter, r *http.Request) {
	h.consultaCep(w, r, 2*time.Second, 0)
}

func (h *CepHandler) ConsultaCepViaCepHandler(w http.ResponseWriter, r *http.Request) {
	h.consultaCep(w, r, 0, 2*time.Second)
}

func (h *CepHandler) ConsultaCepTimeOutHandler(w http.ResponseWriter, r *http.Request) {
	h.consultaCep(w, r, 2*time.Second, 2*time.Second)
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func validarCep(cep string) bool {
	return cepRegex.MatchString(cep)
}

func (h *CepHandler) consultaCep(w http.ResponseWriter, r *http.Request, viaCepDelay, brasilApiDelay time.Duration) {
	cep := chi.URLParam(r, "cep")

	if !validarCep(cep) {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "CEP inválido"})
		return
	}

	ctx := r.Context()
	viaCepCh := make(chan dto.Cep, 1)
	brasilApiCh := make(chan dto.Cep, 1)

	go func() {
		if viaCepDelay > 0 {
			time.Sleep(viaCepDelay)
		}
		result, err := h.viaCepClient.GetCep(ctx, cep)
		if err != nil {
			log.Printf("[ViaCEP] erro ao consultar CEP %s: %v", cep, err)
			return
		}
		viaCepCh <- *result
	}()

	go func() {
		if brasilApiDelay > 0 {
			time.Sleep(brasilApiDelay)
		}
		result, err := h.brasilApiClient.GetCep(ctx, cep)
		if err != nil {
			log.Printf("[BrasilAPI] erro ao consultar CEP %s: %v", cep, err)
			return
		}
		brasilApiCh <- *result
	}()

	select {
	case result := <-viaCepCh:
		respondJSON(w, http.StatusOK, result)
	case result := <-brasilApiCh:
		respondJSON(w, http.StatusOK, result)
	case <-time.After(1 * time.Second):
		respondJSON(w, http.StatusRequestTimeout, map[string]string{"error": "request timeout"})
	}
}
