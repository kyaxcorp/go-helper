package config

import "github.com/kyaxcorp/go-helper/_struct"

var cfg Config

type Config struct {
	// Stage -> prod, dev, beta, testing etc...
	Stage string `yaml:"stage" mapstructure:"stage" default:"dev"`
	// App Data Path (all other subfolders will be stored there...)
	AppDataPath string `yaml:"app_data_path" mapstructure:"app_data_path" default:".appdata"`
	// Where the FileLocks are saved
	LocksPath string `yaml:"locks_path" mapstructure:"locks_path" default:"locks"`
	// Certs Path -> where we store certificates for different use cases
	CertsPath string `yaml:"certs_path" mapstructure:"certs_path" default:"certs"`
	// Temporary folder...
	TempPath string `yaml:"temp_path" mapstructure:"temp_path" default:"tmp"`
	// How many seconds should wait until the app gracefully shuts down all the services before os.exit(0)
	OnShutdownWaitSeconds int `yaml:"on_shutdown_wait_seconds" mapstructure:"on_shutdown_wait_seconds" default:"2"`
	// Where the PIDs are saved!
	PIDsPath string `yaml:"pids_path" mapstructure:"pids_path" default:"pids"`
}

func GetConfig() Config {
	return cfg
}

func SetDefaults(cf *Config) error {
	return _struct.SetDefaultValues(cf)
}

func init() {
	err := SetDefaults(&cfg)
	if err != nil {
		panic(err)
	}
}
