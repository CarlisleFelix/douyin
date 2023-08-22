package config

type Captcha struct {
	AppID      int
	AppKey     string
	TemplateID int
	SMSSign    string
	Params     []string
}
