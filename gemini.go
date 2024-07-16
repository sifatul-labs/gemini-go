package gemini

var AVAILABLE_MODELS = struct {
	FLASH15    string
	PRO15      string
	EMBEDDING4 string
}{
	"models/gemini-1.5-flash",
	"models/gemini-1.5-pro",
	"models/text-embedding-004",
}

type GenerativeModelConfig struct {
	ApiKey           string
	Model            string
	Stream           bool
	SafetySettings   []SafetySetting
	GenerativeConfig GenerativeConfig
}

type Prompt struct{}

type ModelResponse struct{}

type GenerativeResponse struct{}

type ChatResponse struct{}

type GenerateContentRequest struct {
	Model            string           `json:"model"`
	Contents         []Content        `json:"contents"`
	SafetySettings   []SafetySetting  `json:"safetySettings"`   // Optional. A list of unique SafetySetting instances for blocking unsafe content.
	GenerativeConfig GenerativeConfig `json:"generativeConfig"` // Optional. Configuration options for model generation and outputs.
}

/*
HARM_BLOCK_THRESHOLD_UNSPECIFIED - Threshold is unspecified.
BLOCK_LOW_AND_ABOVE - Content with NEGLIGIBLE will be allowed.
BLOCK_MEDIUM_AND_ABOVE - Content with NEGLIGIBLE and LOW will be allowed.
BLOCK_ONLY_HIGH - Content with NEGLIGIBLE, LOW, and MEDIUM will be allowed.
BLOCK_NONE - All content will be allowed.
*/
type SafetySetting struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"`
}

type GenerativeConfig struct {
	StopSequences   []string `json:"stopSequences"` // max len 5
	MaxOutputTokens int      `json:"maxOutputTokens"`
	Temperature     float64  `json:"temperature"` // 0.0 to 2.0
	TopP            float64  `json:"topP"`
	TopK            float64  `json:"topK"`
	CandidateCount  int      `json:"candidateCount"` // value should always be 1
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role,omitempty"` // Optional. The producer of the content. Must be either 'user' or 'model'
}

// Part can have either text or inlineData and not both
type Part struct {
	Text       string `json:"text"`
	InlineData []Blob `json:"inlineData"` // only use this field for media data
}

type Blob struct {
	MimeType string `json:"mimeType"` // all the supported file formats - https://ai.google.dev/gemini-api/docs/prompting_with_media?lang=python#supported_file_formats
	Data     string `json:"data"`     // base64 encoded string of raw bytes
}

func NewGenerativeModel(config *GenerativeModelConfig) (any, error) {
	return nil, nil
}

func EmptyPrompt() Prompt {
	return Prompt{}
}

func GenerateContent(prompt Prompt) (GenerativeResponse, error) {
	return GenerativeResponse{}, nil
}

// example
/*
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
*/
func init() {
}
