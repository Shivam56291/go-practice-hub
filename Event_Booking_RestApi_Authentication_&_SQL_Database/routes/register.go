package routes

import (
	"event-booking-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID.", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get event.", "error": err.Error()})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register for event.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registered for event successfully."})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userId")

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID.",
			"error":   err.Error(),
		})
		return
	}

	// Ensure event exists
	if _, err := models.GetEventByID(eventID); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found.",
			"error":   err.Error(),
		})
		return
	}

	event := models.Event{ID: eventID}
	if err := event.CancelRegistration(userID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cancel registration.",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Registration cancelled successfully.",
	})
}
