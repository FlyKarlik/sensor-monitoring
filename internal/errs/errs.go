package errs

type Err struct {
	Message string
}

func (e *Err) Error() string {
	return e.Message
}

func New(msg string) *Err {
	return &Err{
		Message: msg,
	}
}

var (
	ErrEmptyAuthKey   = New("empty auth key")
	ErrInvalidAuthKey = New("invalid auth key")
)
