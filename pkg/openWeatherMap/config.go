package openWeatherMap

import (
	"encoding/json"
	"fmt"

	"github.com/gravestench/jh-weather-exercise/pkg/configManager"
)

var _ configManager.Configurable = &Service{}

type Config struct {
	ApiKey string
}

func (s *Service) ConfigFileName() string {
	return "open_weather_map.json"
}

func (s *Service) DefaultConfigData() []byte {
	// should probably bubble up the error, but if
	// this explodes then we are definitely doing
	// something wrong...
	data, _ := json.MarshalIndent(&Config{}, "", "\t")

	return data
}

func (s *Service) IngestConfigData(handle *configManager.ConfigHandle) error {
	data, err := handle.Data()
	if err != nil {
		return fmt.Errorf("loading config data: %v", err)
	}

	if err = json.Unmarshal(data, &s.Config); err != nil {
		return fmt.Errorf("unmarshalling config data: %v", err)
	}

	if s.Config.ApiKey == "" {
		return fmt.Errorf("you need to set up your API key in config file %q", handle.FilePath())
	}

	return nil
}
