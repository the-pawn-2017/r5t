package swaggerui

import (
	"embed"
	"errors"
	"io"
	"io/fs"
	"log"
	"net/http"
	"r5t/spec"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed dist/*
var dist embed.FS

func GenSpec(spec *spec.Spec) echo.HandlerFunc {
	return func(c echo.Context) error {
		re, _ := spec.MarshalJSON()
		return c.JSONBlob(http.StatusOK, re)
	}
}
func GenSwaggerUI(swaggerJSONUrl string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the requested path from the URL.
		requestedPath := c.Param("*")
		// Open the requested file from the embedded filesystem.
		file, err := dist.Open("dist/" + requestedPath)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) { // 注意这里使用errors.Is()和fs.ErrNotExist
				log.Println("文件未找到")
				return echo.ErrNotFound
			}
		}
		defer file.Close()

		// Read the file contents.
		content, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		if requestedPath == "swagger-initializer.js" {
			content = []byte(strings.ReplaceAll(string(content), "./swagger.json", swaggerJSONUrl))
		}
		sRe := strings.Split(requestedPath, ".")

		var fileKind string
		switch sRe[len(sRe)-1] {
		case "css":
			fileKind = "text/css"
		case "html":
			fileKind = "text/html"
		case "js":
			fileKind = "application/x-javascript"
		case "png":
			fileKind = "image/png"
		}
		return c.Blob(http.StatusOK, fileKind, content)
	}
}
