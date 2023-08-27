package response

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type User_Response struct {
	Id              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type Comment_Response struct {
	Id         int64         `json:"id,omitempty"`
	Commenter  User_Response `json:"user,omitempty"` // 评论者是自己，应该如何填写IsFollow字段
	Content    string        `json:"content,omitempty"`
	CreateDate string        `json:"create_date,omitempty"`
}

type Friend_Response struct {
	User_Response
	Message string `json:"message,omitempty"`
	Msgtype int64  `json:"msgType,omitempty"`
}

type Message_Response struct {
	Id         int64  `json:"id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	FromUserID int64  `json:"from_user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type Video_Response struct {
	Id            int64         `json:"id,omitempty"`
	Author        User_Response `json:"author,omitempty"`
	PlayUrl       string        `json:"play_url,omitempty"`
	CoverUrl      string        `json:"cover_url,omitempty"`
	FavoriteCount int64         `json:"favorite_count,omitempty"`
	CommentCount  int64         `json:"comment_count,omitempty"`
	IsFavorite    bool          `json:"is_favorite,omitempty"`
	Title         string        `json:"title,omitempty"`
}
