package dto

import "github.com/FakeqProject/fakeq-app/model"

type FileResponseDto struct {
	ID   uint   `json:"id"`
	CDN  string `json:"cdn"`
	URL  string `json:"url"`
	Size int    `json:"size"`
}

type FileRequestDto struct {
	CDN string `json:"cdn"`
}

func CreateFileResponseDto(file model.File) FileResponseDto {
	return FileResponseDto{
		ID:   file.ID,
		CDN:  file.FileCDN,
		URL:  file.FileURL,
		Size: file.FileSize,
	}
}
