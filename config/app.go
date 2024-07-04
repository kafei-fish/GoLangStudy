package config

// ServerConfig 配置信息
type ServerConfig struct {
	Mysql mysql `yaml:"mysql"` // 嵌入MySQL配置
}
