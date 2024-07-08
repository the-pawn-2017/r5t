package swaggerui

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"r5t/spec"
	"strings"

	"github.com/labstack/echo/v4"
)

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
		fmt.Println(requestedPath, "!!!!!")
		// Open the requested file from the embedded filesystem.
		file, err := os.Open("../../swaggerui/dist/" + requestedPath)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) { // 注意这里使用errors.Is()和fs.ErrNotExist
				log.Println("文件未找到")
				return echo.ErrNotFound
			}
		}
		defer file.Close()

		// Get the file info to determine its content type.
		info, err := file.Stat()
		if err != nil {
			return err
		}

		// Read the file contents.
		content, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		if requestedPath == "swagger-initializer.js" {
			content = []byte(strings.ReplaceAll(string(content), "./swagger.json", swaggerJSONUrl))
		}

		// Return the file contents with the correct content type.
		return c.Blob(http.StatusOK, info.Mode().String(), content)
	}
}
