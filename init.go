package tokenizer

import (
	"fmt"
	"os"
)

var (
	CachedDir       string = "NOT_SETTING"
	tokenizerEnvKey string = "GO_TOKENIZER"
)

func init() {
	// Try TEST_TMPDIR first, then HOME/.cache/tokenizer
	if tmpDir := os.Getenv("TEST_TMPDIR"); tmpDir != "" {
		CachedDir = tmpDir
	} else {
		homeDir := os.Getenv("HOME")
		CachedDir = fmt.Sprintf("/.cache/tokenizer", homeDir)
	}

	initEnv()
}

func initEnv() {
	val := os.Getenv(tokenizerEnvKey)
	if val != "" {
		CachedDir = val
	}

	if _, err := os.Stat(CachedDir); os.IsNotExist(err) {
		if err := os.MkdirAll(CachedDir, 0755); err != nil {
			panic(err)
		}
	}
}