// @title Music Library API
// @version 1.0
// @description This is an API for managing a music library, including adding, updating, deleting, and fetching songs with lyrics.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /

package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/srmbackisdeveloper/test-music-info/config"
	"github.com/srmbackisdeveloper/test-music-info/internal/handlers"
	"github.com/srmbackisdeveloper/test-music-info/internal/repositories"
	"github.com/srmbackisdeveloper/test-music-info/internal/services"
	"github.com/srmbackisdeveloper/test-music-info/pkg/logger"

	_ "github.com/srmbackisdeveloper/test-music-info/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    appLog := logger.New(cfg.LogLevel)
    appLog.Infof("Configuration loaded successfully: %+v", cfg)

    // permanent repo (postgres): musicRepo
    appLog.Debug("Connecting to PostgreSQL...")
    db, err := repositories.NewPostgresDB(cfg.PostgresDSN)
    if err != nil {
        appLog.Fatalf("failed to connect to PostgreSQL: %v", err)
    }
    defer func() {
        sqlDB, _ := db.DB()
        sqlDB.Close()
        appLog.Debug("PostgreSQL connection closed")
    }()
    appLog.Infof("Connected to PostgreSQL successfully")

    musicRepo := repositories.NewMusicRepository(db)

    // cache repo (redis): cacheRepo
    appLog.Debug("Connecting to Redis...")
    cache := repositories.NewRedisClient(cfg.RedisAddress, cfg.RedisPassword, cfg.RedisDB)
    defer cache.Close()
    appLog.Infof("Connected to Redis successfully")

    cacheRepo := repositories.NewCacheRepository(cache)
    
    // services
    appLog.Debug("Initializing music service...")
    musicService := services.NewMusicService(musicRepo, cacheRepo, 4*time.Hour)
    appLog.Infof("Music service initialized successfully")

    // handlers
    appLog.Debug("Initializing handlers...")
    musicHandler := handlers.NewMusicHandler(musicService)
    appLog.Infof("Handlers initialized successfully")

    // server
    appLog.Debug("Setting up server routes...")
    router := gin.Default()
    port := cfg.Port

    
    router.GET("/info", musicHandler.GetSong)
	router.POST("/music", musicHandler.AddSong)
	router.GET("/music", musicHandler.ListSongs)
	router.PUT("/music/:id", musicHandler.UpdateSong)
	router.DELETE("/music/:id", musicHandler.DeleteSong)

    // show lyrics
    router.GET("/lyrics/:id", musicHandler.GetLyrics)
    // swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    appLog.Infof("Server routes setup complete")


    appLog.Infof("Starting server on port %s", port)
    if err := router.Run(":" + port); err != nil {
        appLog.Fatalf("failed to start server: %v", err)
    }
}
