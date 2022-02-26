package service

type DBServiceInterface interface {
	FileServiceInterface
	GroupServiceInterface
	ImageServiceInterface
	MessageFileServiceInterface
	MessageImageServiceInterface
	MessageServiceInterface
	UserServiceInterface
	RoomServiceInterface
	UserServiceInterface
}
