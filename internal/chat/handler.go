package chat

import (
	stream "github.com/GetStream/stream-chat-go/v7"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(client *stream.Client) *Handler {
	return &Handler{Service: NewService(client)}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/chat/:channelID/create", h.CreateChannel)
	r.POST("/chat/:channelID/message", h.SendMessage)
	r.POST("/chat/:channelID/members/add", h.AddMembers)
	r.POST("/chat/:channelID/members/remove", h.RemoveMembers)
}

type CreateChannelRequest struct {
	CreatorID string   `json:"creator_id" binding:"required"`
	Members   []string `json:"members" binding:"required"`
}

func (h *Handler) CreateChannel(c *gin.Context) {
	channelID := c.Param("channelID")
	var req CreateChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channel, err := h.Service.CreateChannel(
		c.Request.Context(),
		"messaging",
		channelID,
		req.CreatorID,
		req.Members,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"channel_id": channel.Channel.ID})
}

type SendMessageRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

func (h *Handler) SendMessage(c *gin.Context) {
	channelID := c.Param("channelID")
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.SendMessage(c.Request.Context(), "messaging", channelID, req.UserID, req.Text); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sent"})
}

type MembersRequest struct {
	Members []string `json:"members" binding:"required"`
}

func (h *Handler) AddMembers(c *gin.Context) {
	channelID := c.Param("channelID")
	var req MembersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.AddMembers(c.Request.Context(), "messaging", channelID, req.Members); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "members added"})
}

func (h *Handler) RemoveMembers(c *gin.Context) {
	channelID := c.Param("channelID")
	var req MembersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.RemoveMembers(c.Request.Context(), "messaging", channelID, req.Members); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "members removed"})
}
