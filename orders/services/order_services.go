package services

import (
	"log"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/config"
	"github.com/tushaar24/mixedWash-backend/orders/services/models"
)

var client = config.GetSupabaseClient()

func FetchAllOrders() ([]models.OrderDTO, error){
	const selectColumns = `*,profiles:user_id(username,mobile_number),addresses:address_id(address_line1,address_line2,city,state,house_building,area,postal_code,latitude,longitude),services:service_id(name)`
	var orders []models.OrderDTO
	_, err := client.
		From("orders").
		Select(selectColumns, "", false).
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

