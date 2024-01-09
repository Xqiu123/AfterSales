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

func GetAllApplications() ([]*ApplicationInfo, error) {
	var applications []*ApplicationInfo
	err := DB.Table("applications").
		Select("applications.id as id, users.user_name as user_name, users.telephone as telephone, cars.brand as brand").
		Joins("join users on users.id = applications.user_id").
		Joins("join cars on cars.id = applications.car_id").
		Scan(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}

type ApplicationInfo struct {
	Class       string `json:"class"`
	Description string `json:"description"`
	Time        string `json:"time"`
	IsEnd       int    `json:"is_end"`
	UserName    string `json:"user_name"`
	Telephone   string `json:"telephone"`
}

func GetMyApplications(userId int) ([]*ApplicationModel, error) {
	var applications []*ApplicationModel
	result := DB.Where("user_id = ?", userId).Find(&applications)
	if result.Error != nil {
		return nil, result.Error
	}
	return applications, nil
}
