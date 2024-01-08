package service

import (
	"as/dao"
	"as/pkg/constvar"
	"as/pkg/errno"
	"as/pkg/token"
	"time"
)

// GetInfo ... 获取用户信息
func GetInfo(ids []int) ([]*dao.UserModel, error) {
	return dao.GetUserByIds(ids)
}

func GetProfile(id int) (*dao.UserModel, error) {

	return dao.GetUser(id)
}

// List ... 获取用户列表
func List(role string, offset, limit, lastId int) ([]*dao.UserModel, error) {

	// 过滤条件
	filter := &dao.UserModel{Role: role}

	// DB 查询
	return dao.ListUser(offset, limit, lastId, filter)
}

// Login ... 登录
// 如果无 code，则返回 oauth 的地址，让前端去请求 oauth，
// 否则，用 code 获取 oauth 的 access token，并生成该应用的 auth token，返回给前端。
func Login(telephone string, password string) (string, error) {
	user, err := dao.GetUserByTelephone(telephone)

	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errno.ErrUserNotFound
		// // if err := dao.GetUserInfoFormOne(telephone, password); err != nil {
		// // 	return errno.ServerErr(errno.ErrRegister, err.Error())
		// // }
		// info := &dao.RegisterInfo{
		// 	Telephone: telephone,
		// 	Password:  password,
		// 	Role:      constvar.NormalRole,
		// 	Name:      name,
		// }
		// // 用户未注册，自动注册
		// if err := dao.RegisterUser(info); err != nil {
		// 	return "", err
		// }
		//
		// // 注册后重新查询
		// user, err = dao.GetUserByTelephone(telephone)
		// if err != nil {
		// 	return "", err
		// }
		//
		// if err := dao.AddRole("user", user.Id, constvar.NormalRole); err != nil {
		// 	return "", err
		// }
	} else {
		if !user.CheckPassword(password) {
			return "", errno.ErrPasswordIncorrect
		}
	}

	role := constvar.Normal
	if user.Role == constvar.AdminRole {
		role = constvar.Admin
	}

	// 生成 auth token
	Token, err := token.GenerateToken(&token.TokenPayload{
		Id:      uint32(user.Id),
		Role:    uint32(role),
		Expired: time.Hour * 24 * time.Duration(10),
	})
	if err != nil {
		return "", err
	}

	return Token, nil
}
func Register(telephone, name, password, sex string) error {
	return dao.RegisterUser(
		&dao.RegisterInfo{
			Telephone: telephone,
			Name:      name,
			Password:  password,
			Role:      "user",
			Sex:       sex,
		})
}

// UpdateInfo ... 更新用户信息
func UpdateInfo(id int, name string) error {
	user, err := dao.GetUser(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errno.ErrUserNotFound
	}

	user.Name = name

	if err := user.Update(); err != nil {
		return err
	}

	return nil
}
