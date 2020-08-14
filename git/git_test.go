package git

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckout(t *testing.T) {
	dir := GetTempDir()
	checkedOutDir := Checkout("https://github.com/jmoeser/terraform-modules", dir)

	defer os.RemoveAll(checkedOutDir)

	f, err := os.Open(filepath.Join(checkedOutDir, "README.md"))
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)
	// Advance through the first line
	scanner.Scan()
	if !strings.Contains(scanner.Text(), "Terraform Modules") {
		t.Errorf("README does not contain expected text")
	}
}
