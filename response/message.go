package response

type Message_Chat_Response struct {
	Comment_Response
	MessageList []Message_Response `json:"message_list,omitempty"`
}

type Message_Action_Response struct {
	Comment_Response
}
