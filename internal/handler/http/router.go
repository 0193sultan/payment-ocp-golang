package http

import (
	"github.com/0193sultan/payment-ocp-golang/internal/factory/payment"
	"github.com/0193sultan/payment-ocp-golang/internal/service/payment"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Factory + register processors
	factory := payment.NewPaymentFactory()
	factory.Register(&payment.BkashProcessor{})
	factory.Register(&payment.NagadProcessor{})

	// Service
	service := payment.NewPaymentService(factory)

	// Handler
	handler := NewPaymentHandler(service)

	api := r.Group("/api/v1")
	{
		api.POST("/pay", handler.Pay)
	}

	return r
}
