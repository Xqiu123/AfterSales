package user

type RegisterRequest struct {
	Telephone string `json:"telephone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Sex       string `json:"sex"`
}

// LoginRequest 请求
type LoginRequest struct {
	Telephone string `json:"telephone" binding:"required"`
	Password  string `json:"password" binding:"required"`
} // @name LoginRequest

// LoginResponse login 请求响应
type LoginResponse struct {
	Token string `json:"token"`
} // @name LoginResponse

// GetInfoRequest 获取 info 请求
type GetInfoRequest struct {
	Ids []int `json:"ids" binding:"required"`
} // @name GetInfoRequest

type userInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
}

// GetProfileRequest 获取 profile 请求
type GetProfileRequest struct {
	Id int `json:"id"`
} // @name GetProfileRequest

// type UserProfile struct {
// 	Id                        int    `json:"id"`
// 	Name                      string `json:"name"`
// 	Avatar                    string `json:"avatar"`
// 	Email                     string `json:"email"`
// 	Role                      string `json:"role"`
// }

type user struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Role   string `json:"role"`
} // @name user

// ListResponse 获取 userList 响应
type ListResponse struct {
	Count int    `json:"count"`
	List  []user `json:"list"`
} // @name ListResponse

// UpdateInfoRequest 更新 userInfo 请求
type UpdateInfoRequest struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Signature string `json:"signature"`
} // @name UpdateInfoRequest
