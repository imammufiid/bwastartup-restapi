package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	GetCampaignTransactions(c *gin.Context)
	GetUserTransactions(c *gin.Context)
}

type transactionHandler struct {
	service transaction.Service
}

func InstanceTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service: service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	// get data uri
	var input transaction.GetCampaignTransactionsInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to bind URI", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of campaign's transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to get user's transactions", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("User's transaction", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}