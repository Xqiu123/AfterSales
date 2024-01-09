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
func GetAllApplications() ([]*dao.ApplicationInfo, []*dao.ApplicationInfo, error) {
	applications, err := dao.GetAllApplications()
	if err != nil {
		return nil, nil, err
	}

	driveApplications := make([]*dao.ApplicationInfo, 0)
	repairApplications := make([]*dao.ApplicationInfo, 0)
	for _, application := range applications {
		if application.Class == "drive" {
			driveApplications = append(driveApplications, application)
		} else {
			repairApplications = append(repairApplications, application)
		}
	}

	return driveApplications, repairApplications, nil
}

func GetApplicationsByUserId(userId int) ([]*dao.ApplicationModel, error) {
	return dao.GetMyApplications(userId)
}
