package routes

import (
	"net/http"
	"strconv"

	"github.com/berunda/go-rest/models"
	"github.com/gin-gonic/gin"
)

// Events: GET
func getEvents(context *gin.Context) {
	// Get All Events
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}
	// Return Events
	context.JSON(http.StatusOK, events)
}

// Event: POST
func createEvent(context *gin.Context) {
	// Bind JSON Request Body
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	// Create(Save) Event on Database
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}
	// Return new Event
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

// Event: GET
func getEvent(context *gin.Context) {
	// Get Event ID from Param
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messaage": "Could not parse event id."})
		return
	}
	// Get Event By ID
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messaage": "Could not fetch event."})
		return
	}
	// Return Event
	context.JSON(http.StatusOK, event)
}

// Event: PUT
func updateEvent(context *gin.Context) {
	// Get Event ID from Param
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messaage": "Could not parse event id."})
		return
	}
	// Check if Event exists
	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messaage": "Could not fetch event."})
		return
	}
	// Bind JSON Request Body
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messaage": "Could not parse request data."})
		return
	}
	// Update Event
	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messaage": "Could not update event"})
		return
	}
	// Return Success
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

// EVENT: Delete
func deleteEvent(context *gin.Context) {
	// Get Event ID from Param
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"messaage": "Could not parse event id."})
		return
	}
	// Get Event By ID
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messaage": "Could not fetch event."})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messaage": "Could not delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
