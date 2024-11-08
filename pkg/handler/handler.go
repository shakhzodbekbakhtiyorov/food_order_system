package handler

import (
	"gogogo/pkg/service"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/redis.v5"

	_ "gogogo/docs"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services    *service.Service
	redisClient *redis.Client
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	h.redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	router.Static("/static", "./static/uploads")

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		file := api.Group("/file")
		{
			file.POST("/upload", h.uploadFile)
		}

		admin := api.Group("/admin")
		{
			admin.POST("/users", h.createUser)
		}

		products := api.Group("/products")
		{
			products.POST("/", h.createProduct)
			products.PUT("/:id", h.updateProduct)
			products.GET("/", h.getAllProducts)
			products.GET("/:id", h.getProductById)
			products.DELETE("/:id", h.deleteProduct)
		}

		categories := api.Group("/categories")
		{
			categories.POST("/", h.createCategory)
			categories.PUT("/:id", h.updateCategory)
			categories.GET("/", h.getAllCategories)
			categories.GET("/:id", h.getCategoryById)
			categories.DELETE("/:id", h.deleteCategory)
		}

		orders := api.Group("/orders")
		{
			orders.POST("/", h.createOrder)
			orders.PUT("/:id", h.updateOrder)
			orders.GET("/", h.getAllOrders)
			orders.GET("/:id", h.getOrderById)
			orders.DELETE("/:id", h.deleteOder)
		}

		order_tems := api.Group("/order_items")
		{
			order_tems.POST("/", h.createOrderItem)
			order_tems.PUT("/:id", h.updateOrderItem)
			order_tems.DELETE("/:id", h.deleteOderItem)
		}

		payments := api.Group("/payments")
		{
			payments.POST("/", h.createPayment)
			payments.PUT("/:id", h.updatePayment)
			payments.GET("/", h.getAllPayments)
			payments.GET("/:id", h.getPaymentById)
			payments.DELETE("/:id", h.deletePayment)
			payments.GET("/by_order/:order_id", h.getPaymentByOrderId)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
