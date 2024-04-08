package main

import "os"

func NewClientFromEnv() {
	client := NewClient(&ClientOptions{
		UserName: os.Getenv("NM_USERNAME"),
		ApiUser:  os.Getenv("NM_API_USER"),
	})

}
