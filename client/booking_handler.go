package client

import (
	"encoding/json"
	"feastogether/config"
	"feastogether/fetch"
	"log"
)

// api
const (
	LOGIN_API      = "https://www.feastogether.com.tw/api/994f5388-d001-4ca4-a7b1-72750d4211cf/custSignIn"
	SAVE_SEATS_API = "https://www.feastogether.com.tw/api/booking/saveSeats"
	SAVE_SAETS_API = "https://www.feastogether.com.tw/api/booking/saveSaets"
	BOOKING_API    = "https://www.feastogether.com.tw/api/booking/booking"
)

// 取得 Token
func GetTokne(user config.UserConfig) string {

	payload := Login{
		Act: user.Account,
		Pwd: user.Password,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal struct to JSON:%v\n", err)
		return ""
	}

	resp, err := fetch.Post(LOGIN_API, payloadBytes, payload.Act, "")
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	var data Response
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Failed to decode response body: %v\n", err)
		return ""
	}
	if data.StatusCode != 1000 {
		log.Println(data.Result.Msg)
		return ""
	}
	return data.Result.CustomerLoginResp.Token
}

// 不清楚是什麼
func GetSaveSaets(act string, token string) string {

	resp, err := fetch.Post(SAVE_SAETS_API, nil, act, token)
	if err != nil {
		log.Println(err)
		return ""
	}

	defer resp.Body.Close()

	var data SaveSaetsResponse
	if json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Failed to decode response body: %v\n", err)
		return ""
	}

	if data.StatusCode != 1000 {
		log.Println(data)
		return ""
	}
	return data.Result
}

// 立即定位
func GetSaveSeats(act string, token string, payload config.RestaurantConfig) string {

	saveSeats := SaveSeats{
		StoreID:     payload.StoreID,
		PeopleCount: payload.PeopleCount,
		MealPeriod:  payload.MealPeriod,
		MealDate:    payload.MealDate,
		MealTime:    payload.MealTime,
		MealSeq:     4,
		Zkde:        nil,
	}

	if saets := GetSaveSaets(act, token); saets != "" {
		saveSeats.Zkde = saets
	}

	payloadBytes, err := json.Marshal(saveSeats)
	if err != nil {
		log.Printf("Failed to marshal struct to JSON:%v\n", err)
		return ""
	}

	resp, err := fetch.Post(SAVE_SEATS_API, payloadBytes, act, token)
	if err != nil {
		log.Println(err)
		return ""
	}

	defer resp.Body.Close()

	var data Response
	if json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Failed to decode response body: %v\n", err)
		return ""
	}

	if data.StatusCode != 1000 {
		log.Println(data.Result.Msg)
		return ""
	}

	return data.Result.ExpirationTime
}

// 送出定位
func SaveBooking(act string, token string, payload config.RestaurantConfig) string {

	booking := Booking{
		StoreID:    payload.StoreID,
		MealPeriod: payload.MealPeriod,
		MealDate:   payload.MealDate,
		MealTime:   payload.MealTime,
		MealSeq:    4,
		Special:    0,
		ChildSeat:  0,
		Adult:      2,
		Child:      0,
		ChargeList: []struct {
			Seq   int "json:\"seq\""
			Count int "json:\"count\""
		}{
			// 大人
			{
				Seq:   201,
				Count: payload.PeopleCount,
			},
			// 小孩
			{
				Seq:   202,
				Count: 0,
			},
		},
		StoreCode:    "NTBQ",
		RedirectType: "iEat_card",
		Domain:       "https://www.feastogether.com.tw",
		PathFir:      "booking",
		PathSec:      "result",
	}

	payloadBytes, err := json.Marshal(booking)
	if err != nil {
		log.Printf("Failed to marshal struct to JSON:%v\n", err)
		return ""
	}

	resp, err := fetch.Post(BOOKING_API, payloadBytes, act, token)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	var data Response
	if json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Failed to decode response body: %v\n", err)
		return ""
	}

	if data.StatusCode != 1000 {
		log.Println(data.Result.Msg)
		return ""
	}
	return data.Message
}
