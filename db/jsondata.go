package db

// EventJSON is any base stat event to the server
type EventJSON struct {
	UniqueServerID  string                `json:"uniqueServerId"`
	UniqueMachineID string                `json:"uniqueMachineId"`
	UniqueRequestID string                `json:"uniqueRequestId"`
	Event           string                `json:"event"`
	Crashing        bool                  `json:"crashing"`
	Server          ServerJSON            `json:"server"`
	System          SystemJSON            `json:"system"`
	Players         PlayersJSON           `json:"players"`
	Plugins         map[string]PluginJSON `json:"plugins"`
}

// ServerJSON includes all data about a specific server and its settings
type ServerJSON struct {
	Port             int     `json:"port"`
	Software         string  `json:"software"`
	FullVersion      string  `json:"fullVersion"`
	Version          string  `json:"version"`
	Build            int     `json:"build"`
	API              string  `json:"api"`
	MinecraftVersion string  `json:"minecraftVersion"`
	Protocol         int     `json:"protocol"`
	TicksPerSecond   int     `json:"ticksPerSecond"`
	TickUsage        float64 `json:"tickUsage"`
	Ticks            int     `json:"ticks"`
}

// SystemJSON includes all data about a specific system a server is on
type SystemJSON struct {
	OperatingSystem string `json:"operatingSystem"`
	Cores           int    `json:"cores"`
	PhpVersion      string `json:"phpVersion"`
	Machine         string `json:"machine"`
	Release         string `json:"release"`
	Platform        string `json:"platform"`
	MainMemory      int    `json:"mainMemory"`
	TotalMemory     int    `json:"totalMemory"`
	AvailableMemory int    `json:"availableMemory"`
	ThreadCount     int    `json:"threadCount"`
}

// PlayersJSON holds information about the server's current players
type PlayersJSON struct {
	Count       int      `json:"count"`
	Limit       int      `json:"limit"`
	CurrentList []string `json:"currentList"`
	HistoryList []string `json:"historyList"`
}

// PluginJSON contains information about a specific plugin included in the server
type PluginJSON struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Enabled bool   `json:"enabled"`
}
