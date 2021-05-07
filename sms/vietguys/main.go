package vietguys

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Send SMS with Vietguys service

// Config ...
type Config struct {
	// Endpoint which will use to call api
	Endpoint string
	// For auth
	User string
	// For auth
	Password string
	// Brand name
	From string
	// Postgresql instance, for saving logs
	PostgreSQL *gorm.DB
	// Table name
	TableName string
}

// Service ...
type Service struct {
	Config Config
	Client *http.Client
}

// NewInstance for using send sms method
func NewInstance(config Config) (*Service, error) {
	if config.Endpoint == "" || config.User == "" || config.Password == "" || config.From == "" {
		return nil, errors.New("please provide all information that needed: endpoint, user, password, from")
	}

	s := Service{
		Config: config,
		Client: &http.Client{},
	}
	return &s, nil
}

// SendOTP ...
func (s Service) SendOTP(phone, content string) (success bool, result Result, rawResult string) {
	// Create payload
	params := url.Values{}
	params.Add("u", s.Config.User)
	params.Add("pwd", s.Config.Password)
	params.Add("from", s.Config.From)
	params.Add("json", "1")
	params.Add("phone", phone)
	params.Add("sms", content)
	payload := strings.NewReader(params.Encode())

	// Create request
	client := s.Client
	req, err := http.NewRequest(http.MethodPost, s.Config.Endpoint, payload)
	if err != nil {
		return
	}

	// Add necessary headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Call
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Make sure close body
	defer res.Body.Close()

	// Ready body
	body, err := ioutil.ReadAll(res.Body)
	rawResult = string(body)
	fmt.Println(rawResult)

	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}

	if result.Error == 0 {
		success = true
	}

	// Save log to db
	if s.Config.PostgreSQL != nil {
		log := Log{
			ID:        newUUID(),
			Carrier:   result.Carrier,
			Type:      "otp",
			Recipient: phone,
			Content:   content,
			CreatedAt: time.Now().UTC(),
			Success:   success,
			Result:    rawResult,
		}
		s.saveLog(log)
	}

	return
}

func newUUID() string {
	value, _ := uuid.NewRandom()
	return value.String()
}
