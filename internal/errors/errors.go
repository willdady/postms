package errors

type NotFound struct{}

func (err *NotFound) Error() string {
	return "Not found"
}

type DeleteIsMissingID struct{}

func (err *DeleteIsMissingID) Error() string {
	return "Can not delete. Struct missing ID."
}

type CursorDecodingError struct{}

func (err *CursorDecodingError) Error() string {
	return "Unable to decode cursor"
}
