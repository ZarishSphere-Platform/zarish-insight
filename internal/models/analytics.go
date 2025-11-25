package models

import "time"

// Dashboard represents an analytics dashboard
type Dashboard struct {
	BaseModel

	Name        string `gorm:"size:255;not null;index" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Layout      string `gorm:"type:jsonb" json:"layout"` // JSON layout configuration
	IsPublic    bool   `gorm:"default:false" json:"is_public"`

	CreatedBy uint `gorm:"index;not null" json:"created_by"`
}

// Report represents a generated report
type Report struct {
	BaseModel

	Name       string `gorm:"size:255;not null;index" json:"name"`
	Type       string `gorm:"size:100;not null" json:"type"` // financial, clinical, operational
	Parameters string `gorm:"type:jsonb" json:"parameters"`
	Data       string `gorm:"type:jsonb" json:"data"`

	GeneratedBy uint      `gorm:"index;not null" json:"generated_by"`
	GeneratedAt time.Time `gorm:"index" json:"generated_at"`
}

// Metric represents a tracked metric
type Metric struct {
	BaseModel

	Name      string    `gorm:"size:255;not null;index" json:"name"`
	Category  string    `gorm:"size:100" json:"category"`
	Value     float64   `json:"value"`
	Unit      string    `gorm:"size:50" json:"unit"`
	Timestamp time.Time `gorm:"index;not null" json:"timestamp"`

	Metadata string `gorm:"type:jsonb" json:"metadata"`
}

// TableName overrides
func (Dashboard) TableName() string {
	return "dashboards"
}

func (Report) TableName() string {
	return "reports"
}

func (Metric) TableName() string {
	return "metrics"
}
