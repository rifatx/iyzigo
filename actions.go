package iyzigo

import (
	"encoding/json"
)

func (req *ApiTestRequest) Retrieve() (res *ApiTest) {
	output, err := doGet("/payment/test")
	if err != nil {
		logError(err)
		return nil
	}
	res = &ApiTest{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveBinNumberRequest) RetrieveBinNumber() (res *BinNumber) {
	output, err := doPost("/payment/bin/check", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BinNumber{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveInstallmentInfoRequest) RetrieveInstallmentInfo() (res *InstallmentInfo) {
	output, err := doPost("/payment/iyzipos/installment", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &InstallmentInfo{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateBasicPaymentRequest) CreateBasicPayment() (res *BasicPayment) {
	output, err := doPost("/payment/auth/basic", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BasicPayment{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreatePaymentPostAuthRequest) CreateBasicPaymentPostAuth() (res *BasicPaymentPostAuth) {
	output, err := doPost("/payment/postauth/basic", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BasicPaymentPostAuth{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateBasicPaymentRequest) CreateBasicPaymentPreAuth() (res *BasicPaymentPreAuth) {
	output, err := doPost("/payment/preauth/basic", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BasicPaymentPreAuth{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateBasicPaymentRequest) CreateBasicThreedsInitialize() (res *BasicThreedsInitialize) {
	output, err := doPost("/payment/3dsecure/initialize/basic", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BasicThreedsInitialize{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateBasicPaymentRequest) CreateBasicThreedsInitializePreAuth() (res *BasicThreedsInitializePreAuth) {
	output, err := doPost("/payment/3dsecure/initialize/preauth/basic", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BasicThreedsInitializePreAuth{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateThreedsPaymentRequest) CreateBasicThreedsPayment() (res *BasicThreedsPayment) {
	output, err := doPost("/payment/3dsecure/auth/basic", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BasicThreedsPayment{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreatePaymentRequest) CreatePayment() (res *Payment) {
	output, err := doPost("/payment/auth", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Payment{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrievePaymentRequest) RetrievePayment() (res *Payment) {
	output, err := doPost("/payment/detail", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Payment{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreatePaymentPostAuthRequest) CreatePaymentPostAuth() (res *PaymentPostAuth) {
	output, err := doPost("/payment/postauth", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &PaymentPostAuth{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreatePaymentRequest) CreatePaymentPreAuth() (res *PaymentPreAuth) {
	output, err := doPost("/payment/preauth", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &PaymentPreAuth{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreatePaymentRequest) CreateThreedsInitialize() (res *ThreedsInitialize) {
	output, err := doPost("/payment/3dsecure/initialize", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &ThreedsInitialize{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreatePaymentRequest) CreateThreedsInitializePreAuth() (res *ThreedsInitializePreAuth) {
	output, err := doPost("/payment/3dsecure/initialize/preauth", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &ThreedsInitializePreAuth{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateThreedsPaymentRequest) CreateThreedsPayment() (res *ThreedsPayment) {
	output, err := doPost("/payment/3dsecure/auth", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &ThreedsPayment{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrievePaymentRequest) RetrieveThreedsPayment() (res *ThreedsPayment) {
	output, err := doPost("/payment/detail", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &ThreedsPayment{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveCheckoutFormRequest) CreateCheckoutForm() (res *CheckoutForm) {
	output, err := doPost("/payment/iyzipos/checkoutform/auth/ecom/detail", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &CheckoutForm{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateCheckoutFormInitializeRequest) CreateCheckoutFormInitialize() (res *CheckoutFormInitialize) {
	output, err := doPost("/payment/iyzipos/checkoutform/initialize/auth/ecom", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &CheckoutFormInitialize{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateCheckoutFormInitializeRequest) CreateCheckoutFormInitializePreAuth() (res *CheckoutFormInitializePreAuth) {
	output, err := doPost("/payment/iyzipos/checkoutform/initialize/preauth/ecom", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &CheckoutFormInitializePreAuth{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveBkmRequest) RetrieveBasicBkm() (res *BasicBkm) {
	output, err := doPost("/payment/bkm/auth/detail/basic", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BasicBkm{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateBasicBkmInitializeRequest) CreateBasicBkmInitialize() (res *BasicBkmInitialize) {
	output, err := doPost("/payment/bkm/initialize/basic", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BasicBkmInitialize{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveBkmRequest) CreateBkm() (res *Bkm) {
	output, err := doPost("/payment/bkm/auth/detail", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Bkm{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateBkmInitializeRequest) CreateBkmInitialize() (res *BkmInitialize) {
	output, err := doPost("/payment/bkm/initialize", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BkmInitialize{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreatePeccoInitializeRequest) CreatePeccoInitialize() (res *PeccoInitialize) {
	output, err := doPost("/payment/pecco/initialize", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &PeccoInitialize{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreatePeccoPaymentRequest) CreatePeccoPayment() (res *PeccoPayment) {
	output, err := doPost("/payment/pecco/auth", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &PeccoPayment{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrievePaymentRequest) RetrievePaymentPreAuth() (res *PaymentPreAuth) {
	output, err := doPost("/payment/detail", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &PaymentPreAuth{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateApprovalRequest) CreateApproval() (res *Approval) {
	output, err := doPost("/payment/iyzipos/item/approve", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Approval{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateApprovalRequest) CreateDisapproval() (res *Disapproval) {
	output, err := doPost("/payment/iyzipos/item/disapprove", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Disapproval{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateCancelRequest) RetrieveCancel() (res *Cancel) {
	output, err := doPost("/payment/cancel", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Cancel{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateRefundRequest) CreateRefund() (res *Refund) {
	output, err := doPost("/payment/refund", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Refund{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateRefundRequest) CreateRefundChargedFromMerchant() (res *RefundChargedFromMerchant) {
	output, err := doPost("/payment/iyzipos/refund/merchant/charge", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &RefundChargedFromMerchant{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateCardRequest) CreateCard() (res *Card) {
	output, err := doPost("/cardstorage/card", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Card{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *DeleteCardRequest) DeleteCard() (res *Card) {
	output, err := doDelete("/cardstorage/card", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &Card{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveCardListRequest) CreateCardList() (res *CardList) {
	output, err := doPost("/cardstorage/cards", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &CardList{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateCrossBookingRequest) CreateCrossBookingFromSubMerchant() (res *CrossBookingFromSubMerchant) {
	output, err := doPost("/crossbooking/receive", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &CrossBookingFromSubMerchant{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateCrossBookingRequest) CreateCrossBookingToSubMerchant() (res *CrossBookingToSubMerchant) {
	output, err := doPost("/crossbooking/send", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &CrossBookingToSubMerchant{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *CreateSubMerchantRequest) CreateSubMerchant() (res *SubMerchant) {
	output, err := doPost("/onboarding/submerchant", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &SubMerchant{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *UpdateSubMerchantRequest) UpdateSubMerchant() (res *SubMerchant) {
	output, err := doPut("/onboarding/submerchant", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &SubMerchant{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveSubMerchantRequest) RetrieveSubMerchant() (res *SubMerchant) {
	output, err := doPost("/onboarding/submerchant/detail", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &SubMerchant{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveTransactionsRequest) RetrievePayoutCompletedTransactionList() (res *PayoutCompletedTransactionList) {
	output, err := doPost("/reporting/settlement/payoutcompleted", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &PayoutCompletedTransactionList{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}

func (req *RetrieveTransactionsRequest) RetrieveBouncedBankTransferList() (res *BouncedBankTransferList) {
	output, err := doPost("/reporting/settlement/bounced", req)
	if err != nil {
		logError(err)
		return nil
	}
	res = &BouncedBankTransferList{}
	err = json.Unmarshal(output, res)
	if err != nil {
		logError(err)
		return nil
	}
	return
}
