package vo

type Page struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}

type Stats struct {
	DeviceAll      int64 `json:"device_all"`
	DeviceUsing    int64 `json:"device_using"`
	DeviceInactive int64 `json:"device_inactive"`
	BookingMy      int64 `json:"booking_my"`
}

type StatusCnt struct {
	Status int32
	Cnt    int64
}
