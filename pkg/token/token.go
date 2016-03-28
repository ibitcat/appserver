package token

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

var (
	jwtBackend *JwtBackend = nil
)

// 初始化jwt
func InitJwtBackend() {
	if jwtBackend == nil { //单例
		jwtBackend = &JwtBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
			secret:     []byte("fuckyou"),
		}
	}
}

func GetJwtBackend() *JwtBackend {
	if jwtBackend == nil {
		InitJwtBackend()
	}

	return jwtBackend
}

// 读取rsa私钥
func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open("res/jwt/rsa_private_key.pem")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

// 读取公钥
func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open("res/jwt/rsa_public_key.pem")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
