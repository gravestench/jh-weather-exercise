package configManager

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Service is responsible for integrating with other parts of the application that
// want a managed configuration file in a common config directory.
type Service struct {
	RootDirectory string
	Handles       map[string]*ConfigHandle
}

func (s *Service) ensureExistingRootConfigDirectory() (err error) {
	s.RootDirectory, err = expandHomeDirectoryPath(s.RootDirectory)
	if err != nil {
		return fmt.Errorf("expanding home directory: %v", err)
	}

	_ = os.MkdirAll(s.RootDirectory, 0o755)

	// if root directory exists and is not empty string, we are done
	handle, err := os.Stat(s.RootDirectory)
	if s.RootDirectory != "" && err == nil && handle.IsDir() {
		return nil
	}

	// otherwise, let's set the root directory to the user config dir
	if s.RootDirectory, err = os.UserConfigDir(); err != nil {
		return fmt.Errorf("getting default config directory: %v", err)
	}

	return nil
}

func (s *Service) newConfigHandle(configPath string) (*ConfigHandle, error) {
	if info, err := os.Stat(configPath); info != nil && info.IsDir() {
		return nil, fmt.Errorf("file path is a directory")
	} else if os.IsNotExist(err) {
		log.Printf("creating config file %q", configPath)

		if _, errCreate := os.Create(configPath); errCreate != nil {
			return nil, fmt.Errorf("creating config file: %v", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("something went wrong: %v", err)
	}

	handle := &ConfigHandle{
		manager:  s,
		filepath: configPath,
	}

	if s.Handles == nil {
		s.Handles = make(map[string]*ConfigHandle)
	}

	s.Handles[configPath] = handle

	return handle, nil
}

func (s *Service) ConfigDirectory() (string, error) {
	if err := s.ensureExistingRootConfigDirectory(); err != nil {
		return "", err
	}

	return s.RootDirectory, nil
}

func expandHomeDirectoryPath(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path = strings.Replace(path, "~", homeDir, 1)

	return path, nil
}
