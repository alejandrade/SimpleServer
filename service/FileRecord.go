package service

type FileRecord struct {
	FileId       string
	User         string
	FileName     string
	FileSize     int64
	FileLocation string
}

func (file *FileRecord) setFileId(FileId string) {
	file.FileId = FileId
}

func (file *FileRecord) setUser(user string) {
	file.User = user
}

func (file *FileRecord) setFileName(fileName string) {
	file.FileName = fileName
}

func (file *FileRecord) setFileSize(fileSize int64) {
	file.FileSize = fileSize
}

func (file *FileRecord) setFileLocation(fileLocation string) {
	file.FileLocation = fileLocation
}
