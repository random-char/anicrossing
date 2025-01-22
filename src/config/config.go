package config

import rl "github.com/gen2brain/raylib-go/raylib"

type Config struct {
	ScreenResolution rl.Vector2
}

func LoadConfig() *Config {
	//todo add config storage

	return &Config{
		ScreenResolution: rl.NewVector2(1000, 500),
	}
}
