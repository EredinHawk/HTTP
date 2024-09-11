//
package get

// Image тип является контейнером для десериализации JSON http.Response в структуру.
type Image struct {
	// URL поле принимает значение ключа 'url' при десериализации JSON http.Response.Body.
	// HTTP ответ состоит из множества значений, но в данном примере нас интересует только ключ `url`.
	URL string
}