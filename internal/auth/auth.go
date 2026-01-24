package auth

import (
	"net/http"
	"errors"
	"strings"
)

//GetAPIKey retrieves the API key from the request header.
// The header should contain the field.
// Authorization: APIKey <api_key>

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization header found.")
	}

	header_split := strings.Split(val, " ")
	if len(header_split) != 2 || header_split[0] != "APIKey" {
		return "", errors.New("invalid authorization header format")
	}

	return header_split[1], nil
}