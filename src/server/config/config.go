package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

// New returns an initialized config instance with configurations read in.
func New(dir string) *Context {
	c := &Context{conf: &sync.Map{}}

	// Set configuration file path.
	viper.SetConfigName("config")
	viper.AddConfigPath(dir)

	// Read in configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}

	// Store configuration content.
	c.conf.Store("mssql", viper.GetString("mssql"))
	c.conf.Store("server.port", viper.GetString("server.port"))
	c.conf.Store("server.website", viper.GetString("server.website"))
	c.conf.Store("email.smtp.port", viper.GetString("email.smtp.port"))
	c.conf.Store("email.smtp.url", viper.GetString("email.smtp.url"))
	c.conf.Store("email.system.account", viper.GetString("email.system.account"))
	c.conf.Store("email.system.password", viper.GetString("email.system.password"))
	c.conf.Store("email.help_desk_ticket_receiver", viper.GetString("email.help_desk_ticket_receiver"))
	c.conf.Store("email.suggestion_receiver", viper.GetString("email.suggestion_receiver"))
	c.conf.Store("cache.enable", viper.GetString("cache.enable"))

	return c
}
