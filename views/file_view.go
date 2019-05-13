package views

import (
	"github.com/Mindyu/blog_system/utils"
	"github.com/Mindyu/blog_system/utils/fileutil"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"time"
)


func Upload(ctx *gin.Context) {
	t := time.Now()
	dir := t.Format("20060102")
	info, err := ctx.FormFile("file")
	if err != nil {
		utils.MakeErrResponse(ctx, err.Error())
		return
	}

	file, err := info.Open()
	ext := fileutil.ExtensionName(info.Filename)
	path := dir + "/"
	err = os.MkdirAll("./public/upload/"+path, os.ModePerm)
	if err != nil {
		utils.MakeErrResponse(ctx, err.Error())
		return
	}

	path += fileutil.UniqueID() + strings.ToLower(ext)
	out, err := os.OpenFile("./public/upload/"+path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.MakeErrResponse(ctx, err.Error())
		return
	}

	defer out.Close()
	io.Copy(out, file)
	utils.MakeOkResponse(ctx, path)
}

func Download(ctx *gin.Context) *os.File {
	filePath := ctx.Param("file")

	out, err := os.OpenFile("./public/upload/"+filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.MakeErrResponse(ctx, err.Error())
		return nil
	}
	file := &os.File{}
	defer out.Close()
	io.Copy(file, out)

	return file
}