package main

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/hex"
    "encoding/pem"
    "fmt"
)

//export Ase256
func Ase256(plaintext string, key string, iv string, blockSize int) string {
    bKey := []byte(key)
    bIV := []byte(iv)
    bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
    block, _ := aes.NewCipher(bKey)
    ciphertext := make([]byte, len(bPlaintext))
    mode := cipher.NewCBCEncrypter(block, bIV)
    mode.CryptBlocks(ciphertext, bPlaintext)
    return base64.StdEncoding.EncodeToString(ciphertext)
}

//export EncryptString
func EncryptString(input string) string {
    publicKeyString := `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAipFm4ybslSPEeZH6mrm8
nPNEQ3+7S698AziRnh9b3XcrL5i+rdGewu0Da0Qdg5ibdfQ1DVHGErueIu0Lelde
//46If0a48bq+exArEln0+gYYXeY8nF0R+6rx2ZUcyaXFjjwPBzYKnJyjWoap/z1
Ex6Nz43cQ8TOD+WYUSio56VjC8fSIfBDnzWRpp6ZhPocCQEyM73iwVh+dfXsgBcK
RLyPEPQ0LHabGute3VqgL895308Hhx7UaUF5dbPbgeN/2tYhWpmO7Bxsx/QerCVL
cklJCOISM3g31UcZUFAQXsY9oG+x7ewliWYRlsbo0Qblm6K2NTuvx4BzOwYKhuQ3
kYSBjkE31gta74Z6RwFAuizoiL/zs90ZN1PFRoPUl4RJQq8s9nPB3OH+vaS4JUml
PbK/2GhnCK1kYukxPlLNH1eCFcP5GurmfTEBULMcX8rmB/ZWJztPcVmGrTw2zF42
kEOk0USP5rJrN71iITrH6wflcNFLIvzst1qBKR/us+wjqtpsKYwYWEdmZGJ9lSx/
QuojxOySjmH1sudQniIZOtBnh9/fZun33nant3P8/pa1qJ018NNYpOOx4fQtIzTd
2I+QjeW1hmS3K/+SsCq9Sb6SPXekd5kCPmPfP6goqDFplK9OEOJ7P/7UiGmuO9Td
ygNwOKTppqa6xCk4XlVb5fkCAwEAAQ==
-----END PUBLIC KEY-----
`

    key, _ := randomHex(32)
    iv, _ := randomHex(16)
    encryptPayload := Ase256(input, key, iv, aes.BlockSize)

    keyStr := key + iv
    encryptKey := encryptStringWithPublicKey(publicKeyString, keyStr)
    if encryptKey == "" {
        return ""
    }
    return encryptPayload + "," + encryptKey
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
    padding := (blockSize - len(ciphertext)%blockSize)
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func randomHex(length int) (string, error) {
    bytes := make([]byte, length/2)
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

func encryptStringWithPublicKey(publicKeyString string, str string) string {
    publicKeyBytes := []byte(publicKeyString)

    block, _ := pem.Decode(publicKeyBytes)
    if block == nil {
        return ""
    }

    publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return ""
    }

    rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
    if !ok {
        return ""
    }

    ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublicKey, []byte(str), nil)
    if err != nil {
        return ""
    }

    base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)
    return base64Ciphertext
}

// main is required to build the shared library
func main() {}
