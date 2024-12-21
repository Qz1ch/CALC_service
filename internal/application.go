package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	calculator "github.com/qz1ch/calc_service/pkg"
)

// внесение порта как индивидуальную переменную
type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Request struct {
	Expression string `json:"expression"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	decoder := json.NewDecoder(r.Body)

	//распаковка и обработка ошибки с кодом 400
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//обработка ошибки с кодом 405
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//обработка ошибки с кодом 500
	if strings.Contains(req.Expression, "$") {
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
		return
	}

	result, err := calculator.Calc(req.Expression)
	var resp Response

	//обработка ошибок с кодом 422
	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		resp.Result = fmt.Sprintf("%f", result)
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Запуск сервера
func (a *Application) RunServer() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	log.Fatal(http.ListenAndServe(":"+a.config.Addr, nil))
}
