package lib

//BlockRequest структура запроса для авторизации платежа
type BlockRequest struct {
	MerchantContactID int `json:"merchant_contact_id"`
	Card              struct {
		PAN    string `json:"pan"`
		EMonth int    `json:"e_month"`
		EYear  int    `json:"e_year"`
		CVV    int    `json:"cvv"`
		Holder string `json:"holder"`
	}
	OrderID string `json:"order_id"`
	Amount  int    `json:"amount"`
}

//BlockResponse структура ответа на запрос для авторизации платежа
type BlockResponse struct {
	DealID int      `json:"deal_id"`
	Amount int      `json:"amount"`
	Error  []string `json:"error"`
}

//ChargeRequest структура запроса платежа
type ChargeRequest struct {
	DealID int `json:"deal_id"`
	Amount int `json:"amount"`
}

//ChargeResponse структура ответа платежа
type ChargeResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
