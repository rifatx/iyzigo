package iyzigo

type (
	IyzipayResource struct {
		Status         string `json:"status"`
		ErrorCode      string `json:"errorCode"`
		ErrorMessage   string `json:"errorMessage"`
		ErrorGroup     string `json:"errorGroup"`
		Locale         string `json:"locale"`
		SystemTime     int64  `json:"systemTime"`
		ConversationId string `json:"conversationId"`
	}
	ApiTest struct {
		IyzipayResource
	}
	Approval struct {
		IyzipayResource
		PaymentTransactionId string `json:"paymentTransactionId"`
	}
	ConvertedPayout struct {
		PaidPrice                     float64 `json:"paidPrice"`
		IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64 `json:"iyziCommissionFee"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
		SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
		IyziConversionRate            float64 `json:"iyziConversionRate"`
		IyziConversionRateAmount      float64 `json:"iyziConversionRateAmount"`
		Currency                      string  `json:"currency"`
	}
	PaymentItem struct {
		ItemId                        string          `json:"itemId"`
		PaymentTransactionId          string          `json:"paymentTransactionId"`
		TransactionStatus             int             `json:"transactionStatus"`
		Price                         float64         `json:"price"`
		PaidPrice                     float64         `json:"paidPrice"`
		MerchantCommissionRate        float64         `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount  float64         `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount      float64         `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64         `json:"iyziCommissionFee"`
		BlockageRate                  float64         `json:"blockageRate"`
		BlockageRateAmountMerchant    float64         `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant float64         `json:"blockageRateAmountSubMerchant"`
		BlockageResolvedDate          string          `json:"blockageResolvedDate"`
		SubMerchantKey                string          `json:"subMerchantKey"`
		SubMerchantPrice              float64         `json:"subMerchantPrice"`
		SubMerchantPayoutRate         float64         `json:"subMerchantPayoutRate"`
		SubMerchantPayoutAmount       float64         `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64         `json:"merchantPayoutAmount"`
		ConvertedPayout               ConvertedPayout `json:"convertedPayout"`
	}
	BasicPaymentResource struct {
		IyzipayResource
		Price                        float64 `json:"price"`
		PaidPrice                    float64 `json:"paidPrice"`
		Installment                  int     `json:"installment"`
		Currency                     string  `json:"currency"`
		PaymentId                    string  `json:"paymentId"`
		MerchantCommissionRate       float64 `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
		IyziCommissionFee            float64 `json:"iyziCommissionFee"`
		CardType                     string  `json:"cardType"`
		CardAssociation              string  `json:"cardAssociation"`
		CardFamily                   string  `json:"cardFamily"`
		CardToken                    string  `json:"cardToken"`
		CardUserKey                  string  `json:"cardUserKey"`
		BinNumber                    string  `json:"binNumber"`
		PaymentTransactionId         string  `json:"paymentTransactionId"`
		AuthCode                     string  `json:"authCode"`
		ConnectorName                string  `json:"connectorName"`
		Phase                        string  `json:"phase"`
	}
	BasicBkm struct {
		BasicPaymentResource
		Token         string `json:"token"`
		CallbackUrl   string `json:"callbackUrl"`
		PaymentStatus string `json:"paymentStatus"`
	}
	BasicBkmInitialize struct {
		IyzipayResource
		HtmlContent string `json:"htmlContent"`
		Token       string `json:"token"`
	}
	BasicPayment struct {
		BasicPaymentResource
	}
	BasicPaymentPostAuth struct {
		BasicPaymentResource
	}
	BasicPaymentPreAuth struct {
		BasicPaymentResource
	}
	BasicThreedsInitialize struct {
		IyzipayResource
		HtmlContent string `json:"threeDSHtmlContent"`
	}
	BasicThreedsInitializePreAuth struct {
		IyzipayResource
		HtmlContent string `json:"threeDSHtmlContent"`
	}
	BasicThreedsPayment struct {
		BasicPaymentResource
	}
	BinNumber struct {
		IyzipayResource
		Bin             string `json:"binNumber"`
		CardType        string `json:"cardType"`
		CardAssociation string `json:"cardAssociation"`
		CardFamily      string `json:"cardFamily"`
		BankName        string `json:"bankName"`
		BankCode        int64  `json:"bankCode"`
	}
	PaymentResource struct {
		IyzipayResource
		Price                        float64       `json:"price"`
		PaidPrice                    float64       `json:"paidPrice"`
		Installment                  int           `json:"installment"`
		Currency                     string        `json:"currency"`
		PaymentId                    string        `json:"paymentId"`
		PaymentStatus                string        `json:"paymentStatus"`
		FraudStatus                  int           `json:"fraudStatus"`
		MerchantCommissionRate       float64       `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount float64       `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount     float64       `json:"iyziCommissionRateAmount"`
		IyziCommissionFee            float64       `json:"iyziCommissionFee"`
		CardType                     string        `json:"cardType"`
		CardAssociation              string        `json:"cardAssociation"`
		CardFamily                   string        `json:"cardFamily"`
		CardToken                    string        `json:"cardToken"`
		CardUserKey                  string        `json:"cardUserKey"`
		BinNumber                    string        `json:"binNumber"`
		BasketId                     string        `json:"basketId"`
		PaymentItems                 []PaymentItem `json:"itemTransactions"`
		ConnectorName                string        `json:"connectorName"`
		AuthCode                     string        `json:"authCode"`
		Phase                        string        `json:"phase"`
	}
	Bkm struct {
		PaymentResource
		Token       string `json:"token"`
		CallbackUrl string `json:"callbackUrl"`
	}
	BkmInitialize struct {
		IyzipayResource
		HtmlContent string `json:"htmlContent"`
		Token       string `json:"token"`
	}
	BankTransfer struct {
		SubMerchantKey             string `json:"subMerchantKey"`
		Iban                       string `json:"iban"`
		ContactName                string `json:"contactName"`
		ContactSurname             string `json:"contactSurname"`
		LegalCompanyTitle          string `json:"legalCompanyTitle"`
		MarketplaceSubMerchantType string `json:"marketplaceSubmerchantType"`
	}
	BouncedBankTransferList struct {
		IyzipayResource
		BankTransfers []BankTransfer `json:"bankTransfers"`
	}
	Cancel struct {
		IyzipayResource
		PaymentId     string  `json:"paymentId"`
		Price         float64 `json:"price"`
		Currency      string  `json:"currency"`
		ConnectorName string  `json:"connectorName"`
	}
	Card struct {
		IyzipayResource
		ExternalId      string `json:"externalId"`
		Email           string `json:"email"`
		CardUserKey     string `json:"cardUserKey"`
		CardToken       string `json:"cardToken"`
		CardAlias       string `json:"cardAlias"`
		BinNumber       string `json:"binNumber"`
		CardType        string `json:"cardType"`
		CardAssociation string `json:"cardAssociation"`
		CardFamily      string `json:"cardFamily"`
		CardBankCode    int64  `json:"cardBankCode"`
		CardBankName    string `json:"cardBankName"`
	}
	CardList struct {
		IyzipayResource
		CardUserKey string `json:"cardUserKey"`
		CardDetails []Card `json:"cardDetails"`
	}
	CheckoutForm struct {
		PaymentResource
		Token       string `json:"token"`
		CallbackUrl string `json:"callbackUrl"`
	}
	CheckoutFormInitializeResource struct {
		IyzipayResource
		Token               string `json:"token"`
		CheckoutFormContent string `json:"checkoutFormContent"`
		TokenExpireTime     int64  `json:"tokenExpireTime"`
		PaymentPageUrl      string `json:"paymentPageUrl"`
	}
	CheckoutFormInitialize struct {
		CheckoutFormInitializeResource
	}
	CheckoutFormInitializePreAuth struct {
		CheckoutFormInitializeResource
	}
	CrossBookingFromSubMerchant struct {
		IyzipayResource
	}
	CrossBookingToSubMerchant struct {
		IyzipayResource
	}
	Disapproval struct {
		IyzipayResource
	}
	InstallmentPrice struct {
		Price             float64 `json:"InstallmentPrice"`
		TotalPrice        float64 `json:"totalPrice"`
		InstallmentNumber int     `json:"installmentNumber"`
	}
	InstallmentDetail struct {
		BinNumber         string             `json:"binNumber"`
		Price             float64            `json:"price"`
		CardType          string             `json:"cardType"`
		CardAssociation   string             `json:"cardAssociation"`
		CardFamilyName    string             `json:"cardFamilyName"`
		Force3Ds          int                `json:"force3Ds"`
		BankCode          int64              `json:"bankCode"`
		BankName          string             `json:"bankName"`
		ForceCvc          int                `json:"forceCvc"`
		InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
	}
	InstallmentInfo struct {
		IyzipayResource
		InstallmentDetails []InstallmentDetail `json:"installmentDetails"`
	}
	Payment struct {
		PaymentResource
	}
	PaymentPostAuth struct {
		PaymentResource
	}
	PaymentPreAuth struct {
		PaymentResource
	}
	PayoutCompletedTransactionList struct {
		IyzipayResource
	}
	PeccoInitialize struct {
		IyzipayResource
		HtmlContent     string `json:"htmlContent"`
		RedirectUrl     string `json:"redirectUrl"`
		Token           string `json:"token"`
		TokenExpireTime int64  `json:"tokenExpireTime"`
	}
	PeccoPayment struct {
		PaymentResource
	}
	Refund struct {
		IyzipayResource
		PaymentId            string  `json:"paymentId"`
		PaymentTransactionId string  `json:"paymentTransactionId"`
		Price                float64 `json:"price"`
		Currency             string  `json:"currency"`
		ConnectorName        string  `json:"connectorName"`
	}
	RefundChargedFromMerchant struct {
		IyzipayResource
		PaymentId            string `json:"paymentId"`
		PaymentTransactionId string `json:"paymentTransactionId"`
		Price                string `json:"price"`
	}
	SubMerchant struct {
		IyzipayResource
		Name                  string `json:"name"`
		Email                 string `json:"email"`
		GsmNumber             string `json:"gsmNumber"`
		Address               string `json:"address"`
		Iban                  string `json:"iban"`
		SwiftCode             string `json:"swiftCode"`
		Currency              string `json:"currency"`
		TaxOffice             string `json:"taxOffice"`
		ContactName           string `json:"contactName"`
		ContactSurname        string `json:"contactSurname"`
		LegalCompanyTitle     string `json:"legalCompanyTitle"`
		SubMerchantExternalId string `json:"subMerchantExternalId"`
		IdentityNumber        string `json:"identityNumber"`
		TaxNumber             string `json:"taxNumber"`
		SubMerchantType       string `json:"subMerchantType"`
		SubMerchantKey        string `json:"subMerchantKey"`
	}
	ThreedsInitialize struct {
		IyzipayResource
		HtmlContent string `json:"threeDSHtmlContent"`
	}
	ThreedsInitializePreAuth struct {
		IyzipayResource
		HtmlContent string `json:"threeDSHtmlContent"`
	}
	ThreedsPayment struct {
		PaymentResource
	}
)
