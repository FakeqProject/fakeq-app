package dto

import "github.com/FakeqProject/fakeq-app/model"

type ImageResponseDto struct {
	ID   uint   `json:"id"`
	CDN  string `json:"cdn"`
	URL  string `json:"url"`
	Size int    `json:"size"`
}

type ImageRequestDto struct {
	CDN string `json:"cdn"`
}

func CreateImageResponseDto(image model.Image) ImageResponseDto {
	return ImageResponseDto{
		ID:   image.ID,
		CDN:  image.ImageCDN,
		URL:  image.ImageURL,
		Size: image.ImageSize,
	}
}
