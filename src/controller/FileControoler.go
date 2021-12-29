package controller

import (
	"durin/src/model"
	"durin/src/util"
	"log"

	"github.com/gin-gonic/gin"
)

func File(c *gin.Context) {
	// 获取文件头
	file, err := c.FormFile("upload")
	if err != nil {
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		log.Println(err)
	}

	filePath := util.UploadImg(openedFile)

	fileDTO := model.FileDTO{
		Path: filePath,
	}
	c.JSON(200, util.NewReturnObject(0, "Success", fileDTO))
}
