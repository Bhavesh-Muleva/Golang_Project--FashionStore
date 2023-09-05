package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
)

func main() {

	stripe.Key = "sk_test_value" // paste the secret key from
	// Handler Functions
	http.HandleFunc("/create-payment-intent", handlerCreatePaymentIntent)
	http.HandleFunc("/handle-health", handleHealth)

	fmt.Println("Listening On Port :4242...")

	// Listener
	if err := http.ListenAndServe(":4242", nil); err != nil {
		log.Fatal(err)
	}
}

func handlerCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome")
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(req.ProductId)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(paymentIntent.ClientSecret)

	var response struct {
		ClientSecret string `json:"clientSecret"`
	}
	response.ClientSecret = paymentIntent.ClientSecret
	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(w, &buf)
	if err != nil {
		fmt.Println(err)
	}

}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("temp")
	var response []byte = []byte("Endpoint is up and running")
	if _, err := w.Write(response); err != nil {
		log.Fatal(err)
	}
}

func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 15500
	case "Forever Shorts":
		return 30000
	}
	return 0
}
