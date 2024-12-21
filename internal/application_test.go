package application

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	// Создание фейкового запроса
	requestBody := []byte(`{"expression": "2+2"}`)
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Создание ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызов функции calculateHandler с фейковым запросом и ResponseRecorder
	calculateHandler(rr, req)

	// Проверка кода ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверка содержимого ответа
	expected := `{"result":"4.000000"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCalculateHandlerInvalidRequest(t *testing.T) {
	// Создание фейкового запроса с неверным JSON
	requestBody := []byte(`{"exp: "2+2"}`)

	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Создание ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызов функции calculateHandler с фейковым запросом и ResponseRecorder
	calculateHandler(rr, req)

	// Проверка кода ответа
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid request: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCalculateHandlerInvalidMethod(t *testing.T) {
	// Создание фейкового GET запроса
	req, err := http.NewRequest("GET", "/api/v1/calculate", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Создание ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызов функции calculateHandler с фейковым запросом и ResponseRecorder
	calculateHandler(rr, req)

	// Проверка кода ответа
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for invalid method: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestCalculateHandlerInternalServerError(t *testing.T) {
	// Создание фейкового запроса с недопустимым выражением
	requestBody := []byte(`{"expression": "2+$"}`)
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Создание ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызов функции calculateHandler с фейковым запросом и ResponseRecorder
	calculateHandler(rr, req)

	// Проверка кода ответа
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code for internal server error: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestCalculateHandlerUnprocessableEntity(t *testing.T) {
	// Создание фейкового запроса с неправильным выражением
	requestBody := []byte(`{"expression": "2++2"}`)
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Создание ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызов функции calculateHandler с фейковым запросом и ResponseRecorder
	calculateHandler(rr, req)

	// Проверка кода ответа
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code for unprocessable entity: got %v want %v",
			status, http.StatusUnprocessableEntity)
	}
}
