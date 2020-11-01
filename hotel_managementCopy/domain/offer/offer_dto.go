package offer

const (
	//StatusActive status of the user is active
	StatusActive = "active"
)

type Offer struct {
	Offers []OfferDetails `json:"offers,omitempty"`
}

type OfferDetails struct {
	CmOfferID    string       `json:"cm_offer_id,omitempty"`
	Hotel        Hotel        `json:"hotel,omitempty"`
	Room         Room         `json:"room,omitempty"`
	RatePlan     RatePlan     `json:"rate_plan,omitempty"`
	OriginalData OriginalData `json:"original_data,omitempty"`
	Capacity     Capacity     `json:"capacity,omitempty"`
	Number       uint         `json:"number,omitempty"`
	Price        float32      `json:"price,omitempty"`
	Currency     string       `json:"currency,omitempty"`
	CheckIn      string       `json:"check_in,omitempty"`
	CheckOut     string       `json:"check_out,omitempty"`
	Fees         []Fee        `json:"fees,omitempty"`
}

type Fee struct {
	Type        string  `json:"type,omitempty"`
	Description string  `json:"description,omitempty"`
	Included    bool    `json:"included,omitempty"`
	Percent     float32 `json:"percent,omitempty"`
}

type Capacity struct {
	MaxAdults     uint `json:"max_adults,omitempty"`
	ExtraChildren uint `json:"extra_children,omitempty"`
}

type OriginalData struct {
	GuaranteePolicy GuaranteePolicy `json:"GuaranteePolicy,omitempty"`
}
type GuaranteePolicy struct {
	Required bool `json:"Required,omitempty"`
}

type RatePlan struct {
	HotelId            string               `json:"hotel_id,omitempty"`
	RatePlanId         string               `json:"rate_plan_id,omitempty"`
	CancellationPolicy []CancellationPolicy `json:"cancellation_policy,omitempty"` //array of object
	Name               string               `json:"name,omitempty"`
	OtherConditions    []string             `json:"other_conditions,omitempty"`
	MealPlan           string               `json:"meal_plan,omitempty"`
}
type CancellationPolicy struct {
	Type             string `json:"type,omitempty"`
	ExpiresDayBefore uint   `json:"expires_day_before,omitempty"`
}

type Room struct {
	HotelId     string   `json:"hotel_id,omitempty"`
	RroomId     string   `json:"rroom_id,omitempty"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name,omitempty"`
	Capacity    Capacity `json:"capacity,omitempty"`
}
type Hotel struct {
	HotelId     string   `json:"hotel_id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Country     string   `json:"country,omitempty"`
	Address     string   `json:"address,omitempty"`
	Latitude    float64  `json:"latitude,omitempty"`
	Longitude   float64  `json:"longitude,omitempty"`
	Telephone   string   `json:"telephone,omitempty"`
	Amenities   []string `json:"amenities,omitempty"`
	Description string   `json:"description,omitempty"`
	RoomCount   uint     `json:"room_count,omitempty"`
	Currency    string   `json:"currency,omitempty"`
}
