package projects

import (
	"net/http"
	"time"

	"chatapp/v/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Repo *Repository
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{Repo: NewRepository(db)}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/projects", h.CreateProject)
	r.GET("/projects/:id", h.GetProject)
	r.GET("/clients/:id/projects", h.ListProjectsByClient)
	r.DELETE("/projects/:id", h.DeleteProject)
}

type CreateProjectRequest struct {
	ClientID    string `json:"client_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (h *Handler) CreateProject(c *gin.Context) {
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientID, err := uuid.Parse(req.ClientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client UUID"})
		return
	}

	project := models.Project{
		ID:          uuid.New(),
		ClientID:    clientID,
		Title:       req.Title,
		Description: req.Description,
		Status:      "draft",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := h.Repo.Create(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *Handler) GetProject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	project, err := h.Repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *Handler) ListProjectsByClient(c *gin.Context) {
	clientID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	projects, err := h.Repo.ListByClient(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *Handler) DeleteProject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	if err := h.Repo.SoftDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
