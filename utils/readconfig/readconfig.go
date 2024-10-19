package readconfig

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config structure to hold the parsed configuration
type Config struct {
	Settings struct {
		RemoteStorage              bool   `yaml:"remoteStorage"`
		StoreEncryptionKeyToServer bool   `yaml:"storeEncryptionKeyToServer"`
		StoreDecryptionKeyToServer bool   `yaml:"storeDecryptionKeyToServer"`
		Email                      string `yaml:"email"`
		Password                   string `yaml:"password"`
	} `yaml:"settings"`
	Backup struct {
		Directories          []string `yaml:"directories"`
		TempDirectory        string   `yaml:"temp_directory"`
		DestinationDirectory string   `yaml:"destination_directory"`
		BackupInterval       string   `yaml:"backup_interval"`
	} `yaml:"backup"`
	Encryption struct {
		Enabled     bool   `yaml:"enabled"`
		Method      string `yaml:"method"`
		KeyLocation string `yaml:"keyLocation"`
	} `yaml:"encryption"`
	Logging struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"logging"`
}

// GetConf reads the configuration file and returns the config struct
func GetConf() Config {
	data, err := os.ReadFile("config/cipher-suitcase.yml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("error unmarshaling YAML: %v", err)
	}

	return config
}
