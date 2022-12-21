package config

// Config for v1.
type Config struct {
	Key    Key     `yaml:"key"`
	Admins []Admin `yaml:"admins"`
}

// Key for v1.
type Key struct {
	Public  string `yaml:"public"`
	Private string `yaml:"private"`
}

// Admin for v1.
type Admin struct {
	ID   string `yaml:"id"`
	Hash string `yaml:"hash"`
}
