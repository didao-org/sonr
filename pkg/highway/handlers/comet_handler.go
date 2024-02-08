package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  API Endpoints                                 ||
// ! ||--------------------------------------------------------------------------------||

// CometAPI is a handler for the node module
var CometAPI = cometAPI{}

// cometAPI is a handler for the node module
type cometAPI struct{}

// GetLatestBlock returns the latest block
func (h cometAPI) GetLatestBlock(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.CometClient(r, w).GetLatestBlock(r.Context(), &cmtservice.GetLatestBlockRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetNodeInfo returns the node info
func (h cometAPI) GetNodeInfo(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.CometClient(r, w).GetNodeInfo(r.Context(), &cmtservice.GetNodeInfoRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetSyncing returns the syncing status
func (h cometAPI) GetSyncing(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.CometClient(r, w).GetSyncing(r.Context(), &cmtservice.GetSyncingRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetBlockByHeight returns a block by height
func (h cometAPI) GetBlockByHeight(w http.ResponseWriter, r *http.Request) {
	heightStr := chi.URLParam(r, "height")
	i, _ := strconv.ParseInt(heightStr, 10, 64)
	res, err := middleware.CometClient(r, w).GetBlockByHeight(r.Context(), &cmtservice.GetBlockByHeightRequest{Height: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetLatestValidatorSet returns the latest validator set
func (h cometAPI) GetLatestValidatorSet(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.CometClient(r, w).GetLatestValidatorSet(r.Context(), &cmtservice.GetLatestValidatorSetRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetValidatorSetByHeight returns a validator set by height
func (h cometAPI) GetValidatorSetByHeight(w http.ResponseWriter, r *http.Request) {
	heightStr := chi.URLParam(r, "height")
	i, _ := strconv.ParseInt(heightStr, 10, 64)
	res, err := middleware.CometClient(r, w).GetValidatorSetByHeight(r.Context(), &cmtservice.GetValidatorSetByHeightRequest{Height: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// RegisterRoutes registers the node routes
func (h cometAPI) RegisterRoutes(r chi.Router) {
	r.Get("/block", h.GetLatestBlock)
	r.Get("/block/{height}", h.GetBlockByHeight)
	r.Get("/health", h.GetNodeInfo)
	r.Get("/syncing", h.GetSyncing)
	r.Get("/validatorsets/latest", h.GetLatestValidatorSet)
	r.Get("/validatorsets/{height}", h.GetValidatorSetByHeight)
}
