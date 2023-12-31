package accrual

import (
	"encoding/json"
	"net/http"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
)

type Loyalty struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

func Get(orderNumber string, accrualAddress string) (*Loyalty, error) {

	client := &http.Client{}

	client.Timeout = config.ClientTimeout

	var loyalty Loyalty

	req, err := http.NewRequest("GET", accrualAddress+"/api/orders/"+orderNumber, nil)
	req.Close = true

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&loyalty)
	if err != nil {

		return nil, err
	}

	return &loyalty, nil
}
