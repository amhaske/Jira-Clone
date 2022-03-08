package skeletons

type UsersBaseReq struct {
	UserId uint `json:"user_id"`
}

type UserIdReq struct {
	UserId int `json:"user_id"`
}

type UserEntry struct {
	UserId   uint `json:"user_id"`
	UserName uint `json:"user_name"`
	UserRole uint `json:"user_role"`
}
