package secret

import (
	"fmt"
	"os"

	env "github.com/joho/godotenv"
)

// ScanEnv возвращает nil, если загрузка переменных окружения
// из файла /.env завершилась успешно. В противном случае вернет ошибку типа error.
//
// Для успешной работы, треубется клонировать репозиторий и в его корне создать
// файл .env и внести в него строку `API_KEY={x}`, где x - персональный API ключ.
// Ключ можно получить по ссылке на сайте https://api.nasa.gov/ после регистрации.
func ScanEnv() error {
	if err := env.Load("./.env"); err != nil {
		return fmt.Errorf("ошибка загрузки /.env файла - %v", err)
	}
	return nil
}

// GetKeyAPI возвращает значение переменной окружения API_KEY из .env файла в корне репозитория.
func GetKeyAPI(keyName string) string {
	key := os.Getenv(keyName)
	return key
}
