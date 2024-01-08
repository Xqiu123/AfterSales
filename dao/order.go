package dao

type OrderModel struct {
	Id         int
	UserId     int
	CreateTime string
	CarId      int
}

func (OrderModel) TableName() string {
	return "orders"
}

// Create ...
func (c *OrderModel) Create() error {
	return DB.Create(c).Error
}

func (c *OrderModel) Delete() error {
	return DB.Delete(c).Error
}

func CreateOrder(order *OrderModel) error {
	err := order.Create()
	return err
}

func DeleteOrder(order *OrderModel) error {
	return DB.Table("orders").Delete(order).Error
}

func ListOrderByUserId(userId int) ([]int, error) {
	var carIds []int
	query := DB.Model(&OrderModel{})
	if userId != 0 {
		query = query.Where("user_id = ?", userId)
	}

	err := DB.Select("car_id").Find(&carIds).Error
	return carIds, err
}
