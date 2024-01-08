package dao

type CarModel struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Brand     string `json:"brand" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`
	Image     string `json:"image" gorm:"default:null"`
	Price     string `json:"price" gorm:"default:null"`
	Color     string `json:"color" gorm:"default:null"`
	Size      string `json:"size" gorm:"default:null"`
	Motor     string `json:"motor" gorm:"default:null"`
	Gearbox   string `json:"gearbox" gorm:"default:null"`
	MaxHorse  string `json:"max_horse" gorm:"default:null"`
	MaxPower  string `json:"max_power" gorm:"default:null"`
	DriveMode string `json:"drive_mode" gorm:"default:null"`
	Fuel      string `json:"fuel" gorm:"default:null"`
}

func (*CarModel) TableName() string {
	return "cars"
}

// CreateCar 添加新的Car
func CreateCar(car *CarModel) error {
	return DB.Create(car).Error
}

// UpdateCar 更新Car
func UpdateCar(car *CarModel) error {
	return DB.Save(car).Error
}

// DeleteCar 删除Car
func DeleteCar(id int) error {
	return DB.Delete(&CarModel{}, id).Error
}

// GetCarByID 通过ID查找Car
func GetCarByID(id int) (*CarModel, error) {
	car := &CarModel{}
	result := DB.First(car, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return car, nil
}

// GetAllCarsByBrand 获取特定品牌或所有品牌的Car
func GetAllCarsByBrand(brand string) ([]*CarModel, error) {
	var cars []*CarModel
	query := DB.Model(&CarModel{})
	if brand != "" {
		query = query.Where("brand = ?", brand)
	}
	result := query.Find(&cars)
	if result.Error != nil {
		return nil, result.Error
	}
	return cars, nil
}

// SearchCarsByName 根据名称搜索Car
func SearchCarsByName(name string) ([]*CarModel, error) {
	var cars []*CarModel
	result := DB.Where("name LIKE ?", "%"+name+"%").Find(&cars)
	if result.Error != nil {
		return nil, result.Error
	}
	return cars, nil
}

// ListCarsByIDs 根据一组ID获取Car
func ListCarsByIDs(ids []int) ([]*CarModel, error) {
	var cars []*CarModel
	result := DB.Find(&cars, ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return cars, nil
}
