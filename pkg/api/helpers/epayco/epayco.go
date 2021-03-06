package epayco

import (
	"net/http"
	"strconv"
)

// Payment -> epayco payment data
type Payment struct {
	Ref              int          `json:"x_ref_payco"`
	Invoice          string       `json:"x_id_invoice"`
	Description      string       `json:"x_description"`
	Amount           int          `json:"x_amount"`
	AmountContry     int          `json:"x_amount_country"`
	AmountOk         int          `json:"x_amount_ok"`
	AmountBase       int          `json:"x_amount_base"`
	Tax              int          `json:"x_tax"`
	CurrencyCode     string       `json:"x_currency_code"`
	BankName         string       `json:"x_bank_name"`
	Cardnumber       string       `json:"x_cardnumber"`
	Quotas           interface{}  `json:"x_quotas"`
	Response         string       `json:"x_response"`
	ResponseCode     ResponseCode `json:"x_cod_response"`
	ApprovalCode     string       `json:"x_approval_code"`
	TransactionID    string       `json:"x_transaction_id"`
	TransactionDate  string       `json:"x_transaction_date"`
	TransactionState string       `json:"x_transaction_state"`
	Franchise        string       `json:"x_franchise"`
	Test             string       `json:"x_test_request"`
	Signature        string       `json:"x_signature"`
}

type Response struct {
	Success bool     `json:"success"`
	Data    *Payment `json:"data"`
}

// CheckoutArgs -> checkout information needed to start a redirect to epayco checkout
type CheckoutArgs struct {
	Key          string `json:"key"`
	Test         bool   `json:"test"`
	Response     string `json:"response"`
	Confirmation string `json:"confirmation"`
}

// ResponseCode payment response code
type ResponseCode int

const (
	Accepted  ResponseCode = 1
	Rejected  ResponseCode = 2
	Pending   ResponseCode = 3
	Failed    ResponseCode = 4
	Reversed  ResponseCode = 6
	Held      ResponseCode = 7
	Started   ResponseCode = 8
	Expired   ResponseCode = 9
	Abandoned ResponseCode = 10
	Canceled  ResponseCode = 11
	AntiFraud ResponseCode = 12
)

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

func ParsePaymentFromRequest(r *http.Request) *Payment {
	return &Payment{
		Ref:              atoi(r.FormValue("x_ref_payco")),
		Invoice:          r.FormValue("x_id_invoice"),
		Description:      r.FormValue("x_description"),
		Amount:           atoi(r.FormValue("x_amount")),
		AmountContry:     atoi(r.FormValue("x_amount_country")),
		AmountOk:         atoi(r.FormValue("x_amount_ok")),
		AmountBase:       atoi(r.FormValue("x_amount_base")),
		Tax:              atoi(r.FormValue("x_tax")),
		CurrencyCode:     r.FormValue("x_currency_code"),
		BankName:         r.FormValue("x_bank_name"),
		Cardnumber:       r.FormValue("x_cardnumber"),
		Quotas:           r.FormValue("x_quotas"),
		Response:         r.FormValue("x_response"),
		ResponseCode:     ResponseCode(atoi(r.FormValue("x_cod_response"))),
		ApprovalCode:     r.FormValue("x_approval_code"),
		TransactionID:    r.FormValue("x_transaction_id"),
		TransactionDate:  r.FormValue("x_transaction_date"),
		TransactionState: r.FormValue("x_transaction_state"),
		Franchise:        r.FormValue("x_franchise"),
		Test:             r.FormValue("x_test_request"),
		Signature:        r.FormValue("x_signature"),
	}
}
