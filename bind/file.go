package bind

import (
	"github.com/FakeqProject/fakeq-app/dto"
	"github.com/FakeqProject/fakeq-app/model"
	"github.com/FakeqProject/fakeq-app/service"
	"github.com/webview/webview"
)

func FileBindCollection(w webview.WebView) {
	w.Bind("getAllFile", GetAllFile())
	w.Bind("getAllFileWithPagination", GetAllFileWithPagination())
	w.Bind("getFileByID", GetFileByID())
	w.Bind("insertNewFile", InsertNewFile())
	w.Bind("updateFileByID", UpdateFileByID())
	w.Bind("deleteFileByID", DeleteFileByID())
}

func GetAllFile() interface{} {
	return func() ([]dto.FileResponseDto, error) {
		files, err := service.GetAllFile()
		if err != nil {
			return nil, err
		}
		var fileDtos []dto.FileResponseDto
		for _, file := range files {
			fileDtos = append(fileDtos, dto.CreateFileResponseDto(file))
		}
		return fileDtos, nil
	}
}

func GetAllFileWithPagination() interface{} {
	return func(page, limit int) ([]dto.FileResponseDto, error) {
		files, err := service.GetAllFileWithPagination((page-1)*limit, limit)
		if err != nil {
			return nil, err
		}
		var fileDtos []dto.FileResponseDto
		for _, file := range files {
			fileDtos = append(fileDtos, dto.CreateFileResponseDto(file))
		}
		return fileDtos, nil
	}
}

func GetFileByID() interface{} {
	return func(id uint) (dto.FileResponseDto, error) {
		file, err := service.GetFileByID(id)
		if err != nil {
			return dto.FileResponseDto{}, err
		}
		return dto.CreateFileResponseDto(file), nil
	}
}

func InsertNewFile() interface{} {
	return func(request dto.FileRequestDto) (dto.FileResponseDto, error) {
		file := model.File{
			FileCDN: request.CDN,
		}
		err := service.InsertNewFile(file)
		if err != nil {
			return dto.FileResponseDto{}, err
		}
		return dto.CreateFileResponseDto(file), nil
	}
}

func UpdateFileByID() interface{} {
	return func(id uint, values map[string]interface{}) error {
		err := service.UpdateFileByID(id, values)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteFileByID() interface{} {
	return func(id uint) error {
		err := service.DeleteFileByID(id)
		if err != nil {
			return err
		}
		return nil
	}
}
