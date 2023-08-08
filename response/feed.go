package response

type Feed_Response struct {
	Response
	VideoList []Video_Response `json:"video_list,omitempty"`
	NextTime  int64            `json:"next_time,omitempty"`
}
type Feed_Novideo_Response struct {
	Response
	NextTime int64 `json:"next_time"`
}
