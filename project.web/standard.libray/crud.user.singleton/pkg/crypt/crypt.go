// Back-End in Go server
// @jeffotoni
// 2019-01-09

package crypt

//
//
//
import (
	"crypto/md5"
	"encoding/base64"
	b64 "encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strconv"

	"time"

	"github.com/jeffotoni/crud.user.singleton/pkg/zerolog"
	"golang.org/x/crypto/bcrypt"
)

var (
	HASH_SALT      = "xxxxxxxxxxxxxxxxxxxxxxxxxxx"
	SHA1_SALT      = "xxxxxxxxxxxxxxxxxxxxxxxxxxx"
	letters        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	UKK_KEY_CIPHER = []byte("KEUEUXUWYEJDUEUE3151SXXY917DICXX")
)

func Random(min, max int) int { rand.Seed(time.Now().Unix()); return rand.Intn(max-min) + min }

func Rand5() string {
	rand.Seed(time.Now().Unix())
	min := 10000
	max := 99999
	return strconv.Itoa(rand.Intn(max-min) + min)
}

func Blowfish(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"crypt.go",
			288,
			"api.crud.user.singleton.com.br",
			"bcrypt.GenerateFromPassword",
			err.Error())
	}
	return string(bytes)
}

func CheckBlowfish(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Md5(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return (fmt.Sprintf("%x", h.Sum(nil)))
}

//
func Base64Enc(textString string) string {
	if len(string(textString)) > 0 {
		text := []byte(textString)
		sEnc := b64.URLEncoding.EncodeToString(text)
		return sEnc
	}
	return ""
}

//
func Base64Dec(textString string) string {
	if len(string(textString)) > 0 {
		sDec, _ := b64.URLEncoding.DecodeString(textString)
		return string(sDec)
	}
	return ""
}

func Encode64String(content string) string {
	if len(content) > 0 {
		return base64.StdEncoding.EncodeToString([]byte(content))
	}
	return ""
}

func Encode64Byte(content []byte) string {
	if len(string(content)) > 0 {
		return base64.StdEncoding.EncodeToString(content)
	}
	return ""
}

func Decode64String(encoded string) string {
	if len(encoded) > 0 {
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			zerolog.Error(
				"1.0.0",
				"crypt.go",
				400,
				"api.crud.user.singleton.com.br",
				"Decode64String base64.StdEncoding.DecodeString",
				err.Error())
			return ""
		}
		return (string(decoded))
	}
	return ""
}
