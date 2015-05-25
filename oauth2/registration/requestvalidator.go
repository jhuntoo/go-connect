package registration

import (
	
)


type Validator interface{
	Validate(*Request) (error)
}

type RequestValidator struct {
	
}

func NewRequestValidator() *RequestValidator {
	return & RequestValidator{}
}


func (p * RequestParser)  Validate(request *Request) (error) {
	return nil	
}

