package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type ClientError struct {
	DisplayMessage string `json:"message"`
	HTTPCode       int    `json:"-"`
}

func (ce *ClientError) Error() string {
	return fmt.Sprintf("an error which will be displayed to the client: %d %v", ce.HTTPCode, ce.DisplayMessage)
}

var UserAlreadyExists = &ClientError{
	DisplayMessage: "user with this login already exists",
	HTTPCode:       http.StatusBadRequest,
}

var InvalidCredentials = &ClientError{
	DisplayMessage: "invalid credentials",
	HTTPCode:       http.StatusUnauthorized,
}

func WriteErrorResponse(w http.ResponseWriter, e error) {
	if unwrapped := errors.Unwrap(e); unwrapped != nil {
		e = unwrapped
	}
	if ce, ok := e.(*ClientError); ok {
		w.WriteHeader(ce.HTTPCode)
		json.NewEncoder(w).Encode(ce)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	log.Printf("ERROR: %v", e.Error())
}
