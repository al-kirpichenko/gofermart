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

	client := &http.Client{}

	var loyalty Loyalty

	// Создаем новый запрос GET к внешнему сервису

	req, err := http.NewRequest("GET", accrualAddress+"/api/orders/"+orderNumber, nil)
	req.Close = true

	if err != nil {
		return nil, err
	}

	var resp *http.Response

	// делаем до 3 попыток с интервалом 10 сек. в случае неудачи
	for i := 0; i < 3; i++ {
		resp, err = client.Do(req)
		if err == nil {
			break
		}
		duration := time.Second * 10
		time.Sleep(duration)
	}

	//resp, err := client.Do(req)

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
