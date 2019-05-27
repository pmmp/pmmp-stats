package web

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pmmp/pmmp-stats/db"
)

func handlePayload(dataBase *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var payloadData db.EventJSON
		err := c.BindJSON(&payloadData)

		if err != nil {
			return
		}

		var data db.Server

		serverData := db.Server{
			MachineData: db.Machine{
				UniqueMachineID: payloadData.UniqueMachineID,
				OperatingSystem: payloadData.System.OperatingSystem,
				Cores:           payloadData.System.Cores,
				PhpVersion:      payloadData.System.PhpVersion,
				Machine:         payloadData.System.Machine,
				Release:         payloadData.System.Release,
				Platform:        payloadData.System.Platform,
				MainMemory:      payloadData.System.MainMemory,
				TotalMemory:     payloadData.System.TotalMemory,
				AvailableMemory: payloadData.System.AvailableMemory,
				ThreadCount:     payloadData.System.ThreadCount,
			},
			UniqueMachineID:    payloadData.UniqueMachineID,
			UniqueServerID:     payloadData.UniqueServerID,
			Port:               payloadData.Server.Port,
			Software:           payloadData.Server.Software,
			FullVersion:        payloadData.Server.FullVersion,
			Version:            payloadData.Server.Version,
			Build:              payloadData.Server.Build,
			API:                payloadData.Server.API,
			MinecraftVersion:   payloadData.Server.MinecraftVersion,
			Protocol:           payloadData.Server.Protocol,
			TicksPerSecond:     payloadData.Server.TicksPerSecond,
			TickUsage:          payloadData.Server.TickUsage,
			Ticks:              payloadData.Server.Ticks,
			PlayerCount:        payloadData.Players.Count,
			PlayerLimit:        payloadData.Players.Limit,
			CurrentPlayerList:  strings.Join(payloadData.Players.CurrentList, ","),
			HistoricPlayerList: strings.Join(payloadData.Players.HistoryList, ","),
			Crashing:           payloadData.Crashing,
			LastEvent:          payloadData.Event,
			LastRequestID:      payloadData.UniqueRequestID,
		}

		for _, plugin := range payloadData.Plugins {
			serverData.Plugins = append(serverData.Plugins, db.Plugin{
				Name:    plugin.Name,
				Version: plugin.Version,
				Enabled: plugin.Enabled,
			})
		}

		// dataBase.Preload("MachineData", db.Machine{
		// 	UniqueMachineID: payloadData.UniqueMachineID,
		// }).Preload("Plugins").Where(db.Server{
		// 	UniqueServerID:  payloadData.UniqueServerID,
		// 	UniqueMachineID: payloadData.UniqueMachineID,
		// }).First(&data)

		dataBase.Where(db.Server{
			UniqueServerID:  payloadData.UniqueServerID,
			UniqueMachineID: payloadData.UniqueMachineID,
		}).Assign(serverData).FirstOrCreate(&data)
	}
}
