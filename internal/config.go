package internal

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

type AppConfig struct {
	Username   string
	ApiUser    string
	ApiKey     string
	ClientIp   string
	DomainName string
	Targets    []string
}

func GetConfigFromEnv() (*AppConfig, error) {
	ip, err := GetMyIp()
	if err != nil {
		log.Println(err)
		ip = os.Getenv("NM_CLIENT_IP")
	}
	ac := &AppConfig{
		Username:   os.Getenv("NM_USERNAME"),
		ApiUser:    os.Getenv("NM_API_USER"),
		ApiKey:     os.Getenv("NM_API_KEY"),
		ClientIp:   ip,
		DomainName: os.Getenv("NM_DOMAIN"),
		Targets:    strings.Split(os.Getenv("NM_TARGETS"), ";"),
	}
	err = ac.validateFields()
	if err != nil {
		return ac, err
	}
	return ac, nil
}

func (ac *AppConfig) validateFields() error {
	if ac.Username == "" {
		return errors.New("username not found in environment")
	}
	if ac.ApiUser == "" {
		return errors.New("api user not found in environment")
	}
	if ac.ApiKey == "" {
		return errors.New("api key not found in environment")
	}
	if ac.ClientIp == "" {
		return errors.New("client ip not found in environment")
	}
	if ac.DomainName == "" {
		return errors.New("domain name not found in environment")
	}
	return nil
}

func (ac *AppConfig) Print(hideKey bool) {
	fmt.Println("Username : " + ac.Username)
	fmt.Println("API User : " + ac.ApiUser)
	apiKey := "******SECRET*******"
	if !hideKey {
		apiKey = ac.ApiKey
	}
	fmt.Println("API Key : " + apiKey)
	fmt.Println("ClientIp : " + ac.ClientIp)
	fmt.Println("Domain : " + ac.DomainName)
	fmt.Println("Targets: " + strings.Join(ac.Targets, ";"))
}

func (ac *AppConfig) WriteToDotEnv() {
	file, err := os.Create(".env")
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to write to .env")
	}
	defer file.Close()

	s := fmt.Sprintf(
		"NM_USERNAME=%s\nNM_API_USER=%s\nNM_API_KEY=%s\nNM_CLIENT_IP=%s\nNM_DOMAIN=%s\nNM_TARGETS=%s",
		ac.Username,
		ac.ApiUser,
		ac.ApiKey,
		ac.ClientIp,
		ac.DomainName,
		strings.Join(ac.Targets, ";"),
	)

	_, err = file.WriteString(s)
	if err != nil {
		log.Println(err)
		log.Fatal("Could not write to env")
	}
	fmt.Println("Successfully upated .env")
}

func RunEditForm(ac *AppConfig) {
	targets := strings.Join(ac.Targets, ";")
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Username").
				Value(&ac.Username).
				Inline(true),
			huh.NewInput().
				Title("API User").
				Value(&ac.ApiUser).
				Inline(true),
			huh.NewInput().
				Title("API Key").
				Value(&ac.ApiKey).
				Inline(true),
			huh.NewInput().
				Title("Client IP (Optional)").
				Value(&ac.ClientIp).
				Inline(true),
			huh.NewInput().
				Title("Domain").
				Value(&ac.DomainName).
				Inline(true),
			huh.NewInput().
				Key("targets").
				Title("Target").
				Value(&targets).
				Inline(true),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	ac.Targets = strings.Split(form.GetString("targets"), ":")
	ac.WriteToDotEnv()
}
