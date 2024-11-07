package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"crypto/sha256"
)

func main() {
	// 4.1生成公钥私钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// 4.2 将生成的私钥和公钥导出到文本文件中，
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyFile, err := os.Create("privateKey.pem")
	if err != nil {
		fmt.Println("Error creating private key file:", err)
		return
	}
	defer privateKeyFile.Close()
	if err := pem.Encode(privateKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}); err != nil {
		fmt.Println("Error encoding private key:", err)
		return
	}
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("Error marshalling public key:", err)
		return
	}
	publicKeyFile, err := os.Create("publicKey.pem")
	if err != nil {
		fmt.Println("Error creating public key file:", err)
		return
	}
	defer publicKeyFile.Close()
	if err := pem.Encode(publicKeyFile, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}); err != nil {
		fmt.Println("Error encoding public key:", err)
		return
	}

	// Read the private and public keys from the files
	privateKeyPEM, err := ioutil.ReadFile("privateKey.pem")
	if err != nil {
		fmt.Println("Error reading private key file:", err)
		return
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		fmt.Println("Failed to parse PEM block containing the key")
		return
	}
	privateKeyFromBytes, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return
	}

	publicKeyPEM, err := ioutil.ReadFile("publicKey.pem")
	if err != nil {
		fmt.Println("Error reading public key file:", err)
		return
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	if publicKeyBlock == nil || publicKeyBlock.Type != "RSA PUBLIC KEY" {
		fmt.Println("Failed to parse PEM block containing the public key")
		return
	}
	publicKeyFromBytes, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Error parsing public key:", err)
		return
	}
	publicKey := publicKeyFromBytes.(*rsa.PublicKey)

	// 4.3 公钥对明文加密，得到密文
	message := "Hello, World!"
	messageBytes := []byte(message)
	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, messageBytes)
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}

	// 4.4 私钥对密文解密，得到明文
	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeyFromBytes, encryptedMessage)
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return
	}
	fmt.Printf("Decrypted message: %s\n", decryptedMessage)

	// 4.5 私钥对消息签名
	hash := sha256.Sum256(messageBytes)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKeyFromBytes, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	}

	// 4.6 公钥对签名进行验证
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		fmt.Println("Error verifying signature:", err)
		return
	}
	fmt.Println("Signature verified.")
}
