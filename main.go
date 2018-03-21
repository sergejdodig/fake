package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("Listening...")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fakeauthToken", postMembersHandlerH).Methods("POST")
	r.HandleFunc("/fakepayments", postMembersHandler).Methods("POST")
	r.HandleFunc("/fakepayments/{idpstrx}", postMembersHandlerQS).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8001", nil)
}

func postMembersHandlerH(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Requested auth: " + r.RequestURI)

	msgresp := "{\"partnerId\":\"4\",\"memberId\":\"12541\",\"result\":{\"code\":\"200\",\"description\":\"Token generated successfully\"},\"timestamp\":\"2018-03-16 18:54:22\",\"LoginName\":\"payneticsmerchant1\",\"AuthToken\":\"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwYXluZXRpY3NtZXJjaGFudDEiLCJyb2xlIjoibWVyY2hhbnQiLCJpc3MiOiJQWiIsImV4cCI6MTUyMTIxMDI2Mn0.QkFkxghYeJSZxREHVjd8qNgGuO3ylzAXLug1mlotYwM\"}"

	w.Write([]byte(msgresp))
}

func postMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Requested payment: " + r.RequestURI)

	msgresp := "{\"paymentId\": \"10617\",\"paymentType\": \"DB\",\"paymentBrand\": \"VISA\",	\"paymentMode\": \"CC\",	\"amount\": \"2.88\",\"currency\": \"EUR\",	\"descriptor\": \"Paynetics EUR Test\",	\"result\": {\"code\": \"00001\",\"description\": \"Transaction succeeded\"},\"card\": {\"bin\": \"444433\",\"last4Digits\": \"1111\",	\"holder\": \"Sergej Dddd\",\"expiryMonth\": \"12\",\"expiryYear\": \"2030\"},\"timestamp\": \"2018-03-16 22:12:10\",\"merchantTransactionId\": \"29002889643\", \"transactionStatus\": \"P\"}"

	w.Write([]byte(msgresp))
}

func postMembersHandlerQS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Requested query status: " + r.RequestURI)

	msgresp := "{\"paymentId\":\"10617\",\"status\":\"capturesuccess\",\"paymentBrand\":\"VISA\",\"paymentMode\":\"CC\",\"firstName\":\"?ü?i\",\"lastName\":\"Üglü\",\"amount\":\"24.46\",\"currency\":\"EUR\",\"descriptor\":\"Paynetics EUR Test\",\"result\":{\"code\":\"00026\",\"description\":\"Your record found successfully\"},\"card\":{\"bin\":\"401849\",\"last4Digits\":\"0013\",\"holder\":\"?ü?i Üglü\",\"expiryMonth\":\"03\",\"expiryYear\":\"2034\"},\"customer\":{\"email\":\"sandip.k@paymentz.com\"},\"timestamp\":\"2018-03-19 17:58:46\",\"merchantTransactionId\":\"29000769815\",\"remark\":\"Approved or completed successfully\",\"transactionStatus\":\"Y\",\"tmpl_amount\":\"24.46\",\"tmpl_currency\":\"EUR\",\"eci\":\"05\"}"

	w.Write([]byte(msgresp))
}
