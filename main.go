package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/EredinHawk/http/get"
	"github.com/EredinHawk/http/secret"
)

// init предварительно загружает переменные окружения в текущий процесс до main функции.
func init() {
	if err := secret.ScanEnv(); err != nil {
		log.Fatal(err)
	}
}

// main запускает http сервер с обработчиком, который выполняет HTTP запрос на сервер NASA
// и возвращает результат в виде URL на изображение.
func main() {
	server := constructServer()

	fmt.Println("Сервер localhost:8080 запущен и прослушивает входящие запросы...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("ошибка при запуске сервера (%v)", err)
	}
}

// constructServer инициализирует http сервер и эндпоинты
func constructServer() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("GET /", handler_API)

	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router, // HTTP мультиплексор, или по другому роутер
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return srv
}

func handler_API(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	request, err := get.ConstructRequest()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer response.Body.Close()

	result, err := get.GetURL(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, "%s", result)
}
