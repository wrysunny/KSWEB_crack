package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
)

const (
	publickey = `
-----BEGIN PUBLIC KEY-----
MIGJAoGBAOi1YPXrXffIg9GmzA0AF4/In8500QVe2Rq292u0FAOd95MYTIkXvSXL
D6PFAW1Sdp4YOrQ3uWcObXmPio6Clndh1wg+deB2caJDgcyTrnTX8dfqZXRFDgLQ
3gFUfI3gTBUm8lPvdvf7OiNNWs3vt/L7PKwKvKHDn8DKLZ6fRhqbAgMBAAE=
-----END PUBLIC KEY-----`
	privatekey = `
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDotWD16133yIPRpswNABePyJ/OdNEFXtkatvdrtBQDnfeTGEyJ
F70lyw+jxQFtUnaeGDq0N7lnDm15j4qOgpZ3YdcIPnXgdnGiQ4HMk6501/HX6mV0
RQ4C0N4BVHyN4EwVJvJT73b3+zojTVrN77fy+zysCryhw5/Ayi2en0YamwIDAQAB
AoGACuyzl+1Bb2fQbb/g2x4NUnNeso3gCi/WYrStj4wtVwYhdfGxa3YxK7G0EuzQ
EQEgDx1KvuR8Jbm9q0lGfe46K1r/Xa/8MJD/vf8QAbVQpBha1n8swLbeGsUbCoqR
JoVyqyS0Tce+ChxlqKUkNcEe+P29thoKIyjErH+GNhrct9ECQQDsThJy8Zge+Pqs
CJ6aBNPEvtvRq+SSfYaGOyVc+vq2XaIhji2D/tNTB6k5a8ltWlKP5Xpo9qqlWJos
by5syFitAkEA/BqSHJIMjkBS7ldVXF9YxL2XlOBmg5FK0kd40RcXIrtsB0SL7azo
mcs2q0yUH+YwB9J200oNuw2+XoDI+RDBZwJALvE3cwQRXx3A1koED76jvvLXQiiu
iHdNMP8w5e6pvW6OVbIj0pPdsSHVeSWzZvjJa/J/Rbiyn5QhVHBlvZBzJQJAGNmV
pXNQAYWdpxi8tUpAucPmeSpVcIqV0XxyEEoyYZ4P2/eJw3fTxbUeQmxd/Xb3LQ41
4EXgbJvCNBaFuOdJ6QJAUMsbF31tQpAqTIn8y0BTxGojugFclYjIRdpN7R4fLyh3
7kNRPQkzmw0sMQSehDCucdMDy6IutAb1Wc/xlq6rOw==
-----END RSA PRIVATE KEY-----`
)

func EncryptMsg(msg string) string {
	// public encrypt message
	// RSA/None/OAEPWithSHA1AndMGF1Padding
	pubkey := FormatPublickey([]byte(publickey))
	SecretMessage := base64.StdEncoding.EncodeToString(PublicEncrypt(pubkey, []byte(msg)))
	return SecretMessage
}
func DecryptMsg(msg string) string {
	prikey := FormatPrivatekey([]byte(privatekey))
	data, _ := base64.StdEncoding.DecodeString(msg)
	ciphertext := PrivateDecrypt(prikey, data)
	return string(ciphertext)
}

func FormatPublickey(publickey []byte) (public *rsa.PublicKey) {
	block, _ := pem.Decode(publickey)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Println("failed to decode PEM block containing public key.")
	}
	pub, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return pub
}

func PublicEncrypt(publickey *rsa.PublicKey, ciphertext []byte) (SecretMessage []byte) {
	// EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)
	// EncryptOAEP encrypts the given message with RSA-OAEP.
	SecretMessage, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, publickey, ciphertext, []byte(""))
	if err != nil {
		log.Println("Encrypt Msg Failed.", err.Error())
	}
	return SecretMessage
}

func FormatPrivatekey(prikey []byte) (public *rsa.PrivateKey) {
	block, _ := pem.Decode(prikey)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Println("failed to decode PEM block containing private key.")
	}
	pri, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return pri
}

func PrivateDecrypt(prikey *rsa.PrivateKey, SecretMessage []byte) (ciphertext []byte) {
	// DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
	// DecryptOAEP decrypts ciphertext using RSA-OAEP.
	ciphertext, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, prikey, SecretMessage, []byte(""))
	if err != nil {
		log.Println("Decrypt Msg Failed.", err.Error())
	}
	return ciphertext
}
