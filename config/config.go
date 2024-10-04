package config

import (
    "log"
    "github.com/spf13/viper"
)


type Config struct {
    Database struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        Name     string `mapstructure:"name"`
    } `mapstructure:"database"`
    Server struct {
        Port int `mapstructure:"port"`
    } `mapstructure:"server"`
}

var AppConfig Config

/**
Load ConFig FilE
*/
func LoadConfig() {
    viper.SetConfigName("config") 
    viper.SetConfigType("toml")   
    viper.AddConfigPath(".")      

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Unable to decode into struct: %v", err)
    }
}
