//go:build tools
// +build tools

package tools

// バージョン追跡を行いたいgoベースのツールをこのインポートステートメントに記載する
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
import (
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "github.com/volatiletech/sqlboiler/v4"
)
