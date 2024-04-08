package main 

import "github.com/namecheap/go-namecheap-sdk/v2"



func NewClientFromEnv() {
	client := NewClient(&ClientOptions{
		UserName:	os.Getenv("NM_USERNAME"),
		ApiUser:	os.Getenv("NM_API_USER"),


}
