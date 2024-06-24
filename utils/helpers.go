package utils

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

func GetiD(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	str, ok := vars["ID"]
	if !ok {
		return 0, fmt.Errorf("missing ID")
	}

	ID, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("invalid ID")
	}
	return uint(ID), nil
}

func IsValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
