package ginWebServer

import (
	"encoding/json"
	"fmt"

	"github.com/gravestench/jh-weather-exercise/pkg/configManager"
)

const (
	defaultPort = 8080
)

type Config struct {
	Port int
}

// implementing this interface allows the server
// to have its config file managed
var _ configManager.Configurable = &Service{}

func (s *Service) ConfigFileName() string {
	return "web_server.json"
}

func (s *Service) DefaultConfigData() []byte {
	cfg := &Config{
		Port: defaultPort,
	}

	// should probably bubble up the error, but if
	// this explodes then we are definitely doing
	// something wrong...
	data, _ := json.MarshalIndent(cfg, "", "\t")

	return data
}

func (s *Service) IngestConfigData(handle *configManager.ConfigHandle) error {
	data, err := handle.Data()
	if err != nil {
		return fmt.Errorf("loading config data: %v", err)
	}

	_ = json.Unmarshal(data, &s.Config)

	if s.Port == 0 {
		return fmt.Errorf("port cannot be 0")
	}

	return nil
}
