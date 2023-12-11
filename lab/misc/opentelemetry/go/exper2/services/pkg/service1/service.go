package service1

import (
	"fmt"
	"errors"
	"context"
	"crypto/ecdsa"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"google.golang.org/protobuf/encoding/protojson"

	pb "null.hexolan.dev/dev/protogen"
)

type rpcService struct {
	pb.UnimplementedAuthServiceServer

	PrivateKey *ecdsa.PrivateKey
	publicKeyJWK *pb.GetJwksResponse
}

func NewRpcService(privateKey *ecdsa.PrivateKey) *rpcService {
	return &rpcService{
		PrivateKey: privateKey,
		publicKeyJWK: nil,
	}
}

func (svc *rpcService) getJwks() (*pb.GetJwksResponse, error) {
	// Assemble the public JWK
	jwk, err := jwk.FromRaw(svc.PrivateKey.PublicKey)
	if err != nil {
		return nil, errors.New("something went wrong parsing public key")
	}

	jwk.Set("use", "sig")  // denote use for signatures
	jwk.Set("alg", fmt.Sprintf("ES%v", svc.PrivateKey.Curve.Params().BitSize))  // ease support for both ES256 and ES512

	// Attempt to marshal to protobuf
	// Convert the JWK to JSON
	jwkBytes, err := json.Marshal(jwk)
	if err != nil {
		return nil, errors.New("something went wrong preparing the JWKs")
	}

	// Convert the JSON to Protobuf
	jwkPB := pb.ECPublicJWK{}
	err = protojson.Unmarshal(jwkBytes, &jwkPB)
	if err != nil {
		return nil, errors.New("something went wrong preparing the JWKs")
	}

	// Succesfully assembled the JWK.
	// Store it for future requests
	svc.publicKeyJWK = &pb.GetJwksResponse{
		Keys: []*pb.ECPublicJWK{&jwkPB},
	}

	return svc.publicKeyJWK, nil
}

func (svc *rpcService) GetJwks(ctx context.Context, req *pb.GetJwksRequest) (*pb.GetJwksResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Info().Any("meta", md).Msg("metadata extract")

	log.Error().Msg("test - try to add to OTEL")

	if svc.publicKeyJWK != nil {
		return svc.publicKeyJWK, nil
	}

	return svc.getJwks()
}