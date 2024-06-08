package benchmark

import "github.com/zacksfF/TempoScale-Distributed-Infrastructure-for-Time-Series-Data/internal/config"

const testFile2 = "./test1-zlib.orc"

type noopMembreship int

func (m noopMembreship) Memmber() []string {
	return []string{
		"127.0.0.1:9876",
	}
}

type benchMockConfig struct {
	dir string
}


func (m *benchMockConfig) Configure(c *config.Config) error


