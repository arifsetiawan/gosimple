package main

import "github.com/spf13/viper"

func loadDefaultSettings() {
	viper.SetDefault("CouchbaseURI", "couchbase://192.168.99.100")
	viper.SetDefault("CouchbaseBucket", "gobase")
	viper.SetDefault("CouchbasePassword", "Test1234")
	viper.SetDefault("LogPath", "logs/echo")
	viper.SetDefault("Port", ":4000")
}

func initializeConfig() error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("gobase")
	loadDefaultSettings()
	return nil
}
