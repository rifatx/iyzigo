package iyzigo

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func logRequest(req interface{}) {
	fmt.Print("REQLOG: ")
	fmt.Println(req)
}

func logError(err interface{}) {
	fmt.Print("ERROR: ")
	fmt.Println(err)
}

func getRandomString() string {
	return strings.Replace(time.Now().Format(timeFormat), ".", "", 1)
}
func SetOptions(o Options) {
	options = o
}

func (c *client) generateHash(rs string, object interface{}) string {
	s := InitObjectPrinter().
		SetObjectEncapsulation("[", "]").
		SetArrayEncapsulation("[", "]").
		SetKeyValueSeparator("=").
		SetFieldSeparator(",").
		SetArraySeparator(", ").
		GetObjectString(object)
	h := sha1.Sum([]byte(c.ApiKey + rs + c.SecretKey + s))
	return base64.StdEncoding.EncodeToString(h[:])
}

func (c *client) prepareAuthorizationString(rs string, object interface{}) string {
	return iyziws_header_name + c.ApiKey + colon + c.generateHash(rs, object)
}

func (c *client) getHttpHeader(object interface{}) http.Header {
	h := http.Header{}
	rs := getRandomString()
	h.Add("Accept", "application/json")
	h.Add("Content-Type", "application/json; charset=utf-8")
	h.Add("Expect", "100-continue")
	h.Add("Connection", "Keep-Alive")
	h.Add("User-Agent", "")
	h.Add(random_header_name, rs)
	if object != nil {
		h.Add(authorization, c.prepareAuthorizationString(rs, object))
	}
	return h
}

func maskSensitiveInfo(object interface{}) interface{} {
	switch object.(type) {
	case CreatePaymentRequest:
		cpr, _ := object.(CreatePaymentRequest)
		startDigits, endDigits := 1, 4
		cpr.PaymentCard.CardNumber = fmt.Sprintf("%s%s%s", cpr.PaymentCard.CardNumber[0:startDigits], strings.Repeat("*", len(cpr.PaymentCard.CardNumber)-startDigits-endDigits), cpr.PaymentCard.CardNumber[len(cpr.PaymentCard.CardNumber)-endDigits:])
		cpr.PaymentCard.Cvc = strings.Repeat("*", 3)
		cpr.PaymentCard.ExpireMonth = strings.Repeat("*", 2)
		cpr.PaymentCard.ExpireYear = strings.Repeat("*", 2)
		return cpr
	default:
		return object
	}
}

func FormatPriceString(price string) string {
	if !strings.Contains(price, ".") {
		return fmt.Sprintf("%s.0", price)
	}
	return strings.TrimRight(price, "0")
}

func FormatPrice(price float64) string {
	return FormatPriceString(strconv.FormatFloat(float64(price), 'f', -1, 32))
}
