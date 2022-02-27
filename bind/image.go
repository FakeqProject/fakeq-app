package bind

import (
	"github.com/FakeqProject/fakeq-app/dto"
	"github.com/FakeqProject/fakeq-app/model"
	"github.com/FakeqProject/fakeq-app/service"
	"github.com/webview/webview"
)

func ImageBindCollection(w webview.WebView) {
	w.Bind("getAllImage", GetAllImage())
	w.Bind("getAllImageWithPagination", GetAllImageWithPagination())
	w.Bind("getImageByID", GetImageByID())
	w.Bind("insertNewImage", InsertNewImage())
	w.Bind("updateImageByID", UpdateImageByID())
	w.Bind("deleteImageByID", DeleteImageByID())
}

func GetAllImage() interface{} {
	return func() ([]dto.ImageResponseDto, error) {
		images, err := service.GetAllImage()
		if err != nil {
			return nil, err
		}
		var imageDtos []dto.ImageResponseDto
		for _, image := range images {
			imageDtos = append(imageDtos, dto.CreateImageResponseDto(image))
		}
		return imageDtos, nil
	}
}

func GetAllImageWithPagination() interface{} {
	return func(page, limit int) ([]dto.ImageResponseDto, error) {
		images, err := service.GetAllImageWithPagination((page-1)*limit, limit)
		if err != nil {
			return nil, err
		}
		var imageDtos []dto.ImageResponseDto
		for _, image := range images {
			imageDtos = append(imageDtos, dto.CreateImageResponseDto(image))
		}
		return imageDtos, nil
	}
}

func GetImageByID() interface{} {
	return func(id uint) (dto.ImageResponseDto, error) {
		image, err := service.GetImageByID(id)
		if err != nil {
			return dto.ImageResponseDto{}, err
		}
		return dto.CreateImageResponseDto(image), nil
	}
}

func InsertNewImage() interface{} {
	return func(request dto.ImageRequestDto) (dto.ImageResponseDto, error) {
		image := model.Image{
			ImageCDN: request.CDN,
		}
		err := service.InsertNewImage(image)
		if err != nil {
			return dto.ImageResponseDto{}, err
		}
		return dto.CreateImageResponseDto(image), nil
	}
}

func UpdateImageByID() interface{} {
	return func(id uint, values map[string]interface{}) error {
		err := service.UpdateImageByID(id, values)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteImageByID() interface{} {
	return func(id uint) error {
		err := service.DeleteImageByID(id)
		if err != nil {
			return err
		}
		return nil
	}
}
