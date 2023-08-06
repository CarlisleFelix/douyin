package response

type Friend_Response struct {
	User_Response
	Message string `json:"message,omitempty"`
	Msgtype int64  `json:"msgType,omitempty"`
}
