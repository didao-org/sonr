package common

import (
	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	cmtcservice "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// identityv1 "github.com/didao-org/sonr/x/identity/types"
	servicev1 "github.com/didao-org/sonr/x/service/types"
)

type clients struct {
	echo.Context
}

func Clients(e echo.Context) *clients {
	return &clients{e}
}

// BankClient returns a new bank client.
func (e *clients) Bank() bankv1beta1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return bankv1beta1.NewQueryClient(cc)
	}
	return nil
}

// CometClient returns a new comet client.
func (e *clients) Comet() cmtcservice.ServiceClient {
	if cc := GrpcClientConn(e); cc != nil {
		return cmtcservice.NewServiceClient(cc)
	}
	return nil
}

// GovClient creates a new gov client.
func (e *clients) Gov() govv1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return govv1.NewQueryClient(cc)
	}
	return nil
}

// // IdentityClient creates a new identity client.
// func (e *clients) Identity() identityv1.QueryClient {
// 	if cc := GrpcClientConn(e); cc != nil {
// 		return identityv1.NewQueryClient(cc)
// 	}
// 	return nil
// }

// ServiceClient creates a new service client.
func (e *clients) Service() servicev1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return servicev1.NewQueryClient(cc)
	}
	return nil
}

// StakingClient creates a new staking client.
func (e *clients) Staking() stakingv1beta1.QueryClient {
	if cc := GrpcClientConn(e); cc != nil {
		return stakingv1beta1.NewQueryClient(cc)
	}
	return nil
}

// TxClient creates a new transaction client.
func (e *clients) Tx() tx.ServiceClient {
	if cc := GrpcClientConn(e); cc != nil {
		return tx.NewServiceClient(cc)
	}
	return nil
}

// GrpcClientConn creates a gRPC client connection.
func GrpcClientConn(e echo.Context) *grpc.ClientConn {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"localhost:9090",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return nil
	}
	return grpcConn
}
