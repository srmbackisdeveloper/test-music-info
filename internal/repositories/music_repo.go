package repositories

import (
	"log"

	"github.com/srmbackisdeveloper/test-music-info/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MusicRepository struct {
	DB *gorm.DB
}

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
        return nil, err
    }
	
	if err := sqlDB.Ping(); err != nil {
        return nil, err
    }

	log.Println("Connected to PostgreSQL")
    
    if err := db.AutoMigrate(&models.Music{}); err != nil {
        return nil, err
    }
    log.Println("Database migrated successfully")

    return db, nil
}

func NewMusicRepository(db *gorm.DB) *MusicRepository {
    return &MusicRepository{DB: db}
}

// methods:
func (repo *MusicRepository) AddSong(song *models.Music) error {
    return repo.DB.Create(song).Error
}

func (repo *MusicRepository) GetSong(group, title string) (*models.Music, error) {
    var song models.Music
    err := repo.DB.Where("group_name = ? AND title = ?", group, title).First(&song).Error
    if err != nil {
        return nil, err
    }
    return &song, nil
}

func (repo *MusicRepository) UpdateSong(song *models.Music) error {
    return repo.DB.Save(song).Error
}

func (repo *MusicRepository) DeleteSong(id uint) error {
    return repo.DB.Delete(&models.Music{}, id).Error
}

func (repo *MusicRepository) ListSongs(filter map[string]interface{}, limit, offset int) ([]models.Music, error) {
	var songs []models.Music

	query := repo.DB.Model(&models.Music{})
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}

	err := query.Limit(limit).Offset(offset).Find(&songs).Error
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (repo *MusicRepository) CountSongs(filter map[string]interface{}) (int, error) {
	var count int64

	query := repo.DB.Model(&models.Music{})
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}


func (repo *MusicRepository) GetSongByID(id uint) (*models.Music, error) {
	var song models.Music
	err := repo.DB.First(&song, id).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}

// populate some data: Обогащенную информацию положить в БД postgres 
func (repo *MusicRepository) SeedData() error {
	// Define sample songs
	songs := []models.Music{
		{Group: "Muse", Title: "Supermassive Black Hole", Text: "Ooh baby,\n\n don't you know I suffer?", Link: "https://www.example.com/song1"},
		{Group: "Muse", Title: "Hysteria", Text: "It's bugging me,\n\n grating me", Link: "https://www.example.com/song2"},
		{Group: "Radiohead", Title: "Creep", Text: "I'm a creep,\n\n I'm a weirdo", Link: "https://www.example.com/song3"},
		{Group: "Queen", Title: "Bohemian Rhapsody", Text: "Is this the real life?\n\n Is this just fantasy?", Link: "https://www.example.com/song4"},
		{Group: "Queen", Title: "We Will Rock You", Text: "Buddy,\n\n you're a boy,\n\n make a big noise", Link: "https://www.example.com/song5"},
		{Group: "The Beatles", Title: "Hey Jude", Text: "Hey Jude,\n\n don't make it bad", Link: "https://www.example.com/song6"},
		{Group: "The Beatles", Title: "Let It Be", Text: "When I find myself in times of trouble", Link: "https://www.example.com/song7"},
		{Group: "Coldplay", Title: "Fix You", Text: "When you try your best but you don't succeed", Link: "https://www.example.com/song8"},
		{Group: "Coldplay", Title: "Yellow", Text: "Look at the stars,\n\n look how they shine for you", Link: "https://www.example.com/song9"},
		{Group: "Pink Floyd", Title: "Comfortably Numb", Text: "Hello?\n\n Is there anybody in there?", Link: "https://www.example.com/song10"},
	}

	for _, song := range songs {
		existingSong, err := repo.GetSong(song.Group, song.Title)
		if err == nil && existingSong != nil {
			continue
		}

		if err := repo.AddSong(&song); err != nil {
			return err
		}
	}

	log.Println("SeedData: Database seeded successfully with sample songs")
	return nil
}
