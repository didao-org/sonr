package motor

import (
	"context"
	"fmt"

	st "github.com/sonr-io/sonr/x/schema/types"
	mt "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func (mtr *motorNodeImpl) QueryWhatIs(ctx context.Context, request mt.QueryWhatIsRequest) (mt.QueryWhatIsResponse, error) {
	return mtr.QueryWhatIsWithSchemaCallback(ctx, request, noSchemaCallback)
}

func (mtr *motorNodeImpl) QueryWhatIsWithSchemaCallback(ctx context.Context, request mt.QueryWhatIsRequest, callback SchemaCallback) (mt.QueryWhatIsResponse, error) {
	resp, err := mtr.schemaQueryClient.WhatIs(ctx, &st.QueryWhatIsRequest{
		Creator: request.Creator,
		Did:     request.Did,
	})
	if err != nil {
		return mt.QueryWhatIsResponse{}, err
	}

	whatIsBytes, err := resp.WhatIs.Marshal()
	if err != nil {
		return mt.QueryWhatIsResponse{}, fmt.Errorf("marshal WhatIs: %s", err)
	}

	// store reference to schema
	mtr.resources.StoreWhatIs(resp.WhatIs, callback)

	return mt.QueryWhatIsResponse{
		WhatIs: whatIsBytes,
	}, nil
}