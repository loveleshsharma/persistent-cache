package request

type SetRequest struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Expiry int64  `json:"expiryInMinutes"`
}
