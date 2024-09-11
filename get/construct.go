package get

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/EredinHawk/http/secret"
)

// ConstructRequest возвращает инициализированный тип *http.Request,
// который будет использован при отправке запроса к конечной точке API
func ConstructRequest() (*http.Request, error) {
	request := &http.Request{
		Method: http.MethodGet,
		Header: map[string][]string{
			"Accept": {"application/json"},
		},
		URL: &url.URL{
			Scheme:   "https",
			Host:     "api.nasa.gov",
			Path:     "/planetary/apod",
			RawQuery: fmt.Sprintf("api_key=%v", secret.GetKeyAPI("API_KEY")),
		},
	}

	return request, nil
}
