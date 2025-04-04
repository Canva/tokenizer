package main

import (
	"fmt"
	"log"

	"github.com/Canva/tokenizer"
	"github.com/Canva/tokenizer/model/bpe"
	"github.com/Canva/tokenizer/pretokenizer"
)

func runTest() {
	model, err := bpe.NewBpeFromFiles("model/es-vocab.json", "model/es-merges.txt")
	if err != nil {
		log.Fatal(err)
	}

	tk := tokenizer.NewTokenizer(model)

	bl := pretokenizer.NewBertPreTokenizer()

	tk.WithPreTokenizer(bl)

	sentence := "Mi estas Julien."

	inputSeq := tokenizer.NewInputSequence(sentence)

	en, err := tk.Encode(tokenizer.NewSingleEncodeInput(inputSeq), false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sentence: '%v'\n", sentence)

	fmt.Printf("Tokens: %+v\n", en.GetTokens())

	for _, tok := range en.GetTokens() {
		fmt.Printf("'%v'\n", tok)
	}

}
