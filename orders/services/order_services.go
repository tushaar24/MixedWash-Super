package services

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tushaar24/mixedWash-backend/config"
	"github.com/tushaar24/mixedWash-backend/orders/services/models"
	"log"
)

var client = config.GetSupabaseClient()

func FetchAllOrders() ([]models.OrderDTO, error) {

	// handle orders_temp

	const selectColumns = `*,profiles:user_id(username,mobile_number, email),delivery_time:time_slots!delivery_slot_id(label),pickup_time:time_slots!pickup_slot_id(label),addresses:address_id(address_line1,address_line2,city,state,house_building,area,postal_code,latitude,longitude),services:service_id(name)`
	var orders []models.OrderDTO

	_, err := client.
		From("orders").
		Select(selectColumns, "", false).
		ExecuteTo(&orders)

	if err != nil {
		log.Fatalf("query error: %v", err)
		return nil, err
	}

	return orders, nil
}

func CreateCustomer(customerCreationDTO models.CustomerCreationDTO) (string, error) {
	data, _, err := client.
		From("temp_customers").
		Insert(customerCreationDTO, false, "", "", "").
		Execute()

	if err != nil {
		log.Printf("query error: %v", err)
		return "0", err
	}

	// Define a struct to match the returned row (with at least the ID field)
	type InsertedCustomer struct {
		ID string `json:"id"`
	}

	var inserted []InsertedCustomer

	err = json.Unmarshal(data, &inserted)
	if err != nil {
		log.Printf("unmarshal error: %v", err)
		return "0", err
	}

	if len(inserted) == 0 {
		return "0", fmt.Errorf("no row returned after insert")
	}

	return inserted[0].ID, nil
}

func CreateOrderAdmin(order models.OrderCreationDTO) error {

	var userId []UserId

	_, err := client.
		From("temp_customers").
		Select("id", "", false).
		Eq("id", order.UserID.String()).
		ExecuteTo(&userId)

	if err != nil {
		log.Fatal("query error: ", err)
		return err
	}

	if len(userId) != 0 {

		_, _, err := client.
			From("orders_temp").
			Insert(order, false, "", "", "").
			Execute()

		if err != nil {
			log.Fatal("query error: ", err)
			return err
		}

	} else {

		_, _, err := client.
			From("orders").
			Insert(order, false, "", "", "").
			Execute()

		if err != nil {
			log.Fatal("query error: ", err)
			return err
		}

	}

	return nil
}

func GetAllOrderOfUser(userId uuid.UUID) ([]models.OrderDTO, error) {

	// Handle orders_temp

	var orders []models.OrderDTO

	_, err := client.
		From("orders").
		Select("*", "", false).
		Eq("user_id", userId.String()).
		ExecuteTo(&orders) // fills &orders, returns row-count

	if err != nil {
		log.Fatalf("query error: %v", err)
		return nil, err
	}

	return orders, nil
}

func GetCustomerByPhoneNo(phoneNumber string) (*models.CustomerByPhoneDTO, error) {

	var customerList []models.CustomerByPhoneDTO
	var tempCustomerList []models.TempCustomerByPhoneDTO

	_, err := client.
		From("profiles").
		Select("id, username", "", false).
		Eq("mobile_number", phoneNumber).
		ExecuteTo(&customerList)

	if err != nil {
		log.Fatalf("query error 1: %v", err)
		return nil, err
	}

	if len(customerList) == 0 {

		_, err1 := client.
			From("temp_customers").
			Select("id, customer_name", "", false).
			Eq("customer_phone_number", phoneNumber).
			ExecuteTo(&tempCustomerList)

		if err1 != nil {
			log.Fatalf("query error 2: %v", err1)
			return nil, err1
		}

		if len(tempCustomerList) != 0 {
			var tempCustomer = tempCustomerList[0]
			var cus = models.CustomerByPhoneDTO{
				CustomerId:   tempCustomer.CustomerID,
				CustomerName: tempCustomer.CustomerName,
			}
			return &cus, nil
		}

		return nil, nil

	}

	var customer = customerList[0]

	return &customer, nil
}

func GetCustomerAddresses(userId string) ([]models.CustomerAddressByUserIdDTO, error) {

	var addresses []models.CustomerAddressByUserIdDTO

	var userIds []UserId

	_, err := client.
		From("temp_customers").
		Select("id", "", false).
		Eq("id", userId).
		ExecuteTo(&userIds)

	if err != nil {
		log.Fatalf("query error: %v", err)
		return nil, err
	}

	if len(userIds) != 0 {

		_, err := client.
			From("addresses_temp").
			Select("id, address_line1, address_line2, city, state, house_building, area, postal_code", "", false).
			Eq("user_id", userId).
			ExecuteTo(&addresses)

		if err != nil {
			log.Fatalf("query error: %v", err)
			return nil, err
		}


	} else {
		_, err := client.
			From("addresses").
			Select("id, address_line1, address_line2, city, state, house_building, area, postal_code", "", false).
			Eq("user_id", userId).
			ExecuteTo(&addresses)

		if err != nil {
			log.Fatalf("query error: %v", err)
			return nil, err
		}
	}

	return addresses, nil

}

func AddAddressAdmin(address models.AddAddressAdminDTO) error {

	_, _, err := client.
		From("addresses_temp").
		Insert(address, false, "", "", "").
		Execute()

	if err != nil {
		log.Fatalf("query error: %v", err)
		return err
	}

	return nil
}

func GetAdminOrderCreationScreen() (*models.OrderCreationScreenDTO, error) {

	var services []models.OrderCreationScreenServicesDTO

	_, servicesError := client.
		From("services").
		Select("id, name", "", false).
		ExecuteTo(&services)

	if servicesError != nil {
		log.Fatalf("services query error: %v", servicesError)
		return nil, servicesError
	}

	var pickupTimeSlot []models.OrderCreationScreenPickupSlotDTO
	var deliveryTimeSlot []models.OrderCreationScreenDeliverySlotDTO

	_, pickupTimeSlotError := client.
		From("time_slots").
		Select("id, label", "", false).
		ExecuteTo(&pickupTimeSlot)

	if pickupTimeSlotError != nil {
		log.Fatalf("pickupslot query error: %v", pickupTimeSlotError)
		return nil, pickupTimeSlotError
	}

	_, deliveryTimeSlotError := client.
		From("time_slots").
		Select("id, label", "", false).
		ExecuteTo(&deliveryTimeSlot)

	if deliveryTimeSlotError != nil {
		log.Fatalf("deliveryslot query error: %v", deliveryTimeSlotError)
		return nil, deliveryTimeSlotError
	}

	return &models.OrderCreationScreenDTO{
		Services:     services,
		PickupSlot:   pickupTimeSlot,
		DeliverySlot: deliveryTimeSlot,
	}, nil
}

type UserId struct {
	UserId string `json:"id"`
}
