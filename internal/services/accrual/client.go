package accrual

import (
	"encoding/json"
	"net/http"
	"time"
)

type Loyalty struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

func GetLoyalty(orderNumber string, accrualAddress string) (*Loyalty, error) {

	var (
		loyalty *Loyalty
		err     error
	)

	for i := 0; i < 3; i++ {
		loyalty, err = get(orderNumber, accrualAddress)
		if err == nil {
			return loyalty, nil
		}
		duration := time.Second * 10
		time.Sleep(duration)
	}
	return nil, err
}

func get(orderNumber string, accrualAddress string) (*Loyalty, error) {

	client := &http.Client{}

	var loyalty Loyalty

	req, err := http.NewRequest("GET", accrualAddress+"/api/orders/"+orderNumber, nil)
	req.Close = true

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&loyalty)
	if err != nil {

		return nil, err
	}

	return &loyalty, nil
}
