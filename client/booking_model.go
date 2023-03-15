package client

type Login struct {
	Act string `json:"act"`
	Pwd string `json:"pwd"`
}

type CustomerLoginResp struct {
	Address     string `json:"address"`
	Area        string `json:"area"`
	BirthDay    string `json:"birthDay"`
	BirthMonth  string `json:"birthMonth"`
	BirthYear   string `json:"birthYear"`
	City        string `json:"city"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	MemberShip  string `json:"memberShip"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Token       string `json:"token"`
}

type Result struct {
	// success
	ACT               string            `json:"_ACT"`
	P                 string            `json:"_P"`
	CustomerLoginResp CustomerLoginResp `json:"customerLoginResp"`

	// Save Seats
	ExpirationTime string `json:"expirationTime"`

	// booking
	BookingID    string `json:"bookingId"`
	BookingState string `json:"bookingState"`
	BrandName    string `json:"brandName"`
	ExpireTime   string `json:"expireTime"`
	MealDate     string `json:"mealDate"`
	MealPeriod   string `json:"mealPeriod"`
	MealTime     string `json:"mealTime"`
	PaymentState string `json:"paymentState"`
	StoreName    string `json:"storeName"`

	// error
	Msg string `json:"msg"`
}

type Response struct {
	// success
	Message    string `json:"message"`
	Result     Result `json:"result"`
	StatusCode int    `json:"statusCode"`
	Timestamp  string `json:"timestamp"`

	// error
	Path string `json:"path"`
}

type SaveSaetsResponse struct {
	// success
	Message    string `json:"message"`
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Timestamp  string `json:"timestamp"`
}

type SaveSeats struct {
	StoreID     string `json:"storeId"`
	PeopleCount int    `json:"peopleCount"`
	MealPeriod  string `json:"mealPeriod"`
	MealSeq     int    `json:"mealSeq"`
	MealDate    string `json:"mealDate"`
	MealTime    string `json:"mealTime"`
	Zkde        any    `json:"zkde"`
}

type Booking struct {
	StoreID     string `json:"storeId"`
	MealDate    string `json:"mealDate"`
	MealPurpose string `json:"mealPurpose"`
	MealSeq     int    `json:"mealSeq"`
	MealTime    string `json:"mealTime"`
	MealPeriod  string `json:"mealPeriod"`
	Special     int    `json:"special"`
	ChildSeat   int    `json:"childSeat"`
	Adult       int    `json:"adult"`
	Child       int    `json:"child"`
	ChargeList  []struct {
		Seq   int `json:"seq"`
		Count int `json:"count"`
	} `json:"chargeList"`
	StoreCode    string `json:"storeCode"`
	RedirectType string `json:"redirectType"`
	Domain       string `json:"domain"`
	PathFir      string `json:"pathFir"`
	PathSec      string `json:"pathSec"`
}
