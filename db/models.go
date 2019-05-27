package db

import (
	"time"
)

// DefaultModel defines the default model types
type DefaultModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Plugin contains a DB model containing plugin information
type Plugin struct {
	DefaultModel
	Name    string `gorm:"PRIMARY_KEY" sql:"index"`
	Version string `gorm:"PRIMARY_KEY" sql:"index"`
	Enabled bool   `gorm:"PRIMARY_KEY" sql:"index"`
}

// Machine contains a DB model containing machine information
type Machine struct {
	DefaultModel
	UniqueMachineID string `gorm:"PRIMARY_KEY" sql:"index"`
	OperatingSystem string
	Cores           int
	PhpVersion      string
	Machine         string
	Release         string
	Platform        string
	MainMemory      int
	TotalMemory     int
	AvailableMemory int
	ThreadCount     int
}

// Server contains a DB model containing server information
type Server struct {
	DefaultModel
	UniqueServerID     string  `gorm:"PRIMARY_KEY" sql:"index"`
	UniqueMachineID    string  `gorm:"PRIMARY_KEY" sql:"index"`
	MachineData        Machine `gorm:"foreignkey:UniqueMachineID;association_foreignkey:unique_machine_id"`
	Port               int
	Software           string
	FullVersion        string
	Version            string
	Build              int
	API                string
	MinecraftVersion   string
	Protocol           int
	TicksPerSecond     int
	TickUsage          float64
	Ticks              int
	PlayerCount        int
	PlayerLimit        int
	CurrentPlayerList  string
	HistoricPlayerList string
	Plugins            []Plugin `gorm:"many2many:server_plugins;"`
	Crashing           bool
	LastEvent          string
	LastRequestID      string
}
