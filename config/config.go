package config

type Config struct {
	KafkaConfig `ini:"kafka"`
	Collect     `ini:"collect"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `int:"topic"`
}

type Collect struct {
	LogFilePath string `ini:"logfile_path"`
}
