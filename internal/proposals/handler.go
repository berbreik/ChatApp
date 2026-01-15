package proposals

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Service *Service
}

func NewHandler(db *sqlx.DB) *Handler {
	repo := NewRepository(db)
	return &Handler{Service: NewService(repo)}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/projects/:id/proposals", h.SubmitProposal)
	r.GET("/projects/:id/proposals", h.ListProposals)
	r.PUT("/proposals/:id/accept", h.AcceptProposal)
	r.PUT("/proposals/:id/reject", h.RejectProposal)
}

type SubmitProposalRequest struct {
	FreelancerID string `json:"freelancer_id" binding:"required"`
	Content      string `json:"content" binding:"required"`
}

func (h *Handler) SubmitProposal(c *gin.Context) {
	projectID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project UUID"})
		return
	}

	var req SubmitProposalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	freelancerID, err := uuid.Parse(req.FreelancerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid freelancer UUID"})
		return
	}

	proposal, err := h.Service.SubmitProposal(projectID, freelancerID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, proposal)
}

func (h *Handler) ListProposals(c *gin.Context) {
	projectID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project UUID"})
		return
	}

	proposals, err := h.Service.Repo.ListByProject(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, proposals)
}

func (h *Handler) AcceptProposal(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid proposal UUID"})
		return
	}

	if err := h.Service.AcceptProposal(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Proposal accepted"})
}

func (h *Handler) RejectProposal(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid proposal UUID"})
		return
	}

	if err := h.Service.RejectProposal(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Proposal rejected"})
}
