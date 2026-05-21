package richerror

type ErrorCode int

const (
	InvalidCode ErrorCode = iota + 1
	NotFoundCode
	ForbiddenCode
	UnexpectedCode
)

type RichError struct {
	WrappedError error
	Operation    string
	message      string
	code         ErrorCode
	Metadata     map[string]any
}

func (r RichError) Error() string {
	return r.Message()
}

func New(operation string) RichError {
	return RichError{Operation: operation}
}

func (r RichError) WithErr(err error) RichError {
	r.WrappedError = err
	return r
}

func (r RichError) WithMessage(msg string) RichError {
	r.message = msg
	return r
}

func (r RichError) WithCode(code ErrorCode) RichError {
	r.code = code
	return r
}

func (r RichError) WithMetaData(md map[string]any) RichError {
	r.Metadata = md
	return r
}

func (r RichError) Code() ErrorCode {
	if r.code != 0 {
		return r.code
	}

	re, ok := r.WrappedError.(RichError)

	if !ok {
		return 0
	}

	return re.Code()
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	re, ok := r.WrappedError.(RichError)

	if !ok {
		return r.WrappedError.Error()
	}

	return re.Message()
}
