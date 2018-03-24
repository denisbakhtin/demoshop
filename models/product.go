package models

//Product type contains product info
type Product struct {
	ID                 uint    `json:"id"`
	Name               string  `json:"name"`
	FeatureTitle       string  `json:"feature_title"`
	FeatureDescription string  `json:"feature_description"`
	Description        string  `json:"description"`
	Price              float64 `json:"price"`
	ImageURL           string  `json:"image_url"`
	VendorID           uint    `json:"vendor_id"`
	Vendor             Vendor  `json:"vendor"`
}
