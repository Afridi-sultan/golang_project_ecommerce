package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payloader struct {
	Sub         int `json:"sub"`
	FirstName   string `json:"first_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payloader) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}


	//Header
	byteHeaderarr, err := json.Marshal(header)
	if err != nil{
		return "",err
	}

	base64Header := base64UrlEncode(byteHeaderarr)

	//payloader 
	bytePayloaderArr, err := json.Marshal(data)
	if err != nil{
		return "",err
	}

	base64Payloader := base64UrlEncode(bytePayloaderArr)

	//convert in single string base64 header and payloader
	message := base64Header + "."+base64Payloader
	secretByteArr := []byte(secret)
	byteArrMessage := []byte(message)

	//HMAC 
	h := hmac.New(sha256.New, secretByteArr)
	h.Write(byteArrMessage)

	//signature 
	signature := h.Sum(nil)
	base64Signature := base64UrlEncode(signature)

	// JWT 
	jwt := base64Header + "." + base64Payloader + "." + base64Signature
	return jwt, nil

}

//base64 convertion function 
func base64UrlEncode(data []byte)string{
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}


// just testing jwt -> run it from main()
	// jwt, err := util.CreateJwt("my-secret", util.Payloader{
	// 	Sub: "45",
	// 	FirstName: "Afridi Sultan",
	// 	Email: "Afridi@gmail.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(jwt)