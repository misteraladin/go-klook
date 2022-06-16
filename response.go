package klook

type LoginResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        uint   `json:"expires_in"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type SKUOrder struct {
	SKUID    uint   `json:"sku_id"`
	Quantity uint   `json:"quantity"`
	SKUPrice string `json:"sku_price"`
	Currency string `json:"currency"`
}

type SKUProduct struct {
	SKUID     uint   `json:"sku_id"`
	Title     string `json:"title"`
	SKUMinPax uint   `json:"sku_min_pax"`
	SKUMaxPax uint   `json:"sku_max_pax"`
	Required  bool   `json:"required"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Availablilty struct {
	Success bool          `json:"success"`
	Error   ErrorResponse `json:"error"`
}

type Cancellation struct {
	Success        bool          `json:"success"`
	CancelationID  uint          `json:"cancelation_id"`
	Timestamp      string        `json:"timestamp"`
	RefundAmount   string        `json:"refund_amount"`
	TransactionFee string        `json:"transaction_fee"`
	Currency       string        `json:"currency"`
	Error          ErrorResponse `json:"error"`
}

type CancelationRequest struct {
	RefundID       uint   `json:"refund_id"`
	Timestamp      uint   `json:"timestamp"`
	RefundAmount   string `json:"refund_amount"`
	TransactionFee string `json:"transaction_fee"`
	Currency       string `json:"currency"`
	Status         string `json:"status"`
}

type OrderBooking struct {
	ActivityName     string     `json:"activity_name"`
	PackageName      string     `json:"package_name"`
	ConfirmStatus    string     `json:"confirm_status"`
	VoucherUrl       string     `json:"voucher_url"`
	FileType         uint       `json:"file_type"`
	SKUS             []SKUOrder `json:"skus"`
	BookingRefNumber string     `json:"booking_ref_number"`
}

type OrderDetail struct {
	Success            bool               `json:"success"`
	AgentOrderID       string             `json:"agent_order_id"`
	KLKTechOrderID     string             `json:"klktech_order_id"`
	TotalAmount        string             `json:"total_amount"`
	Currency           string             `json:"currency"`
	CancelationRequest CancelationRequest `json:"cancelation_request"`
	Bookings           []OrderBooking     `json:"bookings"`
	Error              ErrorResponse      `json:"error"`
}

type Order struct {
	Success        bool          `json:"success"`
	AgentOrderID   string        `json:"agent_order_id"`
	KLKTechOrderID string        `json:"klktech_order_id"`
	TotalAmount    string        `json:"total_amount"`
	Currency       string        `json:"currency"`
	ConfirmStatus  string        `json:"confirm_status"`
	SKUS           []SKUOrder    `json:"skus"`
	VoucherURL     string        `json:"voucher_url"`
	Error          ErrorResponse `json:"error"`
}

type City struct {
	CityID uint   `json:"city_id"`
	Name   string `json:"name"`
}

type Country struct {
	CountryID uint   `json:"country_id"`
	Name      string `json:"name"`
	Cities    []City `json:"cities"`
}

type Countries struct {
	Success   bool      `json:"success"`
	Countries []Country `json:"countries"`
}

type ActivityThumbnail struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Width       uint   `json:"width"`
	Height      uint   `json:"height"`
}

type ActivitySimple struct {
	ActivityID  uint              `json:"activity_id"`
	Title       string            `json:"title"`
	Subtitle    string            `json:"subtitle"`
	Description string            `json:"description"`
	CityID      uint              `json:"city_id"`
	City        string            `json:"city"`
	CountryID   uint              `json:"country_id"`
	Country     string            `json:"country"`
	Thumbnail   ActivityThumbnail `json:"thumbnail"`
	Price       float32           `json:"price"`
}

type Activities struct {
	Success    bool             `json:"success"`
	Total      uint             `json:"total"`
	Page       uint             `json:"page"`
	PageSize   uint             `json:"page_size"`
	Activities []ActivitySimple `json:"activities"`
}

// Extra Info General
type ExtraInfoGeneral struct {
	TypeID      uint          `json:"type_id"`
	Type        string        `json:"type"`
	Content     string        `json:"content"`
	Hint        string        `json:"hint"`
	Required    bool          `json:"required"`
	TypeSpecial uint          `json:"type_special"`
	Options     []interface{} `json:"options"`
}

