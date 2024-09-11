package main

import (
	"log"
	"net/http"

	"github.com/EredinHawk/http/get"
	"github.com/EredinHawk/http/secret"
)

// init предварительно загружает переменные окружения в текущий процесс до main функции.
func init() {
	if err := secret.ScanEnv(); err != nil {
		log.Fatal(err)
	}
}

// main инициализирует HTTP клиент и запрос к API.
// В случае успешного выполнения запроса возвращается HTTP ответ.
// Из Body HTTP ответа десериализируется поле URL, в котором содержится строка url к ресурсу.
// Перейдя по адресу, получим изображение.
func main() {
	request, err := get.ConstructRequest()
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	get.PrintResponse(response)
}
