package api

// import (
// 	"fmt"
// 	"math"
// 	"regexp"
// 	"strconv"
// 	"strings"

// 	"github.com/cavelms/config"
// 	"github.com/cavelms/pkg/mail"
// 	"github.com/cavelms/internal/model"
// 	"github.com/gofiber/fiber/v2"
// )

// type Wallet struct {
// }

// const DOLLER_RATE = 400
// const APPLICATION_FEE = 10
// const TUITION_FEE = 1000

// // var ctx = context.Background()

// func ProcessPayment(m fiber.Map, user *model.User) error {
// 	rdb := config.RedisClient(0)
// 	defer rdb.Close()

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	naration := m["TransactionNarration"].(string)
// 	reg := regexp.MustCompile("[0-9]+")
// 	code := reg.FindAllString(naration, -1)
// 	paymentType := code[0]

// 	amount := m["TransactionAmount"].(string)
// 	a := strings.Replace(amount, ",", "", -1)
// 	a = strings.Split(a, ".")[0]

// 	amt, err := strconv.ParseFloat(a, 64)
// 	if err != nil {
// 		return err
// 	}

// 	// id, err := rdb.Get(ctx, user.UserID).Result()
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// if err := user.FetchByID(id); err != nil {
// 	// 	return err
// 	// }

// 	wallet := user.Wallet + (amt / DOLLER_RATE)
// 	if paymentType == "10" {
// 		wallet = wallet - APPLICATION_FEE
// 	}
// 	if paymentType == "12" {
// 		wallet = wallet - TUITION_FEE
// 	}

// 	if wallet < 0 {
// 		user.Wallet = wallet
// 		err = processIncompletePayment(user, amt)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func processIncompletePayment(user *model.User, paid float64) error {
// 	balance := math.Abs(user.Wallet)
// 	data := fiber.Map{
// 		"fromAddress": "support@adullam.ng",
// 		"toAddress":   user.Email,
// 		"subject":     "Adullam|Payment Confirmation",
// 		"content": map[string]interface{}{
// 			"filename": "repay.html",
// 			"balance":  balance,
// 			"due":      APPLICATION_FEE,
// 			"paid":     paid / DOLLER_RATE,
// 			"email":    user.Email,
// 		},
// 	}

// 	m := new(mail.Mail)
// 	_, err := m.SendMail(data)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
