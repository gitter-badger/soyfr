package common

// The Response struct implements api2go.Responder
type Response struct {
	Res  interface{}
	Code int
}

// Metadata returns additional meta data
func (r Response) Metadata() map[string]interface{} {
	return map[string]interface{}{
		"author":  "The manyminds crew",
		"license": "MIT",
	}
}

// Result returns the actual payload
func (r Response) Result() interface{} {
	return r.Res
}

// StatusCode sets the return status code
func (r Response) StatusCode() int {
	return r.Code
}
