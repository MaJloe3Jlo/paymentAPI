package lib

//Block структура запроса для авторизации платежа
type Block struct {
	merchantContactId int `json:"merchant_contact_id"`
	card              struct {
		pan    string `json:"card>pan"`
		eMonth int    `json:"card>e_month"`
		eYear  int    `json:"card>e_year"`
		cvv    int    `json:"card>cvv"`
		holder string `json:"card>holder"`
	} `json:"card"`
	orderId string `json:"order_id"`
	amount  int    `json:"amount"`
}

//Charge структура запроса платежа
type Charge struct {
	dealId int `json:"deal_id"`
	amount int `json:"amount"`
}

