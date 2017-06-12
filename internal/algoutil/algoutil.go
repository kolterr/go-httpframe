//Package algoutil contain some scaffold algo
package algoutil

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	mathrand "math/rand"
	"runtime"
	"serviceFrame/internal/cryptoex"
	"serviceFrame/internal/errutil"
	"serviceFrame/internal/security/security"
	"strconv"
	"time"
)

//MD5String md5 digest in string
func MD5String(plain string) string {
	cipher := MD5([]byte(plain))
	return hex.EncodeToString(cipher)
}

//MD5 md5 digest
func MD5(plain []byte) []byte {
	md5Ctx := md5.New()
	md5Ctx.Write(plain)
	cipher := md5Ctx.Sum(nil)
	return cipher[:]
}

//GenRSAKey gen a rsa key pair, the bit size is 512
func GenRSAKey() (privateKey, publicKey string, err error) {
	//public gen the private key
	privKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return "", "", err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privKey)
	privateKey = base64.StdEncoding.EncodeToString(derStream)

	//gen the public key
	pubKey := &privKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", "", err
	}
	publicKey = base64.StdEncoding.EncodeToString(derPkix)
	return privateKey, publicKey, nil
}

// RSAEncrypt encrypt data by rsa
func RSAEncrypt(plain []byte, pubKey string) ([]byte, error) {
	buf, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return nil, err
	}
	p, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, err
	}
	if pub, ok := p.(*rsa.PublicKey); ok {
		return rsa.EncryptPKCS1v15(rand.Reader, pub, plain) //RSA算法加密
	}
	return nil, errutil.ErrIllegalParameter
}

// RsaDecrypt decrypt data by rsa
func RSADecrypt(cipher []byte, privKey string) ([]byte, error) {
	if cipher == nil {
		return nil, errutil.ErrIllegalParameter
	}
	buf, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return nil, err
	}
	priv, err := x509.ParsePKCS1PrivateKey(buf)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipher) //RSA解密算法
}

// Sign with database appsecret string(base64 encode)
func Sign(plain []byte, privKey string) (string, error) {
	if plain == nil {
		return "", errutil.ErrIllegalParameter
	}
	buf, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return "", err
	}
	priv, err := x509.ParsePKCS1PrivateKey(buf)
	if err != nil {
		return "", err
	}
	return cryptoex.Sign(priv, plain)
}

// Verify with database appkey string(base64 encode)
func Verify(pubKey string, data []byte, sign string) error {
	buf, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return err
	}
	p, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return err
	}
	if pub, ok := p.(*rsa.PublicKey); ok {
		return cryptoex.Verify(pub, data, sign)
	}
	return errutil.ErrIllegalParameter
}

func MaskPhone(phone string) (string, error) {
	if !security.ValidatePhone(phone) {
		return "", errutil.ErrWrongPhoneNumber
	}
	return fmt.Sprintf("%s****%s", phone[:3], phone[7:]), nil
}

// 生成随机字符串
func RandStr(strlen int) string {
	mathrand.Seed(time.Now().Unix())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = mathrand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = mathrand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}
