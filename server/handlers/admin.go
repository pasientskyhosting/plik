package handlers

import (
	"net/http"

	"github.com/pasientskyhosting/plik/server/common"

	"github.com/pasientskyhosting/plik/server/context"
)

// GetUsers return users
func GetUsers(ctx *context.Context, resp http.ResponseWriter, req *http.Request) {

	// Check authorization
	if !ctx.IsAdmin() {
		ctx.Forbidden("you need administrator privileges")
		return
	}

	pagingQuery := ctx.GetPagingQuery()

	// Get uploads
	users, cursor, err := ctx.GetMetadataBackend().GetUsers("", false, pagingQuery)
	if err != nil {
		ctx.InternalServerError("unable to get users : %s", err)
		return
	}

	pagingResponse := common.NewPagingResponse(users, cursor)
	common.WriteJSONResponse(resp, pagingResponse)
}

// GetServerStatistics return the server statistics
func GetServerStatistics(ctx *context.Context, resp http.ResponseWriter, req *http.Request) {

	// Check authorization
	if !ctx.IsAdmin() {
		ctx.Forbidden("you need administrator privileges")
		return
	}

	// Get server statistics
	stats, err := ctx.GetMetadataBackend().GetServerStatistics()
	if err != nil {
		ctx.InternalServerError("unable to get server statistics : %s", err)
		return
	}

	common.WriteJSONResponse(resp, stats)
}
