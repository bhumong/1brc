package first

import (
	"testing"

	"github.com/bhumong/1brc/filepath"
)

func TestCalculateAll(t *testing.T) {
	CalculateAll(filepath.GetPath())
}
