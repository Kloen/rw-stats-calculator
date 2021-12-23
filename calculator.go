package main

import (
	"fmt"
	"math"
)

type Professions struct {
	Assassin  Increment
	Compeller Increment
	Destroyer Increment
	Guardian  Increment
	Healer    Increment
	Mage      Increment
	Witch     Increment
}

type Increment struct {
	Hp   [4]float64
	Atk  [4]float64
	Def  [4]float64
	MDef [4]float64
}

type Stats struct {
	Hp   float64
	Atk  float64
	Def  float64
	MDef float64
}

type Doll struct {
	Name       string
	Level      int
	Stats      Stats
	StatGrowth Stats
}

var maxAscLevel = [5]int{30, 40, 60, 70, 90}

var professions = Professions{
	Assassin:  Increment{},
	Compeller: Increment{},
	Destroyer: Increment{},
	Guardian: Increment{
		Hp:   [4]float64{1.375, 1.43, 1.4, 0},
		Atk:  [4]float64{1.375, 1.43, 1.25, 0},
		Def:  [4]float64{1.069, 1.22, 1.17, 0},
		MDef: [4]float64{1.114, 1.22, 1.14, 0},
	},
	Healer: Increment{},
	Mage:   Increment{},
	Witch:  Increment{},
}

func calculateStat(base float64, add float64, targetLevel int, increment [4]float64) float64 {
	var result = base
	var currentLevel = 0
	var ascension = 0

	for currentLevel < targetLevel {
		for currentLevel < maxAscLevel[ascension] && currentLevel < targetLevel {
			result += add
			currentLevel++
		}
		if currentLevel < targetLevel && currentLevel == maxAscLevel[ascension] {
			result = result * increment[ascension]
			ascension++
		}
	}

	return math.Round(result)
}

func main() {
	doll := Doll{
		Name:  "Caledonia",
		Level: 0,
		Stats: Stats{
			Hp:   900,
			Atk:  60,
			Def:  660,
			MDef: 396,
		},
		StatGrowth: Stats{
			Hp:   90,
			Atk:  6,
			Def:  4,
			MDef: 2.4,
		},
	}

	target := Doll{
		Level: 31, // Ascension 1 - 1
		Stats: Stats{
			Hp:   5040,
			Atk:  336,
			Def:  838,
			MDef: 524,
		},
	}

	targetStat := target.Stats.Hp
	result := calculateStat(doll.Stats.Hp, doll.StatGrowth.Hp, target.Level, professions.Guardian.Hp)
	fmt.Println("Expected:", targetStat)
	fmt.Println("Result:", result)
	fmt.Println("HP error:", targetStat-result)
	fmt.Println()

	targetStat = target.Stats.Atk
	result = calculateStat(doll.Stats.Atk, doll.StatGrowth.Atk, target.Level, professions.Guardian.Atk)
	fmt.Println("Expected:", targetStat)
	fmt.Println("Result:", result)
	fmt.Println("Atk error:", targetStat-result)
	fmt.Println()

	targetStat = target.Stats.Def
	result = calculateStat(doll.Stats.Def, doll.StatGrowth.Def, target.Level, professions.Guardian.Def)
	fmt.Println("Expected:", targetStat)
	fmt.Println("Result:", result)
	fmt.Println("Def error:", targetStat-result)
	fmt.Println()

	targetStat = target.Stats.MDef
	result = calculateStat(doll.Stats.MDef, doll.StatGrowth.MDef, target.Level, professions.Guardian.MDef)
	fmt.Println("Expected:", targetStat)
	fmt.Println("Result:", result)
	fmt.Println("MDef error:", targetStat-result)
}
