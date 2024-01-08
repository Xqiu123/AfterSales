package dao

type CollectionModel struct {
	Id         int
	UserId     int
	CreateTime string
	CarId      int
}

func (CollectionModel) TableName() string {
	return "collections"
}

// Create ...
func (c *CollectionModel) Create() error {
	return DB.Create(c).Error
}

func (c *CollectionModel) Delete() error {
	return DB.Delete(c).Error
}

func CreateCollection(collection *CollectionModel) error {
	err := collection.Create()
	return err
}

func DeleteCollection(collection *CollectionModel) error {
	return DB.Table("collections").Where("car_id = ? AND user_id = ?", collection.CarId, collection.UserId).Delete(collection).Error
}

func ListCollectionByUserId(userId int) ([]int, error) {
	var carIds []int
	err := DB.Table("collections").Select("car_id").Where("user_id = ?", userId).Find(&carIds).Error
	return carIds, err
}

func IsUserCollectionCar(userId int, carId int) (bool, error) {
	var count int64
	if err := DB.Table("collections").Where("user_id = ? AND car_id = ?", userId, carId).Count(&count).Error; err != nil {
		return false, err
	}

	return count != 0, nil
}
