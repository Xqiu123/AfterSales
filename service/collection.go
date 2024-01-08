package service

import (
	"as/dao"
	"as/util"
)

func CreateOrRemoveCollection(userId int, carId int) error {
	collection := &dao.CollectionModel{
		CreateTime: util.GetCurrentTime(),
		UserId:     userId,
		CarId:      carId,
	}

	isCollection, err := dao.IsUserCollectionCar(userId, carId)
	if err != nil {
		return err
	}

	if isCollection {
		return dao.DeleteCollection(collection)
	} else {
		return dao.CreateCollection(collection)
	}
}

func ListCollection(userId int) ([]*dao.CarModel, error) {
	carsIds, err := dao.ListCollectionByUserId(userId)
	if err != nil {
		return nil, err
	}

	return dao.ListCarsByIDs(carsIds)
}
