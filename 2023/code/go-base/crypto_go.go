package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/1/4
 * @Desc:
 * @Project: 2023
 */

func main() {
	s := "pfinalclub this string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s, "->", bs)
	fmt.Printf("%x\n", bs)
	// base64
	data := "pfinalclub123?@#$()'-=@~"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
