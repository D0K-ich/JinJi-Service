package gpt

const (
	defaultPort 		= 1234
	defaultAddress 		= "http://localhost:%d/v1/chat/completions"
	defaultModel 		= "TheBloke/Mistral-7B-Instruct-v0.1-GGUF"
	defaultTimeout      = 5
	defaultTemp      	= 0.7
	defaultMaxTokens    = 2048

	defaultContentType 	= "application/json"
)

type RequestGPT struct {
	Model			string			`json:"model"`
	Temperature 	float64 		`json:"temperature"`

	Messages 		MessagesGPT		`json:"messages"`
	MaxTokens		int				`json:"max_tokens"`
}

type MessagesGPT []*MessageGPT
type MessageGPT struct {
	Role 	string	`json:"role"`
	UserId  string	`json:"user_id"`
	Content string	`json:"content"`
}

type ResponseGPT struct {
	Id      	string 		`json:"id"`
	Object  	string 		`json:"object"`
	Model   	string 		`json:"model"`

	Created 	int    		`json:"created"`

	Choices 	Choices		`json:"choices"`
	Usage		*Usage		`json:"usage"`
}

type Choices []*Choice
type Choice struct {
	Index 				int 			`json:"index"`
	Message 			*MessageGPT 	`json:"message"`
	FinishReason 		string			`json:"finish_reason"`
}

type Usage struct {
	PromptTokens     	int 			`json:"prompt_tokens"`
	CompletionTokens 	int 			`json:"completion_tokens"`
	TotalTokens      	int 			`json:"total_tokens"`
}

