package response

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
