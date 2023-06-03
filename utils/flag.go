package utils

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var ViperConfig Config

func InitFlags() {
	log.Println("init flags")
	configPath := flag.String("config", "", "pass config file path")
	flag.Parse()

	initConfig(*configPath)
	log.Printf("got config : %+v\n", ViperConfig)
}

func initConfig(configPath string) {
	if configPath == "" {
		log.Fatalf("Provide config file to start the service")
	}
	file := strings.Split(filepath.Base(configPath), ".")
	dir := filepath.Dir(configPath)
	dir = dir + "/"

	log.Printf("Full Path : %v \n File Details : %v \n Dir Details : %v \n", configPath, file, dir)

	viper.SetConfigType(file[1]) // file extension yaml/json
	viper.SetConfigName(file[0]) // name of config file (without extension)
	viper.AddConfigPath(dir)     // path to look for the config file in
	err := viper.ReadInConfig()  // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	err = viper.Unmarshal(&ViperConfig)
	if err != nil {
		log.Fatalf("Viper config unable to decode into struct, %v", err.Error())
	}

}
