package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := getEnv("PORT", "80")
	staticDir := getEnv("STATIC_DIR", "./www")
	staticPath := getEnv("STATIC_PATH", "/")
	fsEnabled := getEnv("FS_ENABLED", "no")
	fsDir := getEnv("FS_DIR", "./files")
	fsPath := getEnv("FS_PATH", "/files")

	r := gin.Default()

	if fsEnabled == "yes" {
		r.StaticFS(fsPath, http.Dir(fsDir))
	}

	r.Static(staticPath, staticDir)

	r.Run(":" + port)
}

// getEnv gets an environment variable or sets a default if
// one does not exist.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
