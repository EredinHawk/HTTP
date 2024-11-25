package main

import (
	"fmt"
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

// main запускает http сервер с обработчиком, который выполняет HTTP запрос на сервер NASA
// и возвращает результат в виде URL на изображение.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		
		result, err := get.GetURL(response)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", result)
	})
	
	fmt.Println("Сервер прослушивает входящие запросы.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
