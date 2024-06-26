package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api-go/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID!"})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save event!"})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID!"})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch Event!"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to update this event!"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event!"})
		return
	}

	updatedEvent.ID = id
	updatedEvent.Update()
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID!"})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch Event!"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to delete this event!"})
		return
	}

	event.Delete()

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
