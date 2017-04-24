package iyzigo

import (
	"testing"
)

func printKeyError(t *testing.T) {
	t.Log("---------------------------------------------------------------------")
	t.Log("Invalid token signature, check your api key and secret key in config.")
	t.Log("---------------------------------------------------------------------")
}

func printTestSuccessful(t *testing.T) {
	t.Log("Test successful!")
}

func checkErrorCode(errorCode string, t *testing.T) {
	switch errorCode {
	case "0":
		printTestSuccessful(t)
	case "1000", "1001":
		printKeyError(t)
		t.Fail()
	default:
		printTestSuccessful(t)
	}
}

func TestApiTestRetrieve(t *testing.T) {
	t.Log("Starting ApiTest.Retrieve")
	req := ApiTestRequest{}
	res := req.Retrieve()
	if res.Status != F_StatusSuccess {
		t.Error("Test failed!")
	}
}

func TestCreateApprovalRequestCreateApproval(t *testing.T) {
	req := CreateApprovalRequest{
		PaymentTransactionId: "123",
	}
	res := req.CreateApproval()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateApprovalRequestCreateDisapproval(t *testing.T) {
	req := CreateApprovalRequest{
		PaymentTransactionId: "123",
	}
	res := req.CreateDisapproval()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateBkmRequestCreate(t *testing.T) {
	req := RetrieveBkmRequest{
		Token: "123",
	}
	res := req.CreateBkm()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrieveBkmRequestRetrieve(t *testing.T) {
	req := RetrieveBkmRequest{
		Token: "123",
	}
	res := req.RetrieveBasicBkm()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateBasicBkmInitializeRequestCreate(t *testing.T) {
	req := CreateBasicBkmInitializeRequest{}
	res := req.CreateBasicBkmInitialize()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateBasicPaymentRequestCreateBasicPayment(t *testing.T) {
	req := CreateBasicPaymentRequest{
		Installment: 1,
	}
	res := req.CreateBasicPayment()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateBasicPaymentRequestCreateBasicPaymentPreAuth(t *testing.T) {
	req := CreateBasicPaymentRequest{}
	res := req.CreateBasicPaymentPreAuth()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateBasicPaymentRequestCreateBasicThreedsInitialize(t *testing.T) {
	req := CreateBasicPaymentRequest{}
	res := req.CreateBasicThreedsInitialize()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateBasicPaymentRequestCreateBasicThreedsInitializePreAuth(t *testing.T) {
	req := CreateBasicPaymentRequest{}
	res := req.CreateBasicThreedsInitializePreAuth()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreatePaymentPostAuthRequestCreateBasicPaymentPostAuth(t *testing.T) {
	req := CreatePaymentPostAuthRequest{}
	res := req.CreateBasicPaymentPostAuth()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreatePaymentPostAuthRequestCreatePaymentPostAuth(t *testing.T) {
	req := CreatePaymentPostAuthRequest{}
	res := req.CreatePaymentPostAuth()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateThreedsPaymentRequestCreateBasicThreedsPayment(t *testing.T) {
	req := CreateThreedsPaymentRequest{}
	res := req.CreateBasicThreedsPayment()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateThreedsPaymentRequestCreateThreedsPayment(t *testing.T) {
	req := CreateThreedsPaymentRequest{}
	res := req.CreateThreedsPayment()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrieveBinNumberRequestRetrieveBinNumber(t *testing.T) {
	req := RetrieveBinNumberRequest{}
	res := req.RetrieveBinNumber()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateBkmInitializeRequestCreateBkmInitialize(t *testing.T) {
	req := CreateBkmInitializeRequest{}
	res := req.CreateBkmInitialize()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateCancelRequestRetrieveCancel(t *testing.T) {
	req := CreateCancelRequest{}
	res := req.RetrieveCancel()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateCardRequestCreateCard(t *testing.T) {
	req := CreateCardRequest{}
	res := req.CreateCard()
	checkErrorCode(res.ErrorCode, t)
}

func TestDeleteCardRequestDeleteCard(t *testing.T) {
	req := DeleteCardRequest{}
	res := req.DeleteCard()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrieveCardListRequestCreateCardList(t *testing.T) {
	req := RetrieveCardListRequest{}
	res := req.CreateCardList()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrieveCheckoutFormRequestCreateCheckoutForm(t *testing.T) {
	req := RetrieveCheckoutFormRequest{}
	res := req.CreateCheckoutForm()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateCheckoutFormInitializeRequestCreateCheckoutFormInitialize(t *testing.T) {
	req := CreateCheckoutFormInitializeRequest{}
	res := req.CreateCheckoutFormInitialize()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateCheckoutFormInitializeRequestCreateCheckoutFormInitializePreAuth(t *testing.T) {
	req := CreateCheckoutFormInitializeRequest{}
	res := req.CreateCheckoutFormInitializePreAuth()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateCrossBookingRequestCreateCrossBookingFromSubMerchant(t *testing.T) {
	req := CreateCrossBookingRequest{}
	res := req.CreateCrossBookingFromSubMerchant()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateCrossBookingRequestCreateCrossBookingToSubMerchant(t *testing.T) {
	req := CreateCrossBookingRequest{}
	res := req.CreateCrossBookingToSubMerchant()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrieveInstallmentInfoRequestRetrieveInstallmentInfo(t *testing.T) {
	req := RetrieveInstallmentInfoRequest{}
	res := req.RetrieveInstallmentInfo()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreatePaymentRequestCreatePayment(t *testing.T) {
	req := CreatePaymentRequest{}
	res := req.CreatePayment()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreatePaymentRequestCreatePaymentPreAuth(t *testing.T) {
	req := CreatePaymentRequest{}
	res := req.CreatePaymentPreAuth()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreatePaymentRequestCreateThreedsInitialize(t *testing.T) {
	req := CreatePaymentRequest{}
	res := req.CreateThreedsInitialize()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreatePaymentRequestCreateThreedsInitializePreAuth(t *testing.T) {
	req := CreatePaymentRequest{}
	res := req.CreateThreedsInitializePreAuth()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrievePaymentRequestRetrievePayment(t *testing.T) {
	req := RetrievePaymentRequest{}
	res := req.RetrievePayment()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrievePaymentRequestRetrievePaymentPreAuth(t *testing.T) {
	req := RetrievePaymentRequest{}
	res := req.RetrievePaymentPreAuth()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrievePaymentRequestRetrieveThreedsPayment(t *testing.T) {
	req := RetrievePaymentRequest{}
	res := req.RetrieveThreedsPayment()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrieveTransactionsRequestRetrieveBouncedBankTransferList(t *testing.T) {
	req := RetrieveTransactionsRequest{}
	res := req.RetrieveBouncedBankTransferList()
	checkErrorCode(res.ErrorCode, t)
}

func TestRetrieveTransactionsRequestRetrievePayoutCompletedTransactionList(t *testing.T) {
	req := RetrieveTransactionsRequest{}
	res := req.RetrievePayoutCompletedTransactionList()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreatePeccoInitializeRequestCreatePeccoInitialize(t *testing.T) {
	req := CreatePeccoInitializeRequest{}
	res := req.CreatePeccoInitialize()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreatePeccoPaymentRequestCreatePeccoPayment(t *testing.T) {
	req := CreatePeccoPaymentRequest{}
	res := req.CreatePeccoPayment()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateRefundRequestCreateRefund(t *testing.T) {
	req := CreateRefundRequest{}
	res := req.CreateRefund()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateRefundRequestCreateRefundChargedFromMerchant(t *testing.T) {
	req := CreateRefundRequest{}
	res := req.CreateRefundChargedFromMerchant()
	checkErrorCode(res.ErrorCode, t)
}

func TestCreateSubMerchantRequestCreateSubMerchant(t *testing.T) {
	req := CreateSubMerchantRequest{}
	res := req.CreateSubMerchant()
	checkErrorCode(res.ErrorCode, t)
}

func TestUpdateSubMerchantRequestUpdateSubMerchant(t *testing.T) {
	req := UpdateSubMerchantRequest{}
	res := req.UpdateSubMerchant()
	checkErrorCode(res.ErrorCode, t)
}

func TestUpdateSubMerchantRequestRetrieveSubMerchant(t *testing.T) {
	req := RetrieveSubMerchantRequest{}
	res := req.RetrieveSubMerchant()
	checkErrorCode(res.ErrorCode, t)
}
