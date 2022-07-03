package leafbot

// Config
// @Description: 配置信息
//
type Config struct {
	NickName     []string `json:"nick_name" yaml:"nick_name" hjson:"nick_name"`
	Admin        int64    `json:"admin" yaml:"admin" hjson:"admin"`
	SuperUser    []int64  `json:"super_user" yaml:"super_user" hjson:"super_user"`
	CommandStart []string `json:"command_start" yaml:"command_start" hjson:"command_start"`
	LogLevel     string   `json:"log_level" yaml:"log_level"`
}

var DefaultConfig = Config{
	NickName:     []string{"leafBot"},
	Admin:        0,
	SuperUser:    nil,
	CommandStart: []string{"/"},
	LogLevel:     "",
}

var (
	runConfig *Config
)

func GetConfig() *Config {
	return runConfig
}

func LoadConfig(config *Config) {
	runConfig = config
}
