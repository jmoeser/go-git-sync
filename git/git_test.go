package git_test

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jmoeser/go-git-sync/git"
)

func TestCheckout(t *testing.T) {
	dir := git.GetTempDir()
	defer os.RemoveAll(dir)

	checkedOutDir, _, err := git.Checkout("https://github.com/jmoeser/terraform-modules", "master", dir)
	if err != nil {
		t.Error(err)
	}

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
