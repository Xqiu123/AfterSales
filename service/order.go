package service

import (
	"as/dao"
)

func CreateOrder(userId int, carId int) error {
	return dao.CreateOrder(&dao.OrderModel{
		UserId: userId,
		CarId:  carId,
	})
}

func ListOrder(userId int) ([]*dao.CarModel, error) {
	carsIds, err := dao.ListOrderByUserId(userId)
	if err != nil {
		return nil, err
	}

	return dao.ListCarsByIDs(carsIds)
}
