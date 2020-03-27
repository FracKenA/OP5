package op5

type Payload struct {
	Endpoint string
	Body []byte
	Target string
}

type Server struct {
	Host string
	Username string
	Password string
	AllowInsecure bool
}

type APIResponseError struct {
	Error     string `json:"error"`
	FullError string `json:"full_error"`
}

type CommandResult struct {
	Result string `json:"result"`
}