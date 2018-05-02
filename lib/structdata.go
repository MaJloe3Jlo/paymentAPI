package lib

//Block_req структура запроса для авторизации платежа
type Block_req struct {
	MerchantContactId int `json:"merchant_contact_id"`
	Card              struct {
		Pan    string `json:"pan"`
		EMonth int    `json:"e_month"`
		EYear  int    `json:"e_year"`
		Cvv    int    `json:"cvv"`
		Holder string `json:"holder"`
	}
	OrderId string `json:"order_id"`
	Amount  int    `json:"amount"`
}

//Block_resp структура ответа на запрос для авторизации платежа
type Block_resp struct {
	DealId int      `json:"deal_id"`
	Amount int      `json:"amount"`
	Error  []string `json:"error"`
}

//Charge_req структура запроса платежа
type Charge_req struct {
	DealId int `json:"deal_id"`
	Amount int `json:"amount"`
}

//Charge_resp структура ответа платежа
type Charge_resp struct {
	Status string   `json:"status"`
	Error  string `json:"error"`
}
