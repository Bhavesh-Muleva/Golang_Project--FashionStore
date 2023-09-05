# Golang_Project--FashionStore

This is a simple golang project.
In this project i basically created a simple http server in golang and bind it to the already pre-created frontend react component to make logicless website functional.

In this, I have added stripe payment method to make checkout section of website more real.

# Guide to use this project 
make sure you have npm and golang install

1. Clone the project

2. cd frontend
   npm install
   npm start

Note: before moving to next step, create account in stripe account and go to developer setting copy the public key (pk_test_****) and paste it to stripePromise variable in StripePayment.jsx file in frontend folder and then copy the secret key (sk_test_****) into stripe.Key variable created in main.go in backend folder.

3. cd backend
   go run main.go

Backend will start at :4242 Port.

Enjoy! 

if you like this repo do hit star and share...
