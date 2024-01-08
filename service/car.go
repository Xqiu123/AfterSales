package service

import "as/dao"

// CreateCar 添加新的Car
func CreateCar(car *dao.CarModel) error {
	return dao.CreateCar(car)
}

// UpdateCar 更新Car
func UpdateCar(car *dao.CarModel) error {
	return dao.UpdateCar(car)
}

// DeleteCar 删除Car
func DeleteCar(id int) error {
	return dao.DeleteCar(id)
}

// GetCarByID 通过ID查找Car
func GetCarByID(id int) (*dao.CarModel, error) {
	return dao.GetCarByID(id)
}

// GetAllCarsByBrand 获取特定品牌或所有品牌的Car
func GetAllCarsByBrand(brand string) ([]*dao.CarModel, error) {
	return dao.GetAllCarsByBrand(brand)
}

// SearchCarsByName 根据名称搜索Car
func SearchCarsByName(name string) ([]*dao.CarModel, error) {
	return dao.SearchCarsByName(name)
}
