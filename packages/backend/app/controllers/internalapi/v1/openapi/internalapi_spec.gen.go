// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"1UlUAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
