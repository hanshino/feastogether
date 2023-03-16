package main

import (
	"feastogether/client"
	"feastogether/config"
	"fmt"
	"log"
)

func main() {
	// 讀取 config
	cfg, err := config.GetConfig("./")
	if err != nil {
		log.Println(err)
		return
	}

	// 取得 token
	token := client.GetToken(cfg.UserConfig)
	if token == "" {
		return
	}

	// 立即定位 , 取得定位開始 - 過期時間
	expirationTime := client.GetSaveSeats(
		cfg.UserConfig.Account,
		token,
		cfg.RestaurantConfig)

	// 判斷是否取得訂位開始時間
	if expirationTime != "" {
		// 確認定位
		msg := client.SaveBooking(
			cfg.UserConfig.Account,
			client.GetToken(cfg.UserConfig),
			cfg.RestaurantConfig)

		fmt.Println(msg)
	}

}
