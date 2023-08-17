package initialization

import (
	"douyin/global"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func InitializeCos() {
	videoUrl, _ := url.Parse(global.SERVER_CONFIG.Cos.Video_bucket_url)
	videoB := &cos.BaseURL{BucketURL: videoUrl}
	global.SERVER_COS_VIDEO = cos.NewClient(videoB, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.SERVER_CONFIG.Cos.Secretid,
			SecretKey: global.SERVER_CONFIG.Cos.Secretkey,
		},
	})

	coverUrl, _ := url.Parse(global.SERVER_CONFIG.Cos.Cover_bucket_url)
	coverB := &cos.BaseURL{BucketURL: coverUrl}
	global.SERVER_COS_COVER = cos.NewClient(coverB, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.SERVER_CONFIG.Cos.Secretid,
			SecretKey: global.SERVER_CONFIG.Cos.Secretkey,
		},
	})

	avatarUrl, _ := url.Parse(global.SERVER_CONFIG.Cos.Avatar_bucket_url)
	avatarB := &cos.BaseURL{BucketURL: avatarUrl}
	global.SERVER_COS_AVATAR = cos.NewClient(avatarB, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.SERVER_CONFIG.Cos.Secretid,
			SecretKey: global.SERVER_CONFIG.Cos.Secretkey,
		},
	})

}
