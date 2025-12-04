package model

import "time"

// GeoContinent maps to the geo_continents table.
type GeoContinent struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null;uniqueIndex:uniq_geo_continents_name" json:"name"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

// TableName ensures the struct maps to the exact table name in your SQL.
func (GeoContinent) TableName() string {
	return "geo_continents"
}
