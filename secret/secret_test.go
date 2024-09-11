package secret

import (
	"testing"
)

// TestScanEnv проверяет наличие .env файла по пути `/.env`
func TestScanEnv(t *testing.T) {
	if err := ScanEnv(); err != nil {
		t.Error(err)
	}
}