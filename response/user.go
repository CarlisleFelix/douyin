package response

type User_Register_Response struct {
	Response
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type User_Login_Response struct {
	Response
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type User_Interface_Response struct {
	Response
	User_Response `json:"user"`
}
