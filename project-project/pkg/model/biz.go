package model

var (
	Normal         = 1
	Personal int32 = 1
)
var AESKey = "sdfgyrhgbxcdgryfhgywertd"

const (
	NoDeleted = iota
	Deleted
)

const (
	NoArchive = iota
	Archive
)

const (
	Open = iota
	Private
	Custom
)

const (
	Default = "default"
	Simple  = "simple"
)

const (
	NoCollected = iota
	Collected
)
