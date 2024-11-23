package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/srmbackisdeveloper/test-music-info/internal/models"
	"github.com/srmbackisdeveloper/test-music-info/internal/services"
	"github.com/srmbackisdeveloper/test-music-info/internal/types"
	"gorm.io/gorm"
)

type MusicHandler struct {
	MusicService *services.MusicService
}

func NewMusicHandler(musicService *services.MusicService) *MusicHandler {
	return &MusicHandler{MusicService: musicService}
}

// AddSong godoc
// @Summary Add a new song
// @Description Adds a new song to the database
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body types.AddSongRequest true "Request to add a song"
// @Success 201 {object} models.Music "The added song"
// @Failure 400 {object} types.ErrorResponse "Invalid request payload"
// @Failure 500 {object} types.ErrorResponse "Failed to add the song to the database"
// @Router /music [post]
func (h *MusicHandler) AddSong(c *gin.Context) {
	log.Println("AddSong: Received request to add a song")
	var req types.AddSongRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("AddSong: Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: 'group' and 'song' are required"})
		return
	}

	log.Printf("AddSong: Adding song with group '%s' and title '%s'", req.Group, req.Title)
	newSong := &models.Music{
		Group: req.Group,
		Title: req.Title,
	}

	err := h.MusicService.AddSong(newSong)
	if err != nil {
		log.Println("AddSong: Failed to add song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add song"})
		return
	}

	log.Println("AddSong: Song added successfully")
	c.JSON(http.StatusCreated, gin.H{"message": "Song added successfully", "song": newSong})
}

