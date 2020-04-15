package cli

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	defaultConfigName = "config.yaml"
)

func configPath() string {
	return filepath.Join(configHome(), defaultConfigName)
}

func configHome() string {
	if xdgPath := os.Getenv("XDG_CONFIG_HOME"); xdgPath != "" {
		return filepath.Join(xdgPath, "tbledit")
	}
	return filepath.Join(homeDir(), ".config", "tbledit")
}

func homeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Home directory not found. %+v\n", err)
	}
	return home
}

func writeConfig() error {
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return nil
	/*
		f, err := cfgFileWriter()
		if err != nil {
			return err
		}

		defer f.Close()

		b, err := yaml.Marshal(viper.AllSettings())
		if err != nil {
			return errors.New("unable to encode configuration to YAML format")
		}

		_, err = f.Write(b)
		if err != nil {
			return errors.New("unable to write configuration")
		}

		return nil*/
}

func cfgFileWriter() (io.WriteCloser, error) {
	f, err := os.Create(cfgFile)
	if err != nil {
		return nil, err
	}
	if err := os.Chmod(cfgFile, 0600); err != nil {
		return nil, err
	}

	return f, nil
}
