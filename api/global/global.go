package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/importcjj/sensitive"
	"github.com/palp1tate/brevinect/api/config"
	"github.com/palp1tate/brevinect/proto/admin"
	"github.com/palp1tate/brevinect/proto/meeting"
	"github.com/palp1tate/brevinect/proto/third"
	"github.com/palp1tate/brevinect/proto/user"
)

var (
	Debug bool

	Filter *sensitive.Filter

	Translator ut.Translator

	ServerConfig *config.ServerConfig

	NacosConfig *config.NacosConfig

	UserServiceClient userProto.UserServiceClient

	AdminServiceClient adminProto.AdminServiceClient

	MeetingServiceClient meetingProto.MeetingServiceClient

	ThirdPartyServiceClient thirdProto.ThirdPartyServiceClient
)
