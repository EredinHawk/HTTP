package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// PrintResponse печатает значение ключа URL из теле HTTP ответа в консоль.
func PrintResponse(response *http.Response) error {
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	i := Image{}
	json.Unmarshal(result, &i)
	fmt.Println(i.URL)
	return nil
}
