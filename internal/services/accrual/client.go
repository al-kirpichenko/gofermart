package accrual

import (
	"encoding/json"
	"net/http"
)

type Loyalty struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float32 `json:"accrual"`
}

func GetLoyalty(orderNumber string, accrualAddress string) (*Loyalty, error) {

	client := &http.Client{}

	var loyalty Loyalty

	// Создаем новый запрос GET к внешнему сервису

	req, err := http.NewRequest("GET", accrualAddress+"/api/orders/"+orderNumber, nil)

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