// Extra Info Traveler
type ExtraInfoTraveler struct {
	TypeID   uint   `json:"type_id"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Hint     string `json:"hint"`
	Required bool   `json:"required"`
}

// Extra Info
type ExtraInfo struct {
	General   []ExtraInfoGeneral  `json:"general"`
	Travelers []ExtraInfoTraveler `json:"travelers"`
}

type Balance struct {
	Success          bool          `json:"success"`
	RemainingBalance string        `json:"remaining_balance"`
	Currency         string        `json:"currency"`
	Timestamp        uint          `json:"timestamp"`
	Error            ErrorResponse `json:"error"`
}

type ErrorResponseV3 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Help    string `json:"help"`
	TraceID string `json:"trace_id"`
}

type Calendar struct {
	StartTime       string `json:"start_time"`
	BlockOutTimeUtc string `json:"block_out_time_utc"`
	SellingPrice    string `json:"selling_price"`
	TOCSellingPrice string `json:"toc_selling_price"`
	Inventory       uint32 `json:"inventory"`
}

type Calendars struct {
	Month    string     `json:"month"`
	Calendar []Calendar `json:"calendars"`
}

// Schedule v3
type Schedule struct {
	SKUID         uint        `json:"sku_id"`
	PublishStatus uint        `json:"publish_status"`
	Currency      string      `json:"currency"`
	Calendars     []Calendars `json:"calendars"`
}

type SchedulesResponse struct {
	Success   bool            `json:"success"`
	Schedules []Schedule      `json:"schedules"`
	Error     ErrorResponseV3 `json:"error"`
}

// Detail Activity v3
type Activity struct {
	ActivityID           uint      `json:"activity_id"`
	Title                string    `json:"title"`
	Subtitle             string    `json:"subtitle"`
	SupportedLanguages   []string  `json:"supported_languages"`
	AvailableDate        string    `json:"available_date"`
	Location             string    `json:"location"`
	AddressDescMultilang string    `json:"Address_desc_multilang"`
	Currency             string    `json:"currency"`
	Price                string    `json:"price"`
	TOCPrice             string    `json:"toc_price"`
	HasMSP               bool      `json:"has_msp"`
	Images               []Image   `json:"images"`
	CityInfo             []CityV3  `json:"city_info"`
	CategoryInfo         Category  `json:"category_info"`
	SectionInfo          []Section `json:"section_info"`
	PackageList          []Package `json:"package_list"`
}

type DetailActivityResponse struct {
	Success  bool          `json:"success"`
	Activity Activity      `json:"activity"`
	Error    ErrorResponse `json:"error"`
}

type Image struct {
	ImageAlt     string `json:"image_alt"`
	ImageDesc    string `json:"image_desc"`
	Width        uint   `json:"width"`
	Height       uint   `json:"height"`
	ImageType    string `json:"image_type"`
	ImageUrlHost string `json:"image_url_host"`
}

type CityV3 struct {
	CityID      uint   `json:"city_id"`
	CityName    string `json:"city_name"`
	CountryID   uint   `json:"country_id"`
	CountryName string `json:"country_name"`
}

type Category struct {
	SubCategoryID    uint   `json:"sub_category_id"`
	SubCategoryNAme  string `json:"sub_category_name"`
	LeafCategoryID   uint   `json:"leaf_category_id"`
	LeafCategoryName string `json:"leaf_category_name"`
}

type Section struct {
	SectionName string  `json:"section_name"`
	Groups      []Group `json:"Groups"`
}

type Group struct {
	GroupName     string `json:"group_name"`
	GroupType     uint   `json:"group_type"`
	GroupTypeName string `json:"group_type_name"`
	RefFieldTag   string `json:"ref_field_tag"`
	Content       string `json:"content"`
}

type Package struct {
	PackageID                 uint      `json:"package_id"`
	PackageName               string    `json:"package_name"`
	PackageMinPax             uint      `json:"package_min_pax"`
	PackageMaxPax             uint      `json:"package_max_pax"`
	CancellationType          int       `json:"cancellation_type"`
	CancellationTypeMultilang string    `json:"cancellation_type_multilang"`
	VoucherUsage              uint      `json:"voucher_usage"`
	VoucherUsageMultilang     string    `json:"voucher_usage_multilang"`
	TimeSlotType              uint      `json:"timeslot_type"`
	Instant                   uint      `json:"instant"`
	TicketType                uint      `json:"ticket_type"`
	SectionInfo               []Section `json:"section_info"`
	SKUList                   []SKU     `json:"sku_list"`
}

type SKU struct {
	SKUID     uint   `json:"sku_id"`
	MinAge    uint   `json:"min_age"`
	MaxAge    uint   `json:"max_age"`
	SKUMaxPax uint   `json:"sku_max_pax"`
	SKUMinPax uint   `json:"sku_min_pax"`
	Title     string `json:"title"`
	Required  bool   `json:"required"`
}
