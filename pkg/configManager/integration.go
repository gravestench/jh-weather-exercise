package configManager

import (
	"fmt"
	"log"
	"path/filepath"
)

// Configurable represents something that is managed by our
// config file manager
type Configurable interface {
	ConfigFileName() string
	DefaultConfigData() []byte
	IngestConfigData(handle *ConfigHandle) error
}

// InitConfiguration will take a Configurable instance and init the config
// file with default data or load the existing data and pass it back to the
// Configurable to do whatever it needs to do.
func (s *Service) InitConfiguration(c Configurable) error {
	// always make sure our root directory is good
	if err := s.ensureExistingRootConfigDirectory(); err != nil {
		return fmt.Errorf("ensuring existing root config directory: %v", err)
	}

	// create a new config handle for the file within our config dir
	configPath := filepath.Join(s.RootDirectory, c.ConfigFileName())
	handle, err := s.newConfigHandle(configPath)
	if err != nil {
		return fmt.Errorf("getting config handle: %v", err)
	}

	// if there is no existing data in the file, use the defaults
	if existingData, errLoad := handle.Data(); errLoad == nil && len(existingData) == 0 {
		log.Printf("setting default values for config file %q", configPath)

		if errSetDefault := handle.SetData(c.DefaultConfigData()); errSetDefault != nil {
			return fmt.Errorf("applying default configuration data: %v", err)
		}
	}

	log.Printf("ingesting config file %q", configPath)

	// pass the config handle to the configurable so it can do whatever it needs to do
	if errIngest := c.IngestConfigData(handle); errIngest != nil {
		return fmt.Errorf("ingesting config data for %q: %v", c.ConfigFileName(), errIngest)
	}

	return nil
}
