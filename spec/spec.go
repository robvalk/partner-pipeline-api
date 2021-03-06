// Package ApiDefinition provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RXYW/bNhD9KwS7b5Mt2U67VN/SZmsDFF3QehiwNoBp6mSxlUiVPDn1Av/34SjZlmzF",
	"dTFvRT/Eokjq8b1399gHLk1RGg0aHY8feCmsKADB+qdbsQD6m4CTVpWojOYxn2bAdFXMwTKTMoVQOIaG",
	"mTR1gAwzYBZclaNj8xUPuKI1Xyqw9KBFATzmJW0ccCczKET9hVRUOfJ4FHBclTRHaYQFWL5eBx7Ie/X3",
	"6WDmBAIrqyFhSm9AlUY7OALJf6IfVtSL672xeIiJRvd48I+lNUuVQMJSBXnimNAJS5QFSesCwpkaWwhk",
	"s49VFE2kn8YIn3+GuB7eLqlHg+FwOPuo/8zAwmbl3pQZk0ITJzPh5IwZy2aEePZRP0KGo3O1iWjO7tAq",
	"veBrOvuGT++UX6019t2G4ZhspRG0J0eUZa6kIDjhJ0cMPbR2/slCymP+JNwZMazfutDvWn+ty/CVZkDv",
	"WCYcM1JW1kLiFWmW0s7XsITclGC9sS39QlXDlRYEQnKF/YZCVYBDUZTkqfsMyEDKkZymshLYvXCs2YIH",
	"HL6KosyJnnE0jgbR88E4mo4u4slFPJ78PIriKOIBr6XlMU8EwoC+wIN9WgOuEkK02/H55Th9KqP5YDKR",
	"zwcXyeiXwVxM5CCJnsmLp+n4EmDSt0+tY3una1iya1gqvUBDshdKvwG9wKxdc7sN3GeVu0N23iiHRAq9",
	"zh3DTNRO33JNipCpqBC7AF4+6QPaDAhrxco/gyi660iQKzal8Z71VZn8WyVz4ZA1+zwq57M4IiG/S87d",
	"8cz8E0gkvFui3KErM+Hewlds1dvcmByE9s7YMLr9caxydt4/4LgPVV1oB4B8jU395H1yX1eF0AMLIhHz",
	"HJpqpH0DJo2te0PiW/Hr6fSWORRYkTFAVwWPP/AXInkHXypw1GfeGvzNVJrYv9EIVov8Pdgl2BrXXVuU",
	"zsIDOxTgXG9o7eFtvSRnkIX9ETr6vzW++bK0wfZNeWlI6dR4BRVu6q4RnF3d3rDliAd8CdbVsEbDaBgR",
	"cFOCFqXiMZ/4oYCXAjOvQph0PLMA7xB69j31JqHIAWFl1jJX0InyD/t0/K7zVZOQ25S6z4yD+sSFQJmB",
	"87wsRV4Bg69CYv5Ynvs/x/IiOB0AdYDvBoB1ezgKoK9ediSF/rJz4jx/UThhrr8frO/2wnIcRWeLyJbk",
	"PTn5F1hDcV8Y2+rRrnG0n95ccPq/soUddgPeJ21VFMKuKBSM+UxXF9bxH4oFOa8Vw3frgJfG9dj3pQ/T",
	"3UwizNf4C5Oszs8VUbU+EGX033xoX5NdVFL4uEpKcC6t8nzF5spi5kMoA5E09f7G1BgOe9qtwIw6LBWJ",
	"hvt8NWguJazN5KM10W51oShVuBy1Wk140tWjOd8ZTPSCzs5EB3qfhdZBuyGGD9vfN8m6pigHhEOLXfvx",
	"rsU68l8cEnxEKpcZPFsB1di+efSgv/e/AjxyrOj/dvV5e8srwBN42cs6nxCUn7uAaNmk6S6K/ssQo63g",
	"WG7c+SCW2SHtf/gb4w/uWdEP71mFSVSq4HyK17x+uw/QIn9JrDWvbM5jniGWcRhGQ/8vvowunzatja/v",
	"1v8EAAD//8QKO8LuEAAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
