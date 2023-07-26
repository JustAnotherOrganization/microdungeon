package microdungeon

//go:generate go install github.com/dmarkham/enumer@latest
//go:generate go install github.com/golang/mock/mockgen@latest

type DungeonConfig struct {
	MaxRings         uint32
	MaxUnknownArmour uint32
	PadAlignment     int32
}
