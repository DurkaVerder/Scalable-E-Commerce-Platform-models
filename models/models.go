package models

import "time"

// User represents a user in the system with basic authentication details.
type User struct {
	ID       int    `json:"id"`       // Unique identifier for the user.
	Username string `json:"username"` // Username chosen by the user.
	Email    string `json:"email"`    // Email address of the user.
	Password string `json:"password"` // Hashed password for authentication.
}

// Product represents an item available for purchase in the e-commerce platform.
type Product struct {
	ID          int     `json:"id"`          // Unique identifier for the product.
	Name        string  `json:"name"`        // Name of the product.
	Price       float64 `json:"price"`       // Price of the product.
	Description string  `json:"description"` // Description of the product.
	Quantity    int     `json:"quantity"`    // Available quantity of the product in stock.
}

// Notification represents an email notification with a subject and body.
type Notification struct {
	Email   string `json:"email"`   // Recipient's email address.
	Subject string `json:"subject"` // Subject of the email.
	Body    string `json:"body"`    // Body content of the email.
}

// Order represents a customer's order, including details about the user, amount, status, and associated products.
type Order struct {
	ID        int            `json:"id"`         // Unique identifier for the order.
	UserID    int            `json:"user_id"`    // ID of the user who placed the order.
	Amount    float64        `json:"amount"`     // Total amount for the order.
	Status    string         `json:"status"`     // Current status of the order (e.g., pending, completed).
	CreatedAt time.Time      `json:"created_at"` // Timestamp when the order was created.
	Products  []OrderProduct // List of products included in the order.
}

// OrderProduct represents a product within an order, including its name, price, and quantity.
type OrderProduct struct {
	Name     string  `json:"name"`     // Name of the product.
	Price    float64 `json:"price"`    // Price of the product.
	Quantity int     `json:"quantity"` // Quantity of the product in the order.
}
