package tokenizer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Canva/tokenizer/util"
)

func ExampleConfig() {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	tokFile, err := CachedPath("hf-internal-testing/llama-tokenizer", "tokenizer.json")
	if err != nil {
		panic(err)
	}

	f, err := os.Open(tokFile)
	if err != nil {
		panic(err)
	}

	dec := json.NewDecoder(f)

	var config *Config

	err = dec.Decode(&config)
	if err != nil {
		panic(err)
	}

	modelConfig := util.NewParams(config.Model)

	modelType := modelConfig.Get("type", "").(string)
	fmt.Println(modelType)

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Filter out progress message and print only the model type
	output := <-outC
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if !strings.Contains(line, "completed") {
			fmt.Println(line)
		}
	}

	// Output:
	// BPE
}
