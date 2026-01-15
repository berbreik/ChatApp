package chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/chat/:proposalID/channel", h.CreateChannel)
	r.POST("/chat/:proposalID/message", h.SendMessage)
}

func (h *Handler) CreateChannel(c *gin.Context) {
	proposalID := c.Param("proposalID")
	var req struct {
		Members []string `json:"members"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channel, err := h.Service.CreateChannel(proposalID, req.Members)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"channel_id": channel.ID})
}

func (h *Handler) SendMessage(c *gin.Context) {
	proposalID := c.Param("proposalID")
	var req struct {
		UserID string `json:"user_id"`
		Text   string `json:"text"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channel, err := h.Service.Client.Channel("messaging", proposalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.SendMessage(channel, req.UserID, req.Text); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sent"})
}
