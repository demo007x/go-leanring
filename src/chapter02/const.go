package main

import "fmt"

type Weapon int

const (
	Arrow Weapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)

const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main() {
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	fmt.Printf("%d %d %d %d \n",FlagNone, FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b %b \n",FlagNone, FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%s %d", CPU, GPU)
}
