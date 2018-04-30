package lib

//Block_req структура запроса для авторизации платежа
type Block_req struct {
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

//Block_resp структура ответа на запрос для авторизации платежа
type Block_resp struct {
	dealId int      `json:"deal_id"`
	error  []string `json:"error"`
}

//Charge_req структура запроса платежа
type Charge_req struct {
	dealId int `json:"deal_id"`
	amount int `json:"amount"`
}

//Charge_resp структура ответа платежа
type Charge_resp struct {
	status string   `json:"status"`
	error  []string `json:"error"`
}
