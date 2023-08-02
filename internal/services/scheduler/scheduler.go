package scheduler

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/al-kirpichenko/gofermart/cmd/gophermart/config"
	"github.com/al-kirpichenko/gofermart/internal/api"
	"github.com/al-kirpichenko/gofermart/internal/models"
	"github.com/al-kirpichenko/gofermart/internal/services/accrual"
)

func UpdateOrders(s *api.Server) {

	var orders []models.Order

	for {

		s.DB.Where("status IN ?", []string{config.StatusNew, config.StatusRegistered, config.StatusProcessing}).Find(&orders).Limit(config.OrdersBatchNum)

		for _, order := range orders {

			loyalty, err := accrual.Get(order.Number, s.Config.ServiceAddress)

			if err != nil {
				s.Logger.Error("No response from the accrual service", zap.Error(err))
				continue
			}

			var user models.User

			s.DB.Transaction(func(tx *gorm.DB) error {

				s.DB.First(&user, "id = ?", order.UserID)

				order.Accrual = loyalty.Accrual
				order.Status = loyalty.Status

				user.Balance = user.Balance + loyalty.Accrual

				if err := tx.Save(&order).Error; err != nil {
					s.Logger.Error("Don't save order accrual", zap.Error(err))
					return err
				}
				if err := tx.Save(&user).Error; err != nil {
					s.Logger.Error("Don't update user balance", zap.Error(err))
					return err
				}
				return nil
			})
		}
		time.Sleep(config.UpdateDuration)

	}
}
