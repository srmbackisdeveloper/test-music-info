package tests

import (
	"testing"

	"github.com/srmbackisdeveloper/test-music-info/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockMusicService struct {
	mock.Mock
}

func (m *MockMusicService) AddSong(song *models.Music) error {
	args := m.Called(song)
	return args.Error(0)
}

func TestAddSong(t *testing.T) {
	// mockService := new(MockMusicService)
	// handler := handlers.NewMusicHandler(mockService)
	// router := SetupTestRouter()
	// router.POST("/music", handler.AddSong)

	// testSong := &models.Music{Group: "Muse", Title: "Supermassive Black Hole"}
	// mockService.On("AddSong", testSong).Return(nil)

	// w := PerformRequest(router, "POST", "/music", []byte(`{"group": "Muse", "song": "Supermassive Black Hole"}`))
	// assert.Equal(t, http.StatusCreated, w.Code)

	// var response map[string]interface{}
	// err := json.Unmarshal(w.Body.Bytes(), &response)
	// assert.NoError(t, err)
	// assert.Equal(t, "Song added successfully", response["message"])
	// mockService.AssertExpectations(t)
}
