package service

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/muratovdias/diplom/internal/app/repository"
)

type Service struct {
	AuthService
	Trainer
	Client
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo.AuthRepository),
		Trainer:     NewTrainerService(repo.Trainer),
		Client:      NewClientService(repo.Client),
	}
}

func CreateImage(fileHeader *multipart.FileHeader) (string, error) {
	// fmt.Println("filename: ", fileHeader.Filename)
	if fileHeader != nil {
		file, err := fileHeader.Open()
		if err != nil {
			return "", err
		}
		fileHeader.Filename = path.Base(fileHeader.Filename)
		if !strings.Contains(fileHeader.Filename, ".jpg") && !strings.Contains(fileHeader.Filename, "png") {
			return "", errors.New("you shou")
		}
		out, err := os.Create("./ui/img/" + fileHeader.Filename)
		if err != nil {
			return "", err
		}
		if _, err := io.Copy(out, file); err != nil {
			return "", err
		}
		image, err := ConvertToBinary("./ui/img/" + fileHeader.Filename)
		if err != nil {
			return "", err
		}

		return image, nil
	}
	image, err := ConvertToBinary("./ui/img/ava.png")
	if err != nil {
		return "", err
	}

	return image, nil
}

func ConvertToBinary(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding, nil
}
