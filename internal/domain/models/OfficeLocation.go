package models

type OfficeLocationInfo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    int     `json:"radius"`
}
