package seeder

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type seeder func()

type Seed struct {
	Name   string
	Seeder seeder
}

func Run() {
	start := time.Now()

	seeds := []Seed{
		{"user", userSeeder},
		{"class", classSeeder},
	}

	call(seeds)

	fmt.Print("\n\n")
	color.New(color.BlinkSlow, color.BgGreen, color.FgBlack).Printf("Seed finished. (%.2fs)", time.Since(start).Seconds())
	color.Unset()
	fmt.Println()
}

func call(seeds []Seed) {
	gofakeit.Seed(42)

	if name := viper.GetString("name"); name != "" {
		for _, seed := range seeds {
			if seed.Name == name {
				seed.Seeder()
			}
		}
		return
	}

	for _, seeder := range seeds {
		seed(seeder)
	}
}

func seed(seed Seed) {
	start := time.Now()

	color.Green("Seeding %s...", seed.Name)
	seed.Seeder()
	color.Blue("Seeded %s. (%.2f)s", seed.Name, time.Since(start).Seconds())
}
