package response

const JSONSerializationError = "unable to serialize the JSON response"
const JSONParseError = "unable to parse the JSON request body"
const JSONEmptyBodyError = "request body must no be empty"
const Unexpected = "unexpected error occurred"
const BadRequestError = "bad request"

var errorCodes = map[string]string{
	JSONSerializationError: "ERR_JSON_100",
	JSONEmptyBodyError:     "ERR_JSON_101",
	BadRequestError:        "ERR_BAD_REQUEST_400",
	Unexpected:             "ERR_UNEXPECTED_500",
}
