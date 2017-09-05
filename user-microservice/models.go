package main

type BillingAddress struct {
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zipcode     string `json:"zipcode"`
}

type User struct {
	ID      uint           `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Address BillingAddress `json:"address"`
}
