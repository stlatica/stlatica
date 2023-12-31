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

	"H4sIAAAAAAAC/+xabW/byBH+K8K2QFqUMilHdu/0TTnpXKE+2ZWt+9DYEFbUWmbKty5Xrl1DgEk1qXFO",
	"4OAaxxfgirRpXtw6UQsEbdM6aX7MWnL8L4pdkhIlrpy4cYNDICBARM7u7LMz8zwccr0JVMuwLROZxAGZ",
	"TWBDDA1EEPavdEgqWo39rCFHxZpNNMsEGW6g3tdvnn3T2X5E3XvU26Hufeq51G0fv9jq/maXuu1CbskE",
	"EkDr0LB1FMxi7iSgcSeQrAIJmNAYNGL0y4aGUQ1kCG4gCTjqKjIgA0E2bDbUIVgz66DZlADRDKRrJhKi",
	"DI3nRxp1K0Q7OOA8iBsOwkK0zHB+pKE7Icq+8d0RNkMjr4A8xhYuIce2TAfxksCWjTDREDerVo3fRWbD",
	"AJmr4Eo2Vynlf1bOLywCCXxRWFgoFGcq89lS9ov8Yr4EJFAuZsuLP5krFX6ezwEJFOcWK5/PlYvsd6G4",
	"mC8Vs7OVhXzpy3ypki+V5tgUdln4LF8pF7NfZguz2SuzebAcjcGoiUNbk4CBHAfWfcS92YhtMRGaJEHK",
	"+sG76u+472m5N96qXkMqYavM65CIKQOkWPxMgswRo6nb7n77tHv3twMJp+7vqHeTun+l3t9p6zlt7dPW",
	"IXXb/ownrHC2PNHmVYwgQbUKHLnczeNX33a3b1N3j69wv7v/qHvP+0FhYe6TaSX1wwEYk8pkOqmkkkpq",
	"UVEy/N+PlE8zigIksGJhg60CapCgJKOKCI9mwLofBRGWw+4/j46PvulhOblz1GkxBpRLs6zYCTKcwTyu",
	"EmI7GVleX1+XNZMgbEJdXkvJ/jpySoQhuAExhhvsOiJ438doBWTA9+S+QMoBMWSW4UIuVht9CQvzOhD1",
	"UbVSyF2UwE5NKeiTtKIk0eSn1WQ6VUsn4Y9T08l0enp6aiqdVhSen1gYGIp5y/kuVm1ELs/KSNlBWJCR",
	"fh5CP6IksMliPY5tfcXSdetXCDsV1WqIQhAO6O79LRqASaW3LqvNOsJsYX+wZtbP9qaZ9Xdyp6ls0rAP",
	"2npMWy+p9w/qtqn3gHoPqfectp4PMUpMJLQeI5LDx1WgSiw8YZt1Ibedit2o6poah9O5/vT07g4Hcki9",
	"x7wYtnkJeLzId6j7Z+reoO5OYmlpyQwBJKCtUbdtQLMBdX0ja9vYWkPO52E+qHvQuX3z+MUWddudrYfU",
	"PaDuX6i3M8gQ/+EXwK1alo6gyfA6DcOAeOMtwWNF+4y29qj3J3azdTgQtmhcLqCS/Qn+g3wY1cmdo9Pf",
	"Pzh+8RV1D9/88eDk4b8DnfR2eoA7t2+dA94Qb/qNQw9EP0hBoUkxMsQLOloJo5gnkr/37YUETY1mrljx",
	"hRyiQ6KpMBE+MVidMQca8aMW2CuhveLb1xB2fA/KRGpCYXuxbGQyYwZcnlAmLjPpgGSVq4boebTJ/69o",
	"tSYbUUcC9tcRSfBRnAlV6KDpdOfVg87LXeo+efP6Ti/pPA5MpiCbWaiBDJhBpGD4DU20p78qKKVOa7eQ",
	"GwhkCE3cVUas795WLrPBfhvJQzKpKEOPEu5Wrmsr7KLvqNdMVDeIuD3zJ16zUf1/m8k07FwTm9JQEOd+",
	"ygogHdsStG1dU3lS5GuOL879Zc5SgcG+W7BkFdYSLPrIIb5O8i6WdaeZJTOZuBTpxC/xG7Fm/NKS6aNO",
	"CfhnwgZZtbD2a1QTuo+28T1H6Q+3fdMiiRWrYYrR9d4qAmhTHzIzPSVxEF5DOOGDE8EUvrf0IF/+cJAZ",
	"Uk1FiYYJ16Cmw6qOhIAFb2IcbvQJ2lctxhxYZ6IT9vrLbOSAGLL+z//UIGw/WVOaYGMmYgLHTPN+dxrw",
	"4IpV27iwmPVaYkG4ul/tnRy87ty43mn/CzRjwibg02f+W8DEWCbGMjGWCS4TA9QOdYITeoRKyJvB+3XT",
	"L0QdEUFrnOP3R2iGbwxUY6gtEkWmP0QO3+0FjUw6jqJoJYKUjCk/pvyY8j7lI9xMVDcShZyA+pL4VaiE",
	"SAObzghezyBy0aRWLrSTEIa2oarIccb6MNaHsT5wfZhB5GxxGO4LwlM4R96MHMid/UElHMhh8tW0GnXb",
	"p/dunf7hBnWfBN+wT6/f6mzvU+/rzu7dzn/2/W9Q4u8si4HHc8tP9BDxvSWodyLydi0aPvf4KLVpLClj",
	"SRlkfERPepSNa0rDYdTcDL4ln60lwRFRTBHKwdHRudQg/Hr9f21GOLIx4ceE/5gJH9AyJDuv+WX/L1t4",
	"UETnIDbWHCNhWOovgsgBCTSwHhyHZmRZt1Sor1oOyaSVlAKay83/BgAA//9PG2Y2PCUAAA==",
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
