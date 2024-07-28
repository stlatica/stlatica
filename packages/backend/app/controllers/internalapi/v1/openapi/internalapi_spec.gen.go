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
