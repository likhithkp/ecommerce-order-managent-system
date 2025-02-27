package shared

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       any    `json:"data"`
}

type Order struct {
	Products []OrderedProduct `json:"products"`
	Status   string           `json:"status"`
	Payment  PaymentDetails   `json:"paymentDetails"`
}

type OrderedProduct struct {
	ProdID   string `json:"prodId"`
	Quantity int    `json:"quantity"`
}

type PaymentDetails struct {
	PaymentMethod string  `json:"paymentMethod"`
	TotalAmount   float64 `json:"totalAmount"`
}
