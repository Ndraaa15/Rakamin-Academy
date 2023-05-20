package errors

type Errors struct {
	Type    string
	Code    int64
	Message string
}

func (err *Errors) Error() string {
	return err.Message
}

func NewError(errType string, code int64, message string) error {

	errors := &Errors{
		Type:    errType,
		Code:    code,
		Message: message,
	}

	return errors

}
