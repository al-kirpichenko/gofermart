package models

type Withdraw struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}
