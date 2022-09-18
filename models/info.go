package models

type Info struct {
	RunningStatus     bool
	Sing_Box_Version  string
	Sing_BoxA_Version string
	Geodata_Version   string
	InBound           InBound
	OutBound_Count    int
	Route_Count       int
	Outbound          []OutBound
	Route             []Route
}
