package model

import (
	"time"

	modelGeoContinent "apirest/domain/geo/model/geo_continent"
)

// GeoRegion maps to the geo_regions table.
type GeoRegion struct {
	ID          uint64       `gorm:"primaryKey;autoIncrement" json:"id"`
	ContinentID uint64       `gorm:"not null;index" json:"continent_id"`
	Name        string       `gorm:"size:64;not null;index" json:"name"`
	CreatedAt   time.Time    `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"not null" json:"updated_at"`

	// Association to GeoContinent. GORM will create the FK constraint
	// with ON DELETE CASCADE / ON UPDATE CASCADE as specified.
	Continent   *modelGeoContinent.GeoContinent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ContinentID;references:ID" json:"continent,omitempty"`
}

// TableName ensures the struct maps to the exact table name in your SQL.
func (GeoRegion) TableName() string {
	return "geo_regions"
}
