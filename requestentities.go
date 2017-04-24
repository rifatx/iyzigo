package iyzigo

type (
	ApiTestRequest        struct{}
	CreateApprovalRequest struct {
		Locale               string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId       string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		PaymentTransactionId string `opt:"name:paymentTransactionId,omitempty" json:"paymentTransactionId,omitempty"`
	}
	RetrieveBkmRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Token          string `opt:"name:token,omitempty" json:"token,omitempty"`
	}
	BkmInstallmentPrice struct {
		InstallmentNumber int    `opt:"name:installmentNumber,omitempty" json:"installmentNumber,omitempty"`
		TotalPrice        string `opt:"name:totalPrice,omitempty" json:"totalPrice,omitempty"`
	}
	BkmInstallment struct {
		BankId            int64                  `opt:"name:bankId,omitempty" json:"bankId,omitempty"`
		InstallmentPrices []*BkmInstallmentPrice `opt:"name:installmentPrices,omitempty" json:"installmentPrices,omitempty"`
	}
	CreateBasicBkmInitializeRequest struct {
		Locale             string            `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId     string            `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		ConnectorName      string            `opt:"name:connectorName,omitempty" json:"connectorName,omitempty"`
		Price              string            `opt:"name:price,omitempty" json:"price,omitempty"`
		CallbackUrl        string            `opt:"name:callbackUrl,omitempty" json:"callbackUrl,omitempty"`
		BuyerEmail         string            `opt:"name:buyerEmail,omitempty" json:"buyerEmail,omitempty"`
		BuyerId            string            `opt:"name:buyerId,omitempty" json:"buyerId,omitempty"`
		BuyerIp            string            `opt:"name:buyerIp,omitempty" json:"buyerIp,omitempty"`
		PosOrderId         string            `opt:"name:posOrderId,omitempty" json:"posOrderId,omitempty"`
		InstallmentDetails []*BkmInstallment `opt:"name:installmentDetails,omitempty" json:"installmentDetails,omitempty"`
	}
	PaymentCard struct {
		CardHolderName string `opt:"name:cardHolderName,omitempty" json:"cardHolderName,omitempty"`
		CardNumber     string `opt:"name:cardNumber,omitempty" json:"cardNumber,omitempty"`
		ExpireYear     string `opt:"name:expireYear,omitempty" json:"expireYear,omitempty"`
		ExpireMonth    string `opt:"name:expireMonth,omitempty" json:"expireMonth,omitempty"`
		Cvc            string `opt:"name:cvc,omitempty" json:"cvc,omitempty"`
		RegisterCard   int    `opt:"name:registerCard,omitempty" json:"registerCard,omitempty"`
		CardAlias      string `opt:"name:cardAlias,omitempty" json:"cardAlias,omitempty"`
		CardToken      string `opt:"name:cardToken,omitempty" json:"cardToken,omitempty"`
		CardUserKey    string `opt:"name:cardUserKey,omitempty" json:"cardUserKey,omitempty"`
	}
	CreateBasicPaymentRequest struct {
		Locale         string       `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string       `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Price          string       `opt:"name:price,omitempty" json:"price,omitempty"`
		PaidPrice      string       `opt:"name:paidPrice,omitempty" json:"paidPrice,omitempty"`
		Installment    int          `opt:"name:installment,omitempty" json:"installment,omitempty"`
		BuyerEmail     string       `opt:"name:buyerEmail,omitempty" json:"buyerEmail,omitempty"`
		BuyerId        string       `opt:"name:buyerId,omitempty" json:"buyerId,omitempty"`
		BuyerIp        string       `opt:"name:buyerIp,omitempty" json:"buyerIp,omitempty"`
		PosOrderId     string       `opt:"name:posOrderId,omitempty" json:"posOrderId,omitempty"`
		PaymentCard    *PaymentCard `opt:"name:paymentCard,omitempty" json:"paymentCard,omitempty"`
		Currency       string       `opt:"name:currency,omitempty" json:"currency,omitempty"`
		ConnectorName  string       `opt:"name:connectorName,omitempty" json:"connectorName,omitempty"`
		CallbackUrl    string       `opt:"name:callbackUrl,omitempty" json:"callbackUrl,omitempty"`
	}
	CreatePaymentPostAuthRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		PaymentId      string `opt:"name:paymentId,omitempty" json:"paymentId,omitempty"`
		PaidPrice      string `opt:"name:paidPrice,omitempty" json:"paidPrice,omitempty"`
		Ip             string `opt:"name:ip,omitempty" json:"ip,omitempty"`
		Currency       string `opt:"name:currency,omitempty" json:"currency,omitempty"`
	}
	CreateThreedsPaymentRequest struct {
		Locale           string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId   string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		PaymentId        string `opt:"name:paymentId,omitempty" json:"paymentId,omitempty"`
		ConversationData string `opt:"name:conversationData,omitempty" json:"conversationData,omitempty"`
	}
	RetrieveBinNumberRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		BinNumber      string `opt:"name:binNumber,omitempty" json:"binNumber,omitempty"`
	}
	Address struct {
		Description string `opt:"name:address,omitempty" json:"address,omitempty"`
		ZipCode     string `opt:"name:zipCode,omitempty" json:"zipCode,omitempty"`
		ContactName string `opt:"name:contactName,omitempty" json:"contactName,omitempty"`
		City        string `opt:"name:city,omitempty" json:"city,omitempty"`
		Country     string `opt:"name:country,omitempty" json:"country,omitempty"`
	}
	BasketItem struct {
		Id               string `opt:"name:id,omitempty" json:"id,omitempty"`
		Price            string `opt:"name:price,omitempty" json:"price,omitempty"`
		Name             string `opt:"name:name,omitempty" json:"name,omitempty"`
		Category1        string `opt:"name:category1,omitempty" json:"category1,omitempty"`
		Category2        string `opt:"name:category2,omitempty" json:"category2,omitempty"`
		ItemType         string `opt:"name:itemType,omitempty" json:"itemType,omitempty"`
		SubMerchantKey   string `opt:"name:subMerchantKey,omitempty" json:"subMerchantKey,omitempty"`
		SubMerchantPrice string `opt:"name:subMerchantPrice,omitempty" json:"subMerchantPrice,omitempty"`
	}
	Buyer struct {
		Id                  string `opt:"name:id,omitempty" json:"id,omitempty"`
		Name                string `opt:"name:name,omitempty" json:"name,omitempty"`
		Surname             string `opt:"name:surname,omitempty" json:"surname,omitempty"`
		IdentityNumber      string `opt:"name:identityNumber,omitempty" json:"identityNumber,omitempty"`
		Email               string `opt:"name:email,omitempty" json:"email,omitempty"`
		GsmNumber           string `opt:"name:gsmNumber,omitempty" json:"gsmNumber,omitempty"`
		RegistrationDate    string `opt:"name:registrationDate,omitempty" json:"registrationDate,omitempty"`
		LastLoginDate       string `opt:"name:lastLoginDate,omitempty" json:"lastLoginDate,omitempty"`
		RegistrationAddress string `opt:"name:registrationAddress,omitempty" json:"registrationAddress,omitempty"`
		City                string `opt:"name:city,omitempty" json:"city,omitempty"`
		Country             string `opt:"name:country,omitempty" json:"country,omitempty"`
		ZipCode             string `opt:"name:zipCode,omitempty" json:"zipCode,omitempty"`
		Ip                  string `opt:"name:ip,omitempty" json:"ip,omitempty"`
	}
	CreateBkmInitializeRequest struct {
		Locale          string        `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId  string        `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Price           float64       `opt:"name:price,omitempty" json:"price,omitempty"`
		BasketId        string        `opt:"name:basketId,omitempty" json:"basketId,omitempty"`
		PaymentGroup    string        `opt:"name:paymentGroup,omitempty" json:"paymentGroup,omitempty"`
		PaymentSource   string        `opt:"name:paymentSource,omitempty" json:"paymentSource,omitempty"`
		Buyer           *Buyer        `opt:"name:buyer,omitempty" json:"buyer,omitempty"`
		ShippingAddress *Address      `opt:"name:shippingAddress,omitempty" json:"shippingAddress,omitempty"`
		BillingAddress  *Address      `opt:"name:billingAddress,omitempty" json:"billingAddress,omitempty"`
		BasketItems     []*BasketItem `opt:"name:basketItems,omitempty" json:"basketItems,omitempty"`
		CallbackUrl     string        `opt:"name:callbackUrl,omitempty" json:"callbackUrl,omitempty"`
	}
	CreateCancelRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		PaymentId      string `opt:"name:paymentId,omitempty" json:"paymentId,omitempty"`
		Ip             string `opt:"name:ip,omitempty" json:"ip,omitempty"`
	}
	CardInformation struct {
		CardAlias      string `opt:"name:cardAlias,omitempty" json:"cardAlias,omitempty"`
		CardNumber     string `opt:"name:cardNumber,omitempty" json:"cardNumber,omitempty"`
		ExpireYear     string `opt:"name:expireYear,omitempty" json:"expireYear,omitempty"`
		ExpireMonth    string `opt:"name:expireMonth,omitempty" json:"expireMonth,omitempty"`
		CardHolderName string `opt:"name:cardHolderName,omitempty" json:"cardHolderName,omitempty"`
	}
	CreateCardRequest struct {
		Locale         string           `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string           `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		ExternalId     string           `opt:"name:externalId,omitempty" json:"externalId,omitempty"`
		Email          string           `opt:"name:email,omitempty" json:"email,omitempty"`
		CardUserKey    string           `opt:"name:cardUserKey,omitempty" json:"cardUserKey,omitempty"`
		Card           *CardInformation `opt:"name:card,omitempty" json:"card,omitempty"`
	}
	DeleteCardRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		CardUserKey    string `opt:"name:cardUserKey,omitempty" json:"cardUserKey,omitempty"`
		CardToken      string `opt:"name:cardToken,omitempty" json:"cardToken,omitempty"`
	}
	RetrieveCardListRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		CardUserKey    string `opt:"name:cardUserKey,omitempty" json:"cardUserKey,omitempty"`
	}
	RetrieveCheckoutFormRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Token          string `opt:"name:token,omitempty" json:"token,omitempty"`
	}
	CreateCheckoutFormInitializeRequest struct {
		Locale              string        `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId      string        `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Price               float64       `opt:"name:price,omitempty" json:"price,omitempty"`
		PaidPrice           float64       `opt:"name:paidPrice,omitempty" json:"paidPrice,omitempty"`
		BasketId            string        `opt:"name:basketId,omitempty" json:"basketId,omitempty"`
		PaymentGroup        string        `opt:"name:paymentGroup,omitempty" json:"paymentGroup,omitempty"`
		PaymentSource       string        `opt:"name:paymentSource,omitempty" json:"paymentSource,omitempty"`
		Currency            string        `opt:"name:currency,omitempty" json:"currency,omitempty"`
		Buyer               *Buyer        `opt:"name:buyer,omitempty" json:"buyer,omitempty"`
		ShippingAddress     *Address      `opt:"name:shippingAddress,omitempty" json:"shippingAddress,omitempty"`
		BillingAddress      *Address      `opt:"name:billingAddress,omitempty" json:"billingAddress,omitempty"`
		BasketItems         []*BasketItem `opt:"name:basketItems,omitempty" json:"basketItems,omitempty"`
		CallbackUrl         string        `opt:"name:callbackUrl,omitempty" json:"callbackUrl,omitempty"`
		ForceThreeDS        int           `opt:"name:forceThreeDs,omitempty" json:"forceThreeDs,omitempty"`
		CardUserKey         string        `opt:"name:cardUserKey,omitempty" json:"cardUserKey,omitempty"`
		PosOrderId          string        `opt:"name:posOrderId,omitempty" json:"posOrderId,omitempty"`
		EnabledInstallments []int         `opt:"name:enabledInstallments,omitempty" json:"enabledInstallments,omitempty"`
	}
	CreateCrossBookingRequest struct {
		Locale         string  `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string  `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		SubMerchantKey string  `opt:"name:subMerchantKey,omitempty" json:"subMerchantKey,omitempty"`
		Price          float64 `opt:"name:price,omitempty" json:"price,omitempty"`
		Reason         string  `opt:"name:reason,omitempty" json:"reason,omitempty"`
		Currency       string  `opt:"name:currency,omitempty" json:"currency,omitempty"`
	}
	RetrieveInstallmentInfoRequest struct {
		Locale         string  `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string  `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		BinNumber      string  `opt:"name:binNumber,omitempty" json:"binNumber,omitempty"`
		Price          float64 `opt:"name:price,omitempty" json:"price,omitempty"`
	}
	CreatePaymentRequest struct {
		Locale          string        `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId  string        `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Price           string        `opt:"name:price,omitempty" json:"price,omitempty"`
		PaidPrice       string        `opt:"name:paidPrice,omitempty" json:"paidPrice,omitempty"`
		Installment     int           `opt:"name:installment,omitempty" json:"installment,omitempty"`
		PaymentChannel  string        `opt:"name:paymentChannel,omitempty" json:"paymentChannel,omitempty"`
		BasketId        string        `opt:"name:basketId,omitempty" json:"basketId,omitempty"`
		PaymentGroup    string        `opt:"name:paymentGroup,omitempty" json:"paymentGroup,omitempty"`
		PaymentCard     *PaymentCard  `opt:"name:paymentCard,omitempty" json:"paymentCard,omitempty"`
		Buyer           *Buyer        `opt:"name:buyer,omitempty" json:"buyer,omitempty"`
		ShippingAddress *Address      `opt:"name:shippingAddress,omitempty" json:"shippingAddress,omitempty"`
		BillingAddress  *Address      `opt:"name:billingAddress,omitempty" json:"billingAddress,omitempty"`
		BasketItems     []*BasketItem `opt:"name:basketItems,omitempty" json:"basketItems,omitempty"`
		PaymentSource   string        `opt:"name:paymentSource,omitempty" json:"paymentSource,omitempty"`
		Currency        string        `opt:"name:currency,omitempty" json:"currency,omitempty"`
		PosOrderId      string        `opt:"name:posOrderId,omitempty" json:"posOrderId,omitempty"`
		ConnectorName   string        `opt:"name:connectorName,omitempty" json:"connectorName,omitempty"`
		CallbackUrl     string        `opt:"name:callbackUrl,omitempty" json:"callbackUrl,omitempty"`
	}
	RetrievePaymentRequest struct {
		Locale                string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId        string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		PaymentId             string `opt:"name:paymentId,omitempty" json:"paymentId,omitempty"`
		PaymentConversationId string `opt:"name:paymentConversationId,omitempty" json:"paymentConversationId,omitempty"`
	}
	RetrieveTransactionsRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Date           string `opt:"name:date,omitempty" json:"date,omitempty"`
	}
	CreatePeccoInitializeRequest struct {
		Locale          string        `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId  string        `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Price           float64       `opt:"name:price,omitempty" json:"price,omitempty"`
		PaidPrice       float64       `opt:"name:paidPrice,omitempty" json:"paidPrice,omitempty"`
		Currency        string        `opt:"name:currency,omitempty" json:"currency,omitempty"`
		BasketId        string        `opt:"name:basketId,omitempty" json:"basketId,omitempty"`
		PaymentGroup    string        `opt:"name:paymentGroup,omitempty" json:"paymentGroup,omitempty"`
		PaymentSource   string        `opt:"name:paymentSource,omitempty" json:"paymentSource,omitempty"`
		Buyer           *Buyer        `opt:"name:buyer,omitempty" json:"buyer,omitempty"`
		ShippingAddress *Address      `opt:"name:shippingAddress,omitempty" json:"shippingAddress,omitempty"`
		BillingAddress  *Address      `opt:"name:billingAddress,omitempty" json:"billingAddress,omitempty"`
		BasketItems     []*BasketItem `opt:"name:basketItems,omitempty" json:"basketItems,omitempty"`
		CallbackUrl     string        `opt:"name:callbackUrl,omitempty" json:"callbackUrl,omitempty"`
	}
	CreatePeccoPaymentRequest struct {
		Locale         string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Token          string `opt:"name:token,omitempty" json:"token,omitempty"`
	}
	CreateRefundRequest struct {
		Locale               string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId       string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		PaymentTransactionId string `opt:"name:paymentTransactionId,omitempty" json:"paymentTransactionId,omitempty"`
		Price                string `opt:"name:price,omitempty" json:"price,omitempty"`
		Ip                   string `opt:"name:ip,omitempty" json:"ip,omitempty"`
		Currency             string `opt:"name:currency,omitempty" json:"currency,omitempty"`
	}
	CreateSubMerchantRequest struct {
		Locale                string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId        string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Name                  string `opt:"name:name,omitempty" json:"name,omitempty"`
		Email                 string `opt:"name:email,omitempty" json:"email,omitempty"`
		GsmNumber             string `opt:"name:gsmNumber,omitempty" json:"gsmNumber,omitempty"`
		Address               string `opt:"name:address,omitempty" json:"address,omitempty"`
		Iban                  string `opt:"name:iban,omitempty" json:"iban,omitempty"`
		TaxOffice             string `opt:"name:taxOffice,omitempty" json:"taxOffice,omitempty"`
		ContactName           string `opt:"name:contactName,omitempty" json:"contactName,omitempty"`
		ContactSurname        string `opt:"name:contactSurname,omitempty" json:"contactSurname,omitempty"`
		LegalCompanyTitle     string `opt:"name:legalCompanyTitle,omitempty" json:"legalCompanyTitle,omitempty"`
		SwiftCode             string `opt:"name:swiftCode,omitempty" json:"swiftCode,omitempty"`
		Currency              string `opt:"name:currency,omitempty" json:"currency,omitempty"`
		SubMerchantExternalId string `opt:"name:subMerchantExternalId,omitempty" json:"subMerchantExternalId,omitempty"`
		IdentityNumber        string `opt:"name:identityNumber,omitempty" json:"identityNumber,omitempty"`
		TaxNumber             string `opt:"name:taxNumber,omitempty" json:"taxNumber,omitempty"`
		SubMerchantType       string `opt:"name:subMerchantType,omitempty" json:"subMerchantType,omitempty"`
	}
	UpdateSubMerchantRequest struct {
		Locale            string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId    string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		Name              string `opt:"name:name,omitempty" json:"name,omitempty"`
		Email             string `opt:"name:email,omitempty" json:"email,omitempty"`
		GsmNumber         string `opt:"name:gsmNumber,omitempty" json:"gsmNumber,omitempty"`
		Address           string `opt:"name:address,omitempty" json:"address,omitempty"`
		Iban              string `opt:"name:iban,omitempty" json:"iban,omitempty"`
		TaxOffice         string `opt:"name:taxOffice,omitempty" json:"taxOffice,omitempty"`
		ContactName       string `opt:"name:contactName,omitempty" json:"contactName,omitempty"`
		ContactSurname    string `opt:"name:contactSurname,omitempty" json:"contactSurname,omitempty"`
		LegalCompanyTitle string `opt:"name:legalCompanyTitle,omitempty" json:"legalCompanyTitle,omitempty"`
		SubMerchantKey    string `opt:"name:subMerchantKey,omitempty" json:"subMerchantKey,omitempty"`
		IdentityNumber    string `opt:"name:identityNumber,omitempty" json:"identityNumber,omitempty"`
		TaxNumber         string `opt:"name:taxNumber,omitempty" json:"taxNumber,omitempty"`
		Currency          string `opt:"name:currency,omitempty" json:"currency,omitempty"`
		SwiftCode         string `opt:"name:swiftCode,omitempty" json:"swiftCode,omitempty"`
	}
	RetrieveSubMerchantRequest struct {
		Locale                string `opt:"name:locale,omitempty" json:"locale,omitempty"`
		ConversationId        string `opt:"name:conversationId,omitempty" json:"conversationId,omitempty"`
		SubMerchantExternalId string `opt:"name:subMerchantExternalId,omitempty" json:"subMerchantExternalId,omitempty"`
	}
)
