package dao

import (
	"as/pkg/constvar"
	"errors"
	"github.com/ShiinaOrez/GoSecurity/security"
	"gorm.io/gorm"
)

type UserModel struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Telephone string `json:"telephone" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`
	Sex       string `json:"sex" gorm:"default:null"`
	Role      string `json:"role" gorm:"column:role;"`
}

func (*UserModel) TableName() string {
	return "users"
}

// Create ...
func (u *UserModel) Create() error {
	return DB.Create(u).Error
}

// Save ...
func (u *UserModel) Save() error {
	return DB.Save(u).Error
}

func (u *UserModel) Update() error {
	err := DB.Table("users").Where("id = ?", u.Id).Updates(map[string]interface{}{
		"name": u.Name,
	}).Error

	return err
}

// generatePasswordHash pwd -> hashPwd
func generatePasswordHash(password string) string {
	return security.GeneratePasswordHash(password)
}

func (u *UserModel) CheckPassword(password string) bool {
	return security.CheckPasswordHash(password, u.Password)
}

type RegisterInfo struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Sex       string `json:"sex"`
	Role      string `json:"role"`
}

// GetUser get a single user by id
func GetUser(id int) (*UserModel, error) {
	user := &UserModel{}
	err := DB.Where("id = ?", id).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

// GetUserByIds get user by id array
func GetUserByIds(ids []int) ([]*UserModel, error) {
	var list []*UserModel
	if err := DB.Where("id IN (?)", ids).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

// ListUser list users
func ListUser(offset, limit, lastId int, filter *UserModel) ([]*UserModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	var list []*UserModel

	query := DB.Model(&UserModel{}).Where(filter).Offset(offset).Limit(limit)

	if lastId != 0 {
		query = query.Where("id < ?", lastId).Order("id DESC")
	}

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

// GetUserByTelephone get a user by telephone.
func GetUserByTelephone(telephone string) (*UserModel, error) {
	u := &UserModel{}
	err := DB.Where("telephone = ?", telephone).First(u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return u, err
}

func RegisterUser(info *RegisterInfo) error {
	user := &UserModel{
		Name:      info.Name,
		Telephone: info.Telephone,
		Password:  generatePasswordHash(info.Password),
		Role:      info.Role,
	}

	return user.Create()
}
