package server

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Webhooks (ctx *fiber.Ctx) error {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to Load .env")
	}
	
	envMap, envErr := godotenv.Read(".env")
	if envErr != nil {
		log.Fatal("Failed to read .env")
	}

	zoomEvent := new(Event)
	validationEvent := new(Response)
	ctx.BodyParser(zoomEvent)
	ctx.BodyParser(validationEvent)
	webhookTimestamp := ctx.GetReqHeaders()["X-Zm-Request-Timestamp"]
	body := ctx.Request().Body()
	message := "v0:"+webhookTimestamp +":"+ string(body)
	fmt.Println("Event Recieved: ", zoomEvent.Event)
	fmt.Println("Event Recieved: ", zoomEvent.EventTS)
	fmt.Println("Event Recieved: ", zoomEvent.Payload.Operator)
	fmt.Println("Event Recieved: ", zoomEvent.Payload.Object.Duration)
	fmt.Println("Event Recieved: ", zoomEvent.Payload.Object.UUID)
	fmt.Println(InsertEvent(zoomEvent))
	WebhookSecretToken := envMap["SECRETTOKEN"]
	hash := hmac.New(sha256.New, []byte(WebhookSecretToken))
	hash.Write([]byte(message))
	hashForVerify := hex.EncodeToString(hash.Sum(nil))
	signature := fmt.Sprintf("v0=%s", hashForVerify)
//validation that request came from Zoom
	if ctx.GetReqHeaders()["X-Zm-Signature"] == signature {



		//valditation you control webhhook endpoint 
		if zoomEvent.Event == "endpoint.url_validation" {
			fmt.Println("I'm ready to validate...")
		payloadData := zoomEvent.Payload.PlainToken
		salt := envMap["SECRETTOKEN"]
		mac := hmac.New(sha256.New, []byte(salt))

		mac.Write([]byte(payloadData))
		encryptedToken := hex.EncodeToString(mac.Sum(nil))

		response := Response {
			Message: ValidationEvent{ 
				PlainToken:  payloadData,
				Encrypted: encryptedToken,
			},
			Status: 200,
		}
		
		newResponse, err := json.Marshal(response.Message)
		if err != nil {
			log.Fatal(err)
		}

		ctx.Response().Header.Add("content-type", "application/json")
		ctx.Response().AppendBody(newResponse)
		ctx.Response().SetBody(ctx.Response().Body())



	return ctx.SendString(string(newResponse))

	}else {

		type webhhookValidation struct {
			Message string `json:"message"`
			Status int `json:"status"`
		}
			verifiedStruct := webhhookValidation{
				Message : "Authorized Request to Webhook Sample Go App",
				Status: 200,
			}	
			verified, err := json.Marshal(verifiedStruct.Message)
			if err != nil{
				log.Fatal(err)
			}
			ctx.Response().Header.Add("content-type", "application/json")
			ctx.Response().AppendBody(verified)
			ctx.Response().SetBody(ctx.Response().Body())
			fmt.Println(string(ctx.Response().Body()))

		return ctx.SendString(string(verified))
	}
		}
	
	
	return ctx.JSON(zoomEvent)
}



func TokenGeneration(ctx *fiber.Ctx) error {
	type token struct {
		AccessToken string `json:"Access_token"`
		TokenType string `json:"Token_type"`
		ExpiresIn int `json:"Expires_in"`
		Scope string `json:"scope"`
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountID := os.Getenv("ACCOUNTID")
	cliendID := os.Getenv("CLIENTID")
	secret := os.Getenv("CLIENTSECRET")
	authorization := b64.StdEncoding.EncodeToString([]byte(cliendID + ":" + secret))
	newToken := new(token)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://zoom.us/oauth/token?grant_type=account_credentials&account_id=%s", accountID), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", authorization))
	reqClient := &http.Client{}
	resp, err := reqClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &newToken)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	


	return ctx.JSON(newToken) 
}

func CreateMeeting(ctx *fiber.Ctx) error{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	envMap, envErr := godotenv.Read(".env")
	if envErr != nil {
		log.Fatal(envErr)
	}
	token := envMap["TOKEN"]

	var aNewMeeting NewMeeting
	aNewMeeting.Topic = "Meeting Number " + gofakeit.CreditCard().Number
	aNewMeeting.Agenda = gofakeit.HipsterParagraph(1, 3, 9, ",")
	aNewMeeting.Duration = 10
	aNewMeeting.StartTime = time.Time.String(time.Now())
	aNewMeeting.Timezone = gofakeit.TimeZone()
	aNewMeeting.Type = 2

	requestBody, err := json.Marshal(aNewMeeting)
	if err != nil {
		log.Printf("verbose error info: %#v", err)
	}


	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.zoom.us/v2/users/me/meetings"), bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	reqClient := &http.Client{}
	resp, err := reqClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &aNewMeeting)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return ctx.JSON(aNewMeeting)
	
}

//Random Tests


func TestingDotEnv (ctx *fiber.Ctx) error {
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	envMap, envErr := godotenv.Read(".env")
	if envErr != nil {
		log.Fatal(envErr)
	}

	Token := envMap["TOKEN"]
	Token = "some new value"
	fmt.Println(Token)

	newValue := godotenv.Write(envMap , ".env")
	fmt.Println(newValue)
		

	// godotenv.Write(envs, ".env")

	return ctx.JSON(Token +  " line 99")

}



