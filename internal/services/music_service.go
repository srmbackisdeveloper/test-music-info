package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/srmbackisdeveloper/test-music-info/internal/models"
	"github.com/srmbackisdeveloper/test-music-info/internal/repositories"
	"github.com/srmbackisdeveloper/test-music-info/internal/types"
)

type MusicService struct {
	MusicRepo *repositories.MusicRepository
	CacheRepo *repositories.CacheRepository
	CacheTTL time.Duration // time to live
}


func NewMusicService(musicRepo *repositories.MusicRepository, cacheRepo *repositories.CacheRepository, cacheTTL time.Duration) *MusicService {
	return &MusicService{
        MusicRepo:   musicRepo,
        CacheRepo: cacheRepo,
        CacheTTL:  cacheTTL,
    }
}

func (s *MusicService) AddSong(song *models.Music) error {
    if err := s.MusicRepo.AddSong(song); err != nil {
        return err
    }

    return nil
}

func (s *MusicService) GetSong(group, title string) (*types.SongDetail, error) {
    cacheKey := fmt.Sprintf("%s:%s", group, title)

    cachedData, err := s.CacheRepo.GetSongCache(cacheKey)
    if err != nil {
        return nil, err
    }

    if cachedData != "" { // cache hit
        var song models.Music
        if err := json.Unmarshal([]byte(cachedData), &song); err != nil {
            return nil, err
        }

        return &types.SongDetail{
            ReleaseDate: song.ReleaseDate.Format("2006-01-02"),
            Text:        song.Text,
            Link:        song.Link,
        }, nil
    }

    // cache miss
    song, err := s.MusicRepo.GetSong(group, title)
    if err != nil {
        return nil, err
    }

    data, err := json.Marshal(song)
    if err != nil {
        return nil, err
    }
    _ = s.CacheRepo.SetSongCache(cacheKey, string(data), s.CacheTTL)

    return &types.SongDetail{
        ReleaseDate: song.ReleaseDate.Format("2006-01-02"),
        Text:        song.Text,
        Link:        song.Link,
    }, nil
}

func (s *MusicService) UpdateSong(song *models.Music) error {
    if err := s.MusicRepo.UpdateSong(song); err != nil {
        return err
    }

    return nil
}

func (s *MusicService) DeleteSong(id uint) error {
	return s.MusicRepo.DeleteSong(id)
}


func (s *MusicService) ListSongs(filter map[string]interface{}, limit, offset int) ([]models.Music, int, error) {
	// Fetch the filtered and paginated songs
	songs, err := s.MusicRepo.ListSongs(filter, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Fetch the total count of matching songs
	totalSongs, err := s.MusicRepo.CountSongs(filter)
	if err != nil {
		return nil, 0, err
	}

	return songs, totalSongs, nil
}

func (s *MusicService) GetSongByID(id uint) (*models.Music, error) {
	return s.MusicRepo.GetSongByID(id)
}

