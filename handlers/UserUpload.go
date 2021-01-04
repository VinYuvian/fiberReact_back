package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

func UserUpload(ctx *fiber.Ctx) error{
	//file,err:=ctx.FormFile("file") - file upload for not a multi-part
	name:=ctx.FormValue("name")
	paramId := ctx.Params("Id")
	id,err := strconv.Atoi(paramId)
	formData,err:=ctx.MultipartForm()
	if err!=nil{
		return err
	}
	files:=formData.File["file"]
	path:=fmt.Sprintf("uploads/%s",name)
	os.MkdirAll(path,0755)
	for i:=0;i<len(files);i++{
		ctx.SaveFile(files[i],fmt.Sprintf("%s/%s",path,files[i].Filename))
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"succeed","id":id,
	})

	/*ctx.SaveFile(file,fmt.Sprintf("%s/%s",path,file.Filename))
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message":"image upload success",
		"username":name,"id":formData.File,
	})*/
}
