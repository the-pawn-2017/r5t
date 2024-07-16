package swaggerui

import (
	"embed"
	"errors"
	"io"
	"io/fs"
	"strings"

	"github.com/the-pawn-2017/r5t"
)

//go:embed dist/*
var dist embed.FS

func GenSpec(spec *r5t.Spec) ([]byte, error) {
	return spec.MarshalJSON()
}
func GetSwaggerUIFile(swaggerJSONUrl string, fileName string) (string, []byte, error) {
	requestedPath := fileName
	// Open the requested file from the embedded filesystem.
	file, err := dist.Open("dist/" + requestedPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) { // 注意这里使用errors.Is()和fs.ErrNotExist
			return "", nil, err
		}
	}
	defer file.Close()

	// Read the file contents.
	content, err := io.ReadAll(file)
	if err != nil {
		return "", nil, err
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
	return fileKind, content, err
}
