package service

import "as/dao"

// CreateApplication 添加新的Application
func CreateApplication(application *dao.ApplicationModel) error {
	return dao.CreateApplication(application)
}

// UpdateApplication 更新Application
func UpdateApplication(application *dao.ApplicationModel) error {
	return dao.UpdateApplication(application)
}

// DeleteApplication 删除Application
func DeleteApplication(id int) error {
	return dao.DeleteApplication(id)
}

// GetApplicationByID 通过ID查找Application
func GetApplicationByID(id int) (*dao.ApplicationModel, error) {
	return dao.GetApplicationByID(id)
}

// GetAllApplications 获取所有Application
func GetAllApplications() ([]*dao.ApplicationModel, error) {
	return dao.GetAllApplications()
}

// GetAllApplications 获取我的Application
func GetMyApplications(userId int) ([]*dao.ApplicationModel, error) {
	return dao.GetMyApplications(userId)
}
