package utils


type UserError struct{
	Message string
}

func(e *UserError)Error()string{
	err := e.Message

	return err
}