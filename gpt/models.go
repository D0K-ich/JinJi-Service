package gpt

const ( //default values
	defaultRootPath		= "http://localhost:3001/api/"
	defaultModel 		= "open-mistral-7b"
	defaultTimeout      = 5
	defaultTemp      	= 0.7
	defaultMaxTokens    = 2048
	defaultToken 		= "Bearer MEH45JA-48AMFV0-K4BMSBS-VT239C6"
	defaultMode			= "query | chat"

	defaultContentType 	= "application/json"
)

const ( //EPs
	newMessage = "v1/workspace/test/chat"
)

type RequestGPT struct {
	Messages 		*Message		`json:"messages"`
}

type Message struct {
	Message string `json:"message"`
	Mode    string `json:"mode"`
}

type ResponseGPT struct {
	Id           	string 		`json:"id"`
	Type         	string 		`json:"type"`
	TextResponse 	string 		`json:"textResponse"`
	Error 			string 		`json:"error"`
	Close 			bool   		`json:"close"`

	Sources      	Sources		`json:"sources"`
}

type Sources []*Source
type Source struct {
	Title 			string 		`json:"title"`
	Chunk 			string 		`json:"chunk"`
}
