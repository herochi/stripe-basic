package main

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"log"
	"net/http"
)

const (
	priceIDSusBas  string = "price_id_de_suscripcion_basica"
	priceIDSusPrem string = "price_id_de_suscripcion_premium"
	priceIDUnique  string = "price_id_de_producto"
)

func main() {
	stripe.Key = "secret_de_nuestra_cuenta"

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/create-checkout-session-sus-basic", createCheckoutSessionSuscriptionBasic)
	http.HandleFunc("/create-checkout-session-sus-prem", createCheckoutSessionSuscriptionPremium)
	http.HandleFunc("/create-checkout-session-unique", createCheckoutSessionProducUnique)
	addr := "localhost:4242"
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func createCheckoutSessionSuscriptionBasic(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(priceIDSusBas),
				Quantity: stripe.Int64(1),
			},
		},
		//Mode me sirve para especificar si es una compra única o es para una suscripción
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(domain + "/success.html"),
		CancelURL:  stripe.String(domain + "/cancel.html"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}

func createCheckoutSessionSuscriptionPremium(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(priceIDSusPrem),
				Quantity: stripe.Int64(1),
			},
		},
		//Mode me sirve para especificar si es una compra única o es para una suscripción
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String(domain + "/success.html"),
		CancelURL:  stripe.String(domain + "/cancel.html"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}

func createCheckoutSessionProducUnique(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(priceIDUnique),
				Quantity: stripe.Int64(1),
			},
		},
		//Mode me sirve para especificar si es una compra única o es para una suscripción
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success.html"),
		CancelURL:  stripe.String(domain + "/cancel.html"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
