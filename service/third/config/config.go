package config

type ServiceConfig struct {
	Name string   `json:"name"`
	Host string   `json:"host"`
	Tags []string `json:"tags"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RedisConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Database   int    `json:"database"`
	Expiration int    `json:"expiration"`
}

type AliSmsConfig struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SignName        string `json:"signName"`
	TemplateCode    string `json:"templateCode"`
}

type QiNiuYunConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	Domain    string `json:"domain"`
}

type ServerConfig struct {
	Service  ServiceConfig  `json:"service"`
	Consul   ConsulConfig   `json:"consul"`
	Redis    RedisConfig    `json:"redis"`
	AliSms   AliSmsConfig   `json:"sms"`
	QiNiuYun QiNiuYunConfig `json:"qiniuyun"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}
