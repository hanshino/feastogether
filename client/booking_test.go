package client_test

import (
	"feastogether/client"
	"feastogether/config"
	"fmt"
	"log"
	"testing"
)

func TestGetToken(t *testing.T) {
	if cfg, err := config.GetConfig(".."); err != nil {
		log.Println(err)
	} else {
		fmt.Println(client.GetToken(cfg.UserConfig))
	}
}

func TestGetSaveSaets(t *testing.T) {
	if cfg, err := config.GetConfig(".."); err != nil {
		log.Println(err)
	} else {
		fmt.Println(client.GetSaveSaets(
			cfg.UserConfig.Account,
			client.GetToken(cfg.UserConfig)))
	}
}
func TestGetSaveSeats(t *testing.T) {
	if cfg, err := config.GetConfig(".."); err != nil {
		log.Println(err)
	} else {
		fmt.Println(client.GetSaveSeats(
			cfg.UserConfig.Account,
			client.GetToken(cfg.UserConfig),
			cfg.RestaurantConfig))
	}
}

func TestSaveBooking(t *testing.T) {

	if cfg, err := config.GetConfig(".."); err != nil {
		log.Println(err)
	} else {
		fmt.Println(client.GetSaveSeats(
			cfg.UserConfig.Account,
			client.GetToken(cfg.UserConfig),
			cfg.RestaurantConfig))

		fmt.Println(client.SaveBooking(
			cfg.UserConfig.Account,
			client.GetToken(cfg.UserConfig),
			cfg.RestaurantConfig))
	}
}
