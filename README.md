# Gemini SDK for Go

## Installation

```bash
go get github.com/sifatul-labs/gemini-go
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/sifatul-labs/gemini-go/gemini"
)

func main() {
    model, err := gemini.NewGenerativeModel(&gemini.GenerativeModelConfig{
        ApiKey: os.Genenv("GEMINI_API_KEY"),
        Model: gemini.AVAILABLE_MODELS.FLASH15,
    })
    if err != nil {
        log.Panicln("Unable to prepare Gemini:", err)
    }
    prompt := gemini.EmptyPrompt()
    resp, err := model.GenerateContent(prompt)
    if err != nil {
        log.Panicln("Unable to generate content:", err)
    }

    fmt.Printf("Prompt text: %s\nGemini response: %s\n",
        prompt.Contents[0].Text, resp.Parts.String())
}
```
