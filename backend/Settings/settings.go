package Settings

import (
	"os"
)

var OPENAIAPIKEY = os.Getenv("OPENAI_API_KEY")
var CHATGPTAPIBASEURL = "https://api.openai.com/v1/"
