package response

type User_Register_Response struct {
	Response
	UserId int64  `json:"user_id"` // ！notice：改了  string->int
	Token  string `json:"token"`
}

type User_Login_Request struct {
	Response
	UserName int64  `json:"username"` // ！notice：改了  string->int
	Password string `json:"password"`
}

type User_Register_Request struct {
	Response
	UserName int64  `json:"username"` // ！notice：改了  string->int
	Password string `json:"password"`
}

type User_Login_Response struct {
	Response
	UserId int64  `json:"user_id"` // ！notice：改了  string->int
	Token  string `json:"token"`
}

type User_Interface_Response struct {
	Response
	User_Response
}

type User_Empty struct { //notice：更新
	Response
}
