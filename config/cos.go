package config

type Cos struct {
	Video_bucket_url  string `mapstructure:"video_bucket_url" json:"video_bucket_url" yaml:"video_bucket_url"`
	Cover_bucket_url  string `mapstructure:"cover_bucket_url" json:"cover_bucket_url" yaml:"cover_bucket_url"`
	Avatar_bucket_url string `mapstructure:"avatar_bucket_url" json:"avatar_bucket_url" yaml:"avatar_bucket_url"`
	Secretid          string `mapstructure:"secretid" json:"secretid" yaml:"secretid"`
	Secretkey         string `mapstructure:"secretkey" json:"secretkey" yaml:"secretkey"`
}
