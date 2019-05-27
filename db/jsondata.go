package db

// EventJSON is any base stat event to the server
type EventJSON struct {
	UniqueServerID  string                 `json:"uniqueServerId,omitempty"`
	UniqueMachineID string                 `json:"uniqueMachineId,omitempty"`
	UniqueRequestID string                 `json:"uniqueRequestId,omitempty"`
	Event           string                 `json:"event,omitempty"`
	Crashing        bool                   `json:"crashing,omitempty"`
	Server          ServerJSON             `json:"server,omitempty"`
	System          SystemJSON             `json:"system,omitempty"`
	Players         PlayersJSON            `json:"players,omitempty"`
	Plugins         map[string]*PluginJSON `json:"plugins,omitempty"`
}

// ServerJSON includes all data about a specific server and its settings
type ServerJSON struct {
	Port             int     `json:"port,omitempty"`
	Software         string  `json:"software,omitempty"`
	FullVersion      string  `json:"fullVersion,omitempty"`
	Version          string  `json:"version,omitempty"`
	Build            int     `json:"build,omitempty"`
	API              string  `json:"api,omitempty"`
	MinecraftVersion string  `json:"minecraftVersion,omitempty"`
	Protocol         int     `json:"protocol,omitempty"`
	TicksPerSecond   int     `json:"ticksPerSecond,omitempty"`
	TickUsage        float64 `json:"tickUsage,omitempty"`
	Ticks            int     `json:"ticks,omitempty"`
}

// SystemJSON includes all data about a specific system a server is on
type SystemJSON struct {
	OperatingSystem string `json:"operatingSystem,omitempty"`
	Cores           int    `json:"cores,omitempty"`
	PhpVersion      string `json:"phpVersion,omitempty"`
	Machine         string `json:"machine,omitempty"`
	Release         string `json:"release,omitempty"`
	Platform        string `json:"platform,omitempty"`
	MainMemory      int    `json:"mainMemory,omitempty"`
	TotalMemory     int    `json:"totalMemory,omitempty"`
	AvailableMemory int    `json:"availableMemory,omitempty"`
	ThreadCount     int    `json:"threadCount,omitempty"`
}

// PlayersJSON holds information about the server's current players
type PlayersJSON struct {
	Count       int      `json:"count,omitempty"`
	Limit       int      `json:"limit,omitempty"`
	CurrentList []string `json:"currentList,omitempty"`
	HistoryList []string `json:"historyList,omitempty"`
}

// PluginJSON contains information about a specific plugin included in the server
type PluginJSON struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}
