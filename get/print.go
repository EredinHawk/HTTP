package get

import (
	"io"
	"net/http"
)

// GetURL возвращает значение ключа URL из тела HTTP ответа.
func GetURL(response *http.Response) ([]byte, error) {
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
