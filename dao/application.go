package dao

type ApplicationModel struct {
	ID          int    `json:"id" gorm:"primary_key"`
	UserID      int    `json:"user_id" gorm:"not null"`
	CarID       int    `json:"car_id" gorm:"not null"`
	Class       string `json:"class" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Time        string `json:"time" gorm:"not null"`
	IsEnd       int    `json:"is_end" gorm:"not null"`
}

func (*ApplicationModel) TableName() string {
	return "applications"
}

// CreateApplication 添加新的Application
func CreateApplication(application *ApplicationModel) error {
	return DB.Create(application).Error
}

// UpdateApplication 更新Application
func UpdateApplication(application *ApplicationModel) error {
	return DB.Save(application).Error
}

// DeleteApplication 删除Application
func DeleteApplication(id int) error {
	return DB.Delete(&ApplicationModel{}, id).Error
}

// GetApplicationByID 通过ID查找Application
func GetApplicationByID(id int) (*ApplicationModel, error) {
	application := &ApplicationModel{}
	result := DB.First(application, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return application, nil
}

// GetAllApplications 获取所有Application
func GetAllApplications() ([]*ApplicationModel, error) {
	var applications []*ApplicationModel
	result := DB.Find(&applications)
	if result.Error != nil {
		return nil, result.Error
	}
	return applications, nil
}

// GetAllApplications 获取我的Application
func GetMyApplications(userId int) ([]*ApplicationModel, error) {
	var applications []*ApplicationModel
	result := DB.Where("user_id = ?", userId).Find(&applications)
	if result.Error != nil {
		return nil, result.Error
	}
	return applications, nil
}
