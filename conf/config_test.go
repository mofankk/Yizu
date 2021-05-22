package conf

import "testing"

func TestConfig(t *testing.T) {
	cfg := ServerConfig()
	cfg.SaveToFile()
}
