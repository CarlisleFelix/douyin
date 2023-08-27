package response

type Favorite_Action_Response struct {
	Response
}

type Favorite_List_Response struct {
	Response
	VideoList []Video_Response `json:"video_list,omitempty"`
}
