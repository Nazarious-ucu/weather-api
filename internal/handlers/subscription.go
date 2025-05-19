package handlers

import (
	_ "WeatherSubscriptionAPI/internal/models"
	"WeatherSubscriptionAPI/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SubscriptionHandler struct {
	Service *service.SubscriptionService
}

func NewSubscriptionHandler(svc *service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{Service: svc}
}

// Subscribe
// @Summary Subscribe to weather updates
// @Description Subscribe an email to receive weather updates for a specific city.
// @Tags subscription
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Email address to subscribe"
// @Param city formData string true "City for weather updates"
// @Success 200 {object} model.SubscribeResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /subscribe [post]
func (h *SubscriptionHandler) Subscribe(c *gin.Context) {
	email := c.PostForm("email")
	city := c.PostForm("city")
	if email == "" || city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and city are required"})
		return
	}
	token, err := h.Service.Subscribe(email, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Subscribed successfully. Check your email to confirm.",
		"token":   token, // 👈 додаємо сюди
	})

}

// Confirm
// @Summary Confirm subscription
// @Description Confirms the subscription using the token sent in email.
// @Tags subscription
// @Produce json
// @Param token path string true "Confirmation token"
// @Success 200 {object} model.MessageResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /confirm/{token} [get]
func (h *SubscriptionHandler) Confirm(c *gin.Context) {
	token := c.Param("token")
	ok, err := h.Service.Confirm(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscription confirmed."})
}

// Unsubscribe
// @Summary Unsubscribe
// @Description Unsubscribe from weather updates using the token.
// @Tags subscription
// @Produce json
// @Param token path string true "Unsubscribe token"
// @Success 200 {object} model.MessageResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /unsubscribe/{token} [get]
func (h *SubscriptionHandler) Unsubscribe(c *gin.Context) {
	token := c.Param("token")
	ok, err := h.Service.Unsubscribe(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unsubscribed successfully."})
}
