package main

import "prisk/prisk"

type Function string

const (
	GetAreaMapFromNet      Function = "-i"
	GetAreaMapFromFile     Function = "-f"
	MakeDangerAreasMapJson Function = "-m"
	Help                   Function = "-h"
)

func run(f Function) {
	switch f {
	case GetAreaMapFromFile:
		prisk.GetDangerAreaMapFromFile()
	case GetAreaMapFromNet:
		prisk.CreatePRisk().GetAreaList()
	case MakeDangerAreasMapJson:
		prisk.CreatePRisk().GetAll()
	default:
		prisk.CreatePRisk().GetAll()
	}
}
