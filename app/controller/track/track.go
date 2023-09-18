package track

import (
	"be-map-test/app/db"
	"be-map-test/app/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Get a tracking data by name
// @Description Get a tracking data by name
// @Tags Manual Track
// @Param name path string true "Name"
// @Success 200 {object} models.TrackGet
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /track/{name} [get]
func GetAllByName(c *gin.Context) {

	var TrackData []*models.TrackGet
	var count int64
	name := c.Param("name")
	db.DB.Where("is_deleted = ?", false).Find(&TrackData)
	db.DB.Model(&TrackData).Where("is_deleted = ?", false).Where("name = ?", name).Count(&count)
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Total Data": count, "data": "Empty",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Total Data": count, "data": TrackData})

}

// @Summary Create a manual track data
// @Description Create a manual  track data
// @Tags Manual Track
// @Param create_track body models.Tracks true "Track"
// @Success 200 {object} models.Tracks
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /track [post]
func Create(c *gin.Context) {

	var Track *models.Track

	if err := c.ShouldBindJSON(&Track); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newTrack := models.Track{
		TrackID:        Track.TrackID,
		TrackName:      Track.TrackName,
		Name:           Track.Name,
		AdditionalInfo: Track.AdditionalInfo,
		UpdatedAt:      now,
	}
	log.Print(newTrack)
	if err := db.DB.Create(&newTrack).Error; err != nil {
		fmt.Println("Error creating user:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": newTrack})
}

// @Summary Update a manual track data
// @Description Update a manual track data
// @Tags Manual Track
// @Param        id    query     string  true  "query by id"
// @Param Update_manual_track body models.Tracks true "Track"
// @Success 200 {object} models.Tracks
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /track [put]
func Update(c *gin.Context) {
	var Track models.Track
	trackID := c.Query("id")
	if err := c.ShouldBindJSON(&Track); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if db.DB.Table("track_data").Where("id = ?", trackID).Updates(&Track).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "cannot update track"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "track updated!", "data": Track})

}

// @Summary Delete a manual track data
// @Description Delete a manual track data based on ID
// @Tags Manual Track
// @Param        id    query     string  true  "query by id"
// @Success 204 "No Content"
// @Failure 404 {object} ErrorResponse
// @Router /track [delete]
func Delete(c *gin.Context) {

	var Track models.Track
	id := c.Query("id")

	if db.DB.Model(&Track).Where("id = ?", id).Update("is_deleted", true).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Fail to delete data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Deleted!"})
}

type ErrorResponse struct {
	Message string `json:"message"`
}
