package shared

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
}

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Count       int    `json:"count"`
	Category    string `json:"category"`
}

type DbItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Count       int    `json:"count"`
	Category    string `json:"category"`
}
