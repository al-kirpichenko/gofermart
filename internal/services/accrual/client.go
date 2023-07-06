package accrual

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Loyalty struct {
	Order   string `json:"order"`
	Status  string `json:"status"`
	Accrual int    `json:"accrual"`
}

func GetLoyalty(orderNumber int, accrualAddress string) (*Loyalty, error) {

	client := &http.Client{}

	var loyalty Loyalty

	// Создаем новый запрос GET к внешнему сервису

	req, err := http.NewRequest("GET", accrualAddress+"/"+fmt.Sprintf("%v", orderNumber), nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&loyalty)
	if err != nil {

		return nil, err
	}

	return &loyalty, nil
}