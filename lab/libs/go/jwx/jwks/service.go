package main

import (
	"os"
	"errors"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"google.golang.org/protobuf/encoding/protojson"

	pb "null.hexolan.dev/dev/protogen"
)

type rpcService struct {
	RSAPriv *rsa.PrivateKey

	pb.UnimplementedTestServiceServer
}

func loadRSAKeyFile() []byte {
	keyBytes, err := os.ReadFile("./keys/private.pem")
	if err != nil {
		log.Panic().Err(err).Msg("failed to load RSA file")
	}

	return keyBytes
}

func loadRSAKey() *rsa.PrivateKey {
	keyBytes := loadRSAKeyFile()

	block, _ := pem.Decode(keyBytes)
	rsaPriv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Panic().Err(err).Msg("failed to parse RSA key")
	}
	// ^ may not be needed
	// may be able to parse directly
	// https://pkg.go.dev/github.com/lestrrat-go/jwx/v2/jwk#ParseKey

	return rsaPriv
}

func newRpcService() *rpcService {
	// load RSA private key
	rsaPriv := loadRSAKey()

	return &rpcService{
		RSAPriv: rsaPriv,
	}
}

func (svc rpcService) GetJwks(ctx context.Context, req *pb.GetJwksRequest) (*pb.GetJwksResponse, error) {
	rsaPublic := svc.RSAPriv.PublicKey

	// https://github.com/lestrrat-go/jwx/blob/v2.0.18/docs/04-jwk.md#parse-a-key
	jwkPublic, err := jwk.FromRaw(rsaPublic)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse rsa.PublicKey to jwk.Key")
		return nil, errors.New("bruh")
	}

	jwkPublicBytes, err := json.Marshal(jwkPublic)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal jwk.Key")
		return nil, errors.New("bruh")
	}

	jwkPublicProto := pb.JWK{}
	err = protojson.Unmarshal(jwkPublicBytes, &jwkPublicProto)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal jwk.Key to protobuf")
		return nil, errors.New("bruh")
	}

	log.Info().Any("jwk-public", jwkPublic).Msg("")
	log.Info().Bytes("jwk-public-bytes", jwkPublicBytes).Msg("")
	/*
	{"level":"info","jwk-public":{"e":"AQAB","kty":"RSA","n":"wsTMHsjhH6LKOr98ZusWZRCumkk3W6NfWY-_D3mnvsryaGNnHc1v6S2QHy-Qtp1CeADeWYHUKfI8xUyUOr9Y9lnEZ5J9RCETD1ihZgZhakjG9jDdzjtpQBGpH9qYeOz4punN1vJ2V4LcefLiBMdlF9Bb5vo3_Q6q5r598rTzCd6smzkJf5V6cFZradVWajlO16UTi3IkuTzQ_-CXi2RoAY5W8ja6r3bet0moMwP0pSh2KgTDQrPE1GHjV9DmshIzC7wePoF3bjPM9nsiRMy4B69KGf5C6Bvug4IHXm-9FYcZhOAzIR6N__wYQDuUi1QxwC5OdmW5_hhd4sogg5jACQ"},"time":"2023-12-06T01:54:36Z"}
	{"level":"info","jwk-public-bytes":"{\"e\":\"AQAB\",\"kty\":\"RSA\",\"n\":\"wsTMHsjhH6LKOr98ZusWZRCumkk3W6NfWY-_D3mnvsryaGNnHc1v6S2QHy-Qtp1CeADeWYHUKfI8xUyUOr9Y9lnEZ5J9RCETD1ihZgZhakjG9jDdzjtpQBGpH9qYeOz4punN1vJ2V4LcefLiBMdlF9Bb5vo3_Q6q5r598rTzCd6smzkJf5V6cFZradVWajlO16UTi3IkuTzQ_-CXi2RoAY5W8ja6r3bet0moMwP0pSh2KgTDQrPE1GHjV9DmshIzC7wePoF3bjPM9nsiRMy4B69KGf5C6Bvug4IHXm-9FYcZhOAzIR6N__wYQDuUi1QxwC5OdmW5_hhd4sogg5jACQ\"}","time":"2023-12-06T01:54:36Z"}
	*/
	
	/*
	beautified (before protobuf):
	{
		"e": "AQAB",
		"kty": "RSA",
		"n": "wsTMHsjhH6LKOr98ZusWZRCumkk3W6NfWY-_D3mnvsryaGNnHc1v6S2QHy-Qtp1CeADeWYHUKfI8xUyUOr9Y9lnEZ5J9RCETD1ihZgZhakjG9jDdzjtpQBGpH9qYeOz4punN1vJ2V4LcefLiBMdlF9Bb5vo3_Q6q5r598rTzCd6smzkJf5V6cFZradVWajlO16UTi3IkuTzQ_-CXi2RoAY5W8ja6r3bet0moMwP0pSh2KgTDQrPE1GHjV9DmshIzC7wePoF3bjPM9nsiRMy4B69KGf5C6Bvug4IHXm-9FYcZhOAzIR6N__wYQDuUi1QxwC5OdmW5_hhd4sogg5jACQ"
	}
	*/

	return &pb.GetJwksResponse{
		Keys: []*pb.JWK{&jwkPublicProto},
	}, nil
}