package customers

import "github.com/alactic/mygoproject/models/creditcard"

type Customer struct {
	Id          string                  `json:"id, omitempty"`
	Type        string                  `json:"type"`
	Firstname   string                  `json:"firstname"`
	Lastname    string                  `json:"lastname"`
	CreditCards []creditcard.CreditCard `json:"creditcards"`
}
