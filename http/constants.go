package http

type ResponseFormat string

const (
	ResponseFormatJSON ResponseFormat = "json"
	ResponseFormatXML  ResponseFormat = "xml"
)

const (
	DefaultUserAgent = "Agora Chat CLI vX.X.X"
)
