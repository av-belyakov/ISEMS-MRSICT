package datamodels

import (
	"ISEMS-MRSICT/moduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction"
)

//ChannelsListInteractingModules содержит список каналов для межмодульного взаимодействия
type ChannelsListInteractingModules struct {
	ChannelsModuleDataBaseInteraction  moduledatabaseinteraction.ChannelsModuleDataBaseInteraction
	ChannelsModuleAPIRequestProcessing moduleapirequestprocessing.ChannelsModuleAPIRequestProcessing
}
