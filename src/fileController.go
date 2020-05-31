package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UploadFile struct {
}
//
//func (fileUploader UploadFile) UploadSingleFile(ctx *gin.Context) {
//	request := ctx.Request
//	file, handler, err := request.FormFile("ex1")
//	if err != nil {
//		log.Error("Error Retrieving the File")
//		log.Error(err)
//		ctx.JSON(http.StatusInternalServerError, gin.H{})
//	}
//	f, err := os.OpenFile("uploadData/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{})
//	}
//	defer f.Close()
//	io.Copy(f, file)
//
//	log.WithFields(log.Fields{
//		"FileName": handler.Filename,
//		"Filesize": handler.Size,
//	}).Info("File Uploaded Successfully")
//	ctx.JSON(200, gin.H{"home": "sweet home"})
//}

func (fileUploader UploadFile) Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Status": "Working"})
}


func NewFileUploader() UploadFile {
	return UploadFile{}
}

