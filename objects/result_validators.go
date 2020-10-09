package objects

type ResultValidator struct {
	Height string       `json:"height"`
	Result []*Validator `json:"result"`
}
