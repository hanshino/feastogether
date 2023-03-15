package config

import (
	"log"

	"github.com/spf13/viper"
)

type UserConfig struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RestaurantConfig struct {
	StoreID     string `json:"storeId"`
	PeopleCount int    `json:"peopleCount"`
	MealPeriod  string `json:"mealPeriod"`
	MealDate    string `json:"mealDate"`
	MealTime    string `json:"mealTime"`
}

type Config struct {
	UserConfig
	RestaurantConfig
}

func GetConfig(path string) (*Config, error) {
	// Set the name and path of the config file to read
	viper.SetConfigName("config")
	viper.AddConfigPath(path)  // 目錄
	viper.SetConfigType("ini") // 格式

	// Set the config file type to ini format
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to load config file: %v\n", err)
		return nil, err
	}

	// Load the values into the UserConfig struct
	var userConfig UserConfig
	if err := viper.UnmarshalKey("user", &userConfig); err != nil {
		log.Printf("Failed to unmarshal user config: %v\n", err)
		return nil, err
	}

	// Load the values into the RestaurantConfig struct
	var restaurantConfig RestaurantConfig
	if err := viper.UnmarshalKey("restaurant", &restaurantConfig); err != nil {
		log.Printf("Failed to unmarshal restaurant config: %v\n", err)
		return nil, err
	}

	config := Config{
		userConfig,
		restaurantConfig,
	}

	return &config, nil
}
