package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

//Model Struct
type Track struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not_null"`
	TrackID        string         `json:"track_id" binding:"required" gorm:"uniqueIndex;not null"`
	TrackName      string         `json:"track_name" binding:"required" gorm:"not null"`
	Name           string         `json:"name" gorm:"not null" binding:"required"`
	AdditionalInfo AdditionalInfo `json:"additional_info" gorm:"type:json" binding:"required"`
	UpdatedAt      time.Time      `json:"updated_at"`
	IsDeleted      bool           `json:"is_deleted"`
}

type AdditionalInfo struct {
	Type     string     `json:"type"`
	Features []Features `json:"features"`
}

type Features struct {
	Type       string   `json:"type"`
	Properties string   `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

type Geometry struct {
	Coordinates []int  `json:"coordinates"`
	Type        string `json:"type"`
}

// Seeder
type TrackData struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not_null"`
	TrackID        string    `json:"track_id" binding:"required" gorm:"uniqueIndex;not null"`
	TrackName      string    `json:"track_name" binding:"required" gorm:"not null"`
	Name           string    `json:"name" gorm:"not null" binding:"required"`
	AdditionalInfo string    `json:"additional_info" gorm:"not null;type:json" binding:"required"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"id_deleted"`
}

// Model Struct for Swagger
type Tracks struct {
	TrackID        string         `json:"track_id" binding:"required" gorm:"not null"`
	TrackName      string         `json:"track_name" binding:"required" gorm:"not null"`
	Name           string         `json:"name" gorm:"not null" binding:"required"`
	AdditionalInfo AdditionalInfo `json:"additional_info" gorm:"type:json" binding:"required"`
}

type TrackGet struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not_null"`
	TrackID        string         `json:"track_id" binding:"required" gorm:"not null"`
	TrackName      string         `json:"track_name" binding:"required" gorm:"not null"`
	Name           string         `json:"name" gorm:"not null" binding:"required"`
	AdditionalInfo AdditionalInfo `json:"additional_info" gorm:"type:jsonb" binding:"required"`
	UpdatedAt      time.Time      `json:"updated_at"`
	IsDeleted      bool           `json:"is_deleted"`
}

type Tabler interface {
	TableName() string
}

func (Track) TableName() string {
	return "track_data"
}

func (TrackGet) TableName() string {
	return "track_data"
}

// type JSONB []interface{}

// Value Marshal
func (a AdditionalInfo) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *AdditionalInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
