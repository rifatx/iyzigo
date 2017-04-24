package main

import (
	"fmt"

	iyzigo ".."
)

func main() {
	iyzigo.SetOptions(iyzigo.Options{
		ApiKey:    "api key",
		SecretKey: "secret key",
		BaseUrl:   "https://sandbox-api.iyzipay.com",
	})

	req := iyzigo.RetrieveBinNumberRequest{
		BinNumber: "540668",
	}

	res := req.RetrieveBinNumber()

	fmt.Println(res)
}
