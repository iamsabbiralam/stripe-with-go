# Stripe Payment Integration with Go (Gin + GORM)

This repository demonstrates a simple payment flow integration using the **Stripe API** with the **Go programming language**, built using **Gin** as the HTTP router and **GORM** as the ORM layer.

## ✨ Features

- Create and confirm Stripe payment intents.
- Manage and fetch customers from Stripe.
- Save payment data to a PostgreSQL database.
- Well-structured MVC pattern for scalability.
- Clean and simple RESTful API.

## 📂 Project Structure
```
stripe-with-go/
├── pkg/
│ └── payment/
│ ├── controllers/ # Handles HTTP logic
│ ├── models/ # GORM model for payments
│ ├── repositories/ # Handles DB and Stripe logic
├── routes/ # Gin route definitions
├── utils/ # Common utilities (error formatter etc.)
```

## 🔧 Technologies Used

- **Go 1.20+**
- **Gin Web Framework**
- **GORM**
- **Stripe Go SDK (v78)**
- **PostgreSQL**
- **UUID for order tracking**

## 📦 Installation

1. Clone the repository
```bash
git clone https://github.com/your-username/stripe-with-go.git
cd stripe-with-go
```

2. Set environment variables in a .env file:
```
STRIPE_SECRET_KEY=your_stripe_secret_key
CUSTOMER_EMAIL=test@example.com
```

3. Run the app:
```
go run main.go
```

## 🧪 API Endpoint
```
POST /payment/create
```
Creates a Stripe payment intent and stores the order in the database.

Request Body

```
{
  "amount": 99.99
}
```
Response
```
{
  "invoice": "https://receipt.stripe.com/...",
  "message": "✅ Order created successfully after payment!"
}
```

## 📝 Notes
- For simplicity, the customer email is taken from environment variables, but in a real-world app it should come from the authenticated user or request.
- Stripe test keys and test card numbers can be used for simulation.

## 📄 License
MIT License - feel free to use and contribute!

