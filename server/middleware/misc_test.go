package middleware

import (
	"github.com/pasientskyhosting/plik/server/common"
	"github.com/pasientskyhosting/plik/server/context"
	data_test "github.com/pasientskyhosting/plik/server/data/testing"
	"github.com/pasientskyhosting/plik/server/metadata"
)

func newTestingContext(config *common.Configuration) (ctx *context.Context) {
	ctx = &context.Context{}
	config.Debug = true
	ctx.SetConfig(config)
	ctx.SetLogger(config.NewLogger())

	ctx.SetDataBackend(data_test.NewBackend())
	ctx.SetStreamBackend(data_test.NewBackend())

	metadataBackendConfig := &metadata.Config{Driver: "sqlite3", ConnectionString: "/tmp/plik.test.db", EraseFirst: true}
	metadataBackend, err := metadata.NewBackend(metadataBackendConfig)
	if err != nil {
		panic(err)
	}
	ctx.SetMetadataBackend(metadataBackend)

	return ctx
}
