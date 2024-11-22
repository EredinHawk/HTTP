package get

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetURL возвращает значение ключа URL из тела HTTP ответа.
func GetURL(response *http.Response) (string, error) {
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	i := Image{}
	json.Unmarshal(result, &i)
	return i.URL, nil
}
