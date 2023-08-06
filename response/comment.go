package response

type Comment_Response struct {
	Id         int64         `json:"id,omitempty"`
	Commentor  User_Response `json:"user,omitempty"`
	Content    string        `json:"content,omitempty"`
	CreateDate string        `json:"create_date,omitempty"`
}
