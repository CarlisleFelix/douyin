package response

type Comment_Action_Response struct {
	Response
	Comment_Response `json:"comment,omitempty"`
}

type Comment_List_Response struct {
	Response
	CommentList []Comment_Response `json:"comment_list,omitempty"`
}
