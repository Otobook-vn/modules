package vietguys

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/Otobook-vn/modules/constant"
	"github.com/Otobook-vn/modules/mongodb"
	"github.com/Otobook-vn/modules/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBConfig ...
type MongoDBConfig struct {
	Host, User, Password, DBName, Mechanism, Source string
}

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
	// MongoDB config, for save documents
	MongoDB MongoDBConfig
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
	DB     *mongo.Collection
}

var s *Service
var bgCtx = context.Background()

// NewInstance for using send sms method
func NewInstance(config Config) error {
	if config.Endpoint == "" || config.User == "" || config.Password == "" || config.From == "" || config.MongoDB.Host == "" {
		return errors.New("please provide all information that needed: endpoint, user, password, from, mongodb")
	}

	// Connect MongoDB
	err := mongodb.Connect(
		config.MongoDB.Host,
		config.MongoDB.User,
		config.MongoDB.Password,
		config.MongoDB.DBName,
		config.MongoDB.Mechanism,
		config.MongoDB.Source,
	)
	if err != nil {
		fmt.Println("Cannot init module VIETGUYS", err)
		return err
	}

	s = &Service{
		Config: config,
		Client: &http.Client{},
		DB:     mongodb.GetInstance().Collection("log_sms_vietguys"),
	}

	return nil
}

// SendOTP to phone number
// Phone format: 84935123456
func (s Service) SendOTP(phone, ip, content, otpCode string) error {
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

	go func() {
		// Save log to db
		log := Log{
			ID:          utils.GenerateMongoID(),
			Carrier:     result.Carrier,
			Type:        constant.SMSTypeOTP,
			Code:        otpCode,
			IsCodeValid: true,
			PhoneNumber: phone,
			IP:          ip,
			Content:     content,
			CreatedAt:   utils.TimeNow(),
			Success:     result.Error == 0,
			Result:      rawResult,
		}
		s.saveLog(log)
	}()

	return nil
}

// VerifyOTP verify otp code is right or not
func (s Service) VerifyOTP(phone, otpCode string) bool {
	// Just remove char "+" if existed
	strings.Replace(phone, "+", "", 1)

	return s.checkOTP(phone, otpCode)
}
