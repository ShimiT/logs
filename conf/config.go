package conf

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.wdf.sap.corp/devx-wing/logs/logger"
)

const (
	varLogLevel = "log.level"
	// Path to the config env key
	varPathToConfig = "log.cfg"
)

type Configuration struct {
	v *viper.Viper
	l *logger.Logger
}

func New(level string, version string) *Configuration {
	c := Configuration{
		v: viper.New(),
		l: logger.NewLogger(level, version),
	}
	c.v.SetDefault(varPathToConfig, "./logcfg.yaml")
	c.v.SetDefault(varLogLevel, "debug")
	c.v.AutomaticEnv()
	c.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.v.SetTypeByDefaultValue(true)
	c.v.SetConfigFile(c.GetPathToConfig())
	err := c.v.ReadInConfig() // Find and read the config file
	if _, ok := err.(*os.PathError); ok {
		c.l.Warnf("config file '%s' not found", c.GetPathToConfig())
	} else if err != nil { // Handle other errors that occurred while reading the config file
		c.l.Panic(fmt.Errorf("not able to load config file: %s", err))
	}
	c.l.SetLevel(c.GetLogLevel())
	// monitor the changes in the config file
	c.v.WatchConfig()
	c.v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config was changed to: ", c.GetLogLevel())
		c.l.SetLevel(c.GetLogLevel())
	})
	return &c
}

// GetLogLevel returns the log level
func (c *Configuration) GetLogLevel() string {
	s := c.v.GetString(varLogLevel)
	return s
}

// GetPathToConfig returns the path to the config file
func (c *Configuration) GetPathToConfig() string {
	return c.v.GetString(varPathToConfig)
}
