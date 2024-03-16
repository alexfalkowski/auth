package config

// Config for v1.
type Config struct {
	Issuer   string    `yaml:"issuer,omitempty" json:"issuer,omitempty" toml:"issuer,omitempty"`
	Admins   []Admin   `yaml:"admins,omitempty" json:"admins,omitempty" toml:"admins,omitempty"`
	Services []Service `yaml:"services,omitempty" json:"services,omitempty" toml:"services,omitempty"`
}

// Admin for v1.
type Admin struct {
	ID   string `yaml:"id,omitempty" json:"id,omitempty" toml:"id,omitempty"`
	Hash string `yaml:"hash,omitempty" json:"hash,omitempty" toml:"hash,omitempty"`
}

// Service for v1.
type Service struct {
	ID       string `yaml:"id,omitempty" json:"id,omitempty" toml:"id,omitempty"`
	Name     string `yaml:"name,omitempty" json:"name,omitempty" toml:"name,omitempty"`
	Hash     string `yaml:"hash,omitempty" json:"hash,omitempty" toml:"hash,omitempty"`
	Duration string `yaml:"duration,omitempty" json:"duration,omitempty" toml:"duration,omitempty"`
}
