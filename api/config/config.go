package config

type ApiConfig struct {
	Name string   `json:"name"`
	Host string   `json:"host"`
	Tags []string `json:"tags"`
}
type ServiceConfig struct {
	User       string `json:"user"`
	Admin      string `json:"admin"`
	Meeting    string `json:"meeting"`
	ThirdParty string `json:"thirdParty"`
}

type JWTConfig struct {
	SigningKey string `json:"signingKey"`
	Expiration int    `json:"expiration"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type JaegerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	Api     ApiConfig     `json:"api"`
	Service ServiceConfig `json:"service"`
	JWT     JWTConfig     `json:"jwt"`
	Consul  ConsulConfig  `json:"consul"`
	Jaeger  JaegerConfig  `json:"jaeger"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"userService"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataId"`
	Group     string `mapstructure:"group"`
}
