package response

type Relation_Action_Response struct {
	Response
}

type Relation_Follow_List_Response struct {
	Response
	UserList []User_Response `json:"user_list,omitempty"`
}

type Relation_Follower_List_Response struct {
	Response
	UserList []User_Response `json:"user_list,omitempty"`
}

type Relation_Friend_List_Response struct {
	Response
	UserList []Friend_Response `json:"user_list,omitempty"`
}
