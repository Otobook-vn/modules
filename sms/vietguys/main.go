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

	"github.com/Otobook-vn/modules/utils"
	"github.com/kr/pretty"
	"gorm.io/gorm"
)

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
	// Log table name
	LogTableName string
	// Limit send time per ip
	IPMaxSendPerDay int
	// Limit send time per phone number
	PhoneMaxSendPerDay int
}

// Result ...
type Result struct {
	Carrier   string
	Error     int
	ErrorCode int
	MsgID     string
	Message   string
	Log       string
}

// Service ...
type Service struct {
	Config
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

	if s.PostgreSQL != nil {
		if err := config.PostgreSQL.AutoMigrate(
			&Log{tableName: config.LogTableName},
		); err != nil {
			return nil, errors.New("error when create new database table for logging")
		}

		// TODO: add index for db (field phone, ip, created_at)
		//
	}

	return &s, nil
}

// SendOTP to phone number
// Phone format: 84935123456
func (s Service) SendOTP(phone, content, ip string) error {
	// Just remove char "+" if existed
	strings.Replace(phone, "+", "", 1)

	// Check that phone or ip is not over quota
	canSend := s.checkCanSend(phone, ip)
	if canSend {
		return errors.New("ip or phone has reached over limited quota per day")
	}

	// Create payload
	params := url.Values{}
	params.Add("u", s.User)
	params.Add("pwd", s.Password)
	params.Add("from", s.From)
	params.Add("json", "1")
	params.Add("phone", phone)
	params.Add("sms", content)
	payload := strings.NewReader(params.Encode())

	pretty.Println("params", params)
	pretty.Println("payload", payload)

	// Create request
	client := s.Client
	req, err := http.NewRequest(http.MethodPost, s.Endpoint, payload)
	if err != nil {
		return err
	}

	// Add necessary headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Call
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Make sure close body
	defer res.Body.Close()

	// Ready body
	body, err := ioutil.ReadAll(res.Body)
	rawResult := string(body)
	fmt.Println(rawResult)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return err
	}

	var result Result
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	// Save log to db
	if s.PostgreSQL != nil {
		log := Log{
			ID:          utils.NewUUID(),
			Carrier:     result.Carrier,
			Type:        "otp",
			PhoneNumber: phone,
			IP:          ip,
			Content:     content,
			CreatedAt:   time.Now().UTC(),
			Success:     result.Error == 0,
			Result:      rawResult,
		}
		s.saveLog(log)
	}

	return nil
}