// GetSong godoc
// @Summary Retrieve a song
// @Description Fetches a song by its group and title
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string true "The group of the song"
// @Param song query string true "The title of the song"
// @Success 200 {object} types.SongDetail "The requested song"
// @Failure 400 {object} types.ErrorResponse "Invalid or missing query parameters"
// @Failure 404 {object} types.ErrorResponse "Song not found"
// @Router /info [get]
func (h *MusicHandler) GetSong(c *gin.Context) {
	log.Println("GetSong: Received request to fetch a song")
	group := c.Query("group")
	song := c.Query("song")

	if group == "" || song == "" {
		log.Println("GetSong: Missing required query parameters 'group' or 'song'")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group and song are required"})
		return
	}

	log.Printf("GetSong: Fetching song with group '%s' and title '%s'", group, song)
	gotSong, err := h.MusicService.GetSong(group, song)
	if err != nil {
		log.Println("GetSong: Song not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	log.Println("GetSong: Song fetched successfully")
	c.JSON(http.StatusOK, gotSong)
}

// UpdateSong godoc
// @Summary Update a song
// @Description Updates an existing song by its ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "The ID of the song to update"
// @Param song body models.Music true "Updated song details"
// @Success 200 {object} models.Music "The updated song"
// @Failure 400 {object} types.ErrorResponse "Invalid song ID or payload"
// @Failure 404 {object} types.ErrorResponse "Song not found"
// @Failure 500 {object} types.ErrorResponse "Failed to update the song"
// @Router /music/{id} [put]
func (h *MusicHandler) UpdateSong(c *gin.Context) {
	log.Println("UpdateSong: Received request to update a song")
	idParam := c.Param("id")
	songID, err := strconv.Atoi(idParam)
	if err != nil || songID <= 0 {
		log.Println("UpdateSong: Invalid song ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	log.Printf("UpdateSong: Checking existence of song with ID %d", songID)
	existingSong, err := h.MusicService.GetSongByID(uint(songID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("UpdateSong: Song not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		}
		log.Println("UpdateSong: Failed to fetch song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch song"})
		return
	}

	var req models.Music
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UpdateSong: Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	log.Printf("UpdateSong: Updating song with ID %d", songID)
	existingSong.Group = req.Group
	existingSong.Title = req.Title
	existingSong.ReleaseDate = req.ReleaseDate
	existingSong.Text = req.Text
	existingSong.Link = req.Link

	err = h.MusicService.UpdateSong(existingSong)
	if err != nil {
		log.Println("UpdateSong: Failed to update song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
		return
	}

	log.Println("UpdateSong: Song updated successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully", "song": existingSong})
}

// DeleteSong godoc
// @Summary Delete a song
// @Description Deletes a song by its ID
// @Tags Songs
// @Param id path int true "The ID of the song to delete"
// @Success 200 {object} types.MessageResponse "Deletion success message"
// @Failure 400 {object} types.ErrorResponse "Invalid song ID"
// @Failure 404 {object} types.ErrorResponse "Song not found"
// @Failure 500 {object} types.ErrorResponse "Failed to delete the song"
// @Router /music/{id} [delete]
func (h *MusicHandler) DeleteSong(c *gin.Context) {
	log.Println("DeleteSong: Received request to delete a song")
	idParam := c.Param("id")
	songID, err := strconv.Atoi(idParam)
	if err != nil || songID <= 0 {
		log.Println("DeleteSong: Invalid song ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	log.Printf("DeleteSong: Checking existence of song with ID %d", songID)
	_, err = h.MusicService.GetSongByID(uint(songID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("DeleteSong: Song not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		}
		log.Println("DeleteSong: Failed to fetch song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch song"})
		return
	}

	log.Printf("DeleteSong: Deleting song with ID %d", songID)
	err = h.MusicService.DeleteSong(uint(songID))
	if err != nil {
		log.Println("DeleteSong: Failed to delete song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
		return
	}

	log.Println("DeleteSong: Song deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
}

// ListSongs godoc
// @Summary List all songs
// @Description Retrieves a paginated list of songs with optional filters
// @Tags Songs
// @Param group query string false "Filter by group name"
// @Param title query string false "Filter by song title"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Number of songs per page (default: 10)"
// @Success 200 {object} types.PaginatedSongsResponse "Paginated list of songs"
// @Failure 500 {object} types.ErrorResponse "Failed to fetch the list of songs"
// @Router /music [get]
func (h *MusicHandler) ListSongs(c *gin.Context) {
	log.Println("ListSongs: Received request to list songs")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	filter := map[string]interface{}{}
	if group := c.Query("group"); group != "" {
		filter["group_name"] = group
	}
	if title := c.Query("title"); title != "" {
		filter["title"] = title
	}

	log.Printf("ListSongs: Fetching songs with filter %+v, page %d, limit %d", filter, page, limit)
	songs, totalSongs, err := h.MusicService.ListSongs(filter, limit, offset)
	if err != nil {
		log.Println("ListSongs: Failed to list songs")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list songs"})
		return
	}

	totalPages := (totalSongs + limit - 1) / limit
	log.Println("ListSongs: Songs listed successfully")
	c.JSON(http.StatusOK, gin.H{
		"page":       page,
		"limit":      limit,
		"totalPages": totalPages,
		"totalSongs": totalSongs,
		"data":       songs,
	})
}

// GetLyrics godoc
// @Summary Get lyrics of a song
// @Description Retrieves the lyrics of a song in a paginated format
// @Tags Songs
// @Param id path int true "The ID of the song"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Number of verses per page (default: 5)"
// @Success 200 {object} types.PaginatedVersesResponse "Paginated lyrics of the song"
// @Failure 400 {object} types.ErrorResponse "Invalid song ID or pagination parameters"
// @Failure 404 {object} types.ErrorResponse "Song not found"
// @Failure 500 {object} types.ErrorResponse "Failed to fetch the lyrics"
// @Router /lyrics/{id} [get]
func (h *MusicHandler) GetLyrics(c *gin.Context) {
	log.Println("GetLyrics: Received request to fetch lyrics")
	idParam := c.Param("id")
	songID, err := strconv.Atoi(idParam)
	if err != nil || songID <= 0 {
		log.Println("GetLyrics: Invalid song ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}

	log.Printf("GetLyrics: Fetching song with ID %d", songID)
	song, err := h.MusicService.GetSongByID(uint(songID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("GetLyrics: Song not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		}
		log.Println("GetLyrics: Failed to fetch song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve song"})
		return
	}

	verses := splitLyricsIntoVerses(song.Text)
	totalVerses := len(verses)
	start := (page - 1) * limit
	end := start + limit

	if start >= totalVerses {
		log.Println("GetLyrics: No verses for this page")
		c.JSON(http.StatusOK, gin.H{
			"page":       page,
			"limit":      limit,
			"totalPages": (totalVerses + limit - 1) / limit,
			"totalVerses": totalVerses,
			"data":       []string{},
		})
		return
	}
	if end > totalVerses {
		end = totalVerses
	}

	log.Printf("GetLyrics: Returning verses for page %d, limit %d", page, limit)
	c.JSON(http.StatusOK, gin.H{
		"page":       page,
		"limit":      limit,
		"totalPages": (totalVerses + limit - 1) / limit,
		"totalVerses": totalVerses,
		"data":       verses[start:end],
	})
}

func splitLyricsIntoVerses(lyrics string) []string {
	return strings.Split(lyrics, "\n\n")
}
