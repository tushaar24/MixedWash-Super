package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/config"
	"github.com/tushaar24/mixedWash-backend/orders/services/models"
)

var client = config.GetSupabaseClient()

func FetchAllOrders() ([]models.OrderDTO, error){
	var orders []models.OrderDTO
	_, err := client.
		From("orders").
		Select("*", "", false).
		ExecuteTo(&orders)   // fills &orders, returns row-count
	if err != nil {
		log.Fatalf("query error: %v", err)
		return nil, err
	}
	return orders, nil
}

func GetAllOrderOfUser(userId uuid.UUID) ([]models.OrderDTO, error){
	var orders []models.OrderDTO
	_, err := client.
		From("orders").
		Select("*", "", false).
		Eq("user_id", userId.String()).
		ExecuteTo(&orders)   // fills &orders, returns row-count
	if err != nil {
		log.Fatalf("query error: %v", err)
		return nil, err
	}
	return orders, nil
}

