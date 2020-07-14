package lodr

import (
	"testing"
	"time"
)

// TestConfig structure
type TestConfig struct {
	Name  string  `yaml:"yamlName" json:"jsonName" cmd:"app.name" env:"NAME"`
	Float float64 `json:"jsonFloat" cmd:"app.thefloat"`
	DB    struct {
		Host    string        `yaml:"yamlHost" json:"jsonHost" cmd:"app.db.host"`
		Port    int           `yaml:"yamlPort" json:"jsonPort" cmd:"app.db.port"`
		Options []string      `yaml:"yamlOptions" json:"jsonOptions" cmd:"app.db.options"`
		TLS     bool          `yaml:"yamlTLS" json:"jsonTLS" cmd:"tls" env:"IS_TLS"`
		Timeout time.Duration `yaml:"yamlTimeout" json:"jsonTimeout" cmd:"timeout"`
	}
	Log struct {
		Formats []struct {
			Name   string `yaml:"yamlLogFormatName" json:"jsonLogFormatName"`
			Pretty bool
		}
	}
}

func TestLoad(t *testing.T) {
	var tc TestConfig
	tc.Name = "test"
	config := Load(tc)
	if config.object == nil {
		t.Error("config.object should be 'test'")
	}
}
