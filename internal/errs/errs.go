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
	ErrEmptyAuthKey             = New("empty auth key")
	ErrInvalidAuthKey           = New("invalid auth key")
	ErrOutOfPageLimit           = New("out of page limit")
	ErrNoSensorData             = New("no sensor data found")
	ErrFailedToCountSensorData  = New("failed to count sensor data")
	ErrFailedToSearchSensorData = New("failed to search sensor data")
)
