package gogogo

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
}

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	CategoryId int     `json:"category_id" db:"category_id"`
	Image      string  `json:"image"`
}

type CreateProductInput struct {
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	CategoryId int     `json:"category_id"`
	Image      string  `json:"image"`
}

type UpdateProductInput struct {
	Name       *string  `json:"name"`
	Price      *float32 `json:"price"`
	CategoryId *int     `json:"category_id"`
	Image      string   `json:"image"`
}

type Category struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type CreateCategoryInput struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Order struct {
	Order      OrderDB     `json:"order"`
	OrderItems []OrderItem `json:"order_items"`
}

type OrderDB struct {
	Id          int `json:"id"`
	OrderTime   int `json:"order_time"`
	Status      int `json:"status"`
	OrderNumber int `json:"order_number"`
	Amount      int `json:"amount"`
}

type UpdateOrderInput struct {
	Status int `json:"status"`
}

type OrderItem struct {
	Id        int `json:"id"`
	Count     int `json:"count"`
	ProductId int `json:"product_id"`
	OrderId   int `json:"order_id"`
	Status    int `json:"status"`
}

type CreateOrderItem struct {
	OrderId   int `json:"order_id"`
	Count     int `json:"count"`
	ProductId int `json:"product_id"`
}

type UpdateOrderItem struct {
	Count  int `json:"count"`
	Status int `json:"status"`
}

type CreateOrderInput struct {
	Items []CreateOrderItem `json:"items"`
}

type Payment struct {
	Id          int     `json:"id"`
	OrderId     int     `json:"order_id"`
	PaymentType int     `json:"payment_type"`
	Amount      float32 `json:"amount"`
	Change      float32 `json:"change"`
}

type CreatePaymentInput struct {
	OrderId     int     `json:"order_id"`
	PaymentType int     `json:"payment_type"`
	Amount      float32 `json:"amount"`
	Change      float32 `json:"change"`
}

type UpdatePaymentInput struct {
	PaymentType int     `json:"payment_type"`
	Amount      float32 `json:"amount"`
	Change      float32 `json:"change"`
}
