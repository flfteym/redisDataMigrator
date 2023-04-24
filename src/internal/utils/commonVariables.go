package utils

// Config 从redisDataMigrator.yml文件读取的配置
type Config struct {
	SrcRedis struct {
		Version string   `yaml:"version"`
		Passwd  string   `yaml:"passwd"`
		Url     []string `yaml:"url"`
	} `yaml:"srcRedis"`

	DestRedis struct {
		Version string   `yaml:"version"`
		Passwd  string   `yaml:"passwd"`
		Url     []string `yaml:"url"`
	} `yaml:"destRedis"`
}
