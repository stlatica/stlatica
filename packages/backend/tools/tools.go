//go:build tools
// +build tools

package tools

// バージョン追跡を行いたいgoベースのツールをこのインポートステートメントに記載する
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
import (
	_ "github.com/go-playground/validator/v10"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql"
	_ "golang.org/x/tools/cmd/goimports"
)
