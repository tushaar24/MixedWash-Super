package services

import (
	"github.com/tushaar24/mixedWash-backend/config"
	"github.com/tushaar24/mixedWash-backend/drivers/models"
	serviceModels "github.com/tushaar24/mixedWash-backend/orders/services/models"
	"github.com/tushaar24/mixedWash-backend/utils"
	"log"
	"slices"
	"time"
)

var client = config.GetSupabaseClient()

func getAllTasks() ([]models.DriverTasksDTO, error) {

	var tasks []models.DriverTasksDTO

	_, err := client.
		From(utils.DRIVER_TASK_TABLE).
		Select("*", "", false).
		ExecuteTo(&tasks)

	if err != nil {
		log.Fatalf("query error: %v", err)
		return nil, err
	}

	return tasks, nil
}

func GetTodayTask() ([]models.DriverTasksDTO, error) {

	var tasks []models.CreateDriverTaskDTO
	var currentTasks, currentTasksError = getAllTasks()
	var todaysTask []models.DriverTasksDTO

	loc, _ := time.LoadLocation("Asia/Kolkata")
	todayStr := time.Now().In(loc).Format("2006-01-02")

	if currentTasksError != nil {
		log.Fatalf("query error: %v", currentTasksError)
		return nil, currentTasksError
	}

	var todayPickupOrders []serviceModels.OrderTaskDTO
	var todayDeliveryOrders []serviceModels.OrderTaskDTO
	var todayPickupTempOrders []serviceModels.TempOrderTaskDTO
	var todayDeliveryTempOrders []serviceModels.TempOrderTaskDTO

	_, todayPickupOrderError := client.
		From(utils.ORDERS_TABLE).
		Select("*", "", false).
		Eq("pickup_date", todayStr).
		ExecuteTo(&todayPickupOrders)

	if todayPickupOrderError != nil {
		log.Fatalf("query error: %v", todayPickupOrderError)
		return nil, todayPickupOrderError
	}

	_, todayDeliveryOrderError := client.
		From(utils.ORDERS_TABLE).
		Select("*", "", false).
		Eq("delivery_date", todayStr).
		ExecuteTo(&todayDeliveryOrders)

	if todayDeliveryOrderError != nil {
		log.Fatalf("query error: %v", todayDeliveryOrderError)
		return nil, todayDeliveryOrderError
	}

	_, todayPickupTempOrderError := client.
		From(utils.TEMP_ORDER_TABLE).
		Select("*", "", false).
		Eq("pickup_date", todayStr).
		ExecuteTo(&todayPickupTempOrders)

	if todayPickupTempOrderError != nil {
		log.Fatalf("query error: %v", todayPickupTempOrderError)
		return nil, todayPickupTempOrderError
	}

	_, todayDeliveryTempOrderError := client.
		From(utils.TEMP_ORDER_TABLE).
		Select("*", "", false).
		Eq("delivery_date", todayStr).
		ExecuteTo(&todayDeliveryTempOrders)

	if todayDeliveryTempOrderError != nil {
		log.Fatalf("query error: %v", todayDeliveryTempOrderError)
		return nil, todayDeliveryTempOrderError
	}

	for _, todayPickupOrder := range todayPickupOrders {

		found := slices.IndexFunc(currentTasks, func(c models.DriverTasksDTO) bool {
			return c.OrderId == todayPickupOrder.ID
		}) >= 0

		if !found {
			tasks = append(tasks,
				models.CreateDriverTaskDTO{
					OrderId:        &todayPickupOrder.ID,
					TempOrderId:    nil,
					CustomerId:     &todayPickupOrder.UserID,
					TempCustomerId: nil,
					DriverId:       nil,
					Status:         "pending",
					TypeTask:       "pickup",
				})
		}
	}

	for _, todayDeliveryOrder := range todayDeliveryOrders {

		found := slices.IndexFunc(currentTasks, func(c models.DriverTasksDTO) bool {
			return c.OrderId == todayDeliveryOrder.ID
		}) >= 0

		if !found {
			tasks = append(tasks,
				models.CreateDriverTaskDTO{
					OrderId:        &todayDeliveryOrder.ID,
					TempOrderId:    nil,
					CustomerId:     &todayDeliveryOrder.UserID,
					TempCustomerId: nil,
					DriverId:       nil,
					Status:         "pending",
					TypeTask:       "delivery",
				})
		}
	}

	for _, todayPickupTempOrder := range todayPickupTempOrders {

		found := slices.IndexFunc(currentTasks, func(c models.DriverTasksDTO) bool {
			return c.TempOrderId == todayPickupTempOrder.ID
		}) >= 0

		if !found {
			tasks = append(tasks,
				models.CreateDriverTaskDTO{
					OrderId:        nil,
					TempOrderId:    &todayPickupTempOrder.ID,
					CustomerId:     nil,
					TempCustomerId: &todayPickupTempOrder.UserID,
					DriverId:       nil,
					Status:         "pending",
					TypeTask:       "pickup",
				})
		}
	}

	for _, todayDeliveryTempOrder := range todayDeliveryTempOrders {

		found := slices.IndexFunc(currentTasks, func(c models.DriverTasksDTO) bool {
			return c.TempOrderId == todayDeliveryTempOrder.ID
		}) >= 0

		if !found {
			tasks = append(tasks,
				models.CreateDriverTaskDTO{
					OrderId:        nil,
					TempOrderId:    &todayDeliveryTempOrder.ID,
					CustomerId:     nil,
					TempCustomerId: &todayDeliveryTempOrder.UserID,
					DriverId:       nil,
					Status:         "pending",
					TypeTask:       "delivery",
				})
		}
	}

	_, _, insertTasksError := client.
		From(utils.DRIVER_TASK_TABLE).
		Insert(tasks, false, "", "", "").
		Execute()

	if insertTasksError != nil {
		log.Fatalf("query error: %v", insertTasksError)
		return nil, insertTasksError
	}

	_, todayTaskError := client.
		From(utils.DRIVER_TASK_TABLE).
		Select("*", "", false).
		ExecuteTo(&todaysTask)

	if todayTaskError != nil {
		log.Fatalf("query error: %v", insertTasksError)
		return nil, insertTasksError
	}

	return todaysTask, nil

}

func GetDrivers() ([]models.DriverDTO, error) {

	var drivers []models.DriverDTO

	_, err := client.
		From(utils.DRIVER_TABLE).
		Select("id, name, phone_number, salary", "", false).
		ExecuteTo(&drivers)

	if err != nil {
		log.Fatalf("query error: %v", err)
		return nil, err
	}

	return drivers, nil
}

func UpdateDriver(driverId string, taskId string) error {

	_, _, err := client.
		From(utils.DRIVER_TASK_TABLE).
		Update(map[string]interface{}{
			"driver_id": driverId,
		}, "minimal", "").
		Eq("id", taskId).
		Execute()

	if err != nil {
		log.Fatalf("query error: %v", err)
		return err
	}

	return nil
}

func UpdateDriverTaskStatus(taskId string, status string) error {

	_, _,err := client.
	From(utils.DRIVER_TASK_TABLE).
	Update(map[string]interface{}{
			"status": status,
		}, "minimal", "").
	Eq("id", taskId).
	Execute()

	if err != nil {
		log.Fatalf("query error: %v", err)
		return err
	}

	return nil
}

