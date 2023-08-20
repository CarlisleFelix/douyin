package response

type Publish_Action_Response struct {
	Response
}

type Publish_List_Response struct {
	Response
	VideoList []Video_Response `json:"video_list,omitempty"`
}
