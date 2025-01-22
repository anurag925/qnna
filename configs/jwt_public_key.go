package configs

// import (
// 	"crypto/ecdsa"
// 	"log/slog"
// 	"strings"
// 	"sync"

// 	jwt "github.com/golang-jwt/jwt/v5"
// )

// var jwtPublicKey *ecdsa.PublicKey
// var jwtPublicKeyOnce sync.Once

// func GetJwtPublicKey() *ecdsa.PublicKey {
// 	jwtPublicKeyOnce.Do(func() {
// 		pubKey := strings.ReplaceAll(Get().ES512PublicKey, "\\n", "\n")
// 		key, err := jwt.ParseECPublicKeyFromPEM([]byte(pubKey))
// 		if err != nil {
// 			slog.Error("Error parsing public key", "error", err)
// 			panic(err)
// 		}
// 		jwtPublicKey = key
// 	})
// 	return jwtPublicKey
// }
