package handler

import (
	// "fmt"

	"fmt"
	"gapai-skor-api/domain"
	helper_http "gapai-skor-api/transport/http/helper"
	"net/http"
	"os"
	"path/filepath"
	"time"

	// "strconv"

	"github.com/labstack/echo"
)

type QuestionHandler struct {
	usecase domain.QuestionUsecase
}

func QuestionRoute(e *echo.Echo, uc domain.QuestionUsecase) {
	handler := QuestionHandler{
		usecase: uc,
	}
	e.GET("/question/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/question")
	})

	e.GET("/question", handler.GetAllHandler)
	e.POST("/question", handler.Create)
	e.POST("/question_options", handler.CreateWithAnswerOptions)
	e.GET("/question/:id", handler.GetByIDHandler)
	e.GET("/question/test_id/:id", handler.GetByTestIDHandler)
	e.POST("/upload", handler.UploadFile)

}

func (h *QuestionHandler) GetAllHandler(c echo.Context) error {
	// init handler
	data, err := h.usecase.GetAllData()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all question")

	return resp
}

func (h *QuestionHandler) GetByIDHandler(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetByID(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

func (h *QuestionHandler) GetByTestIDHandler(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetByTestID(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

func (h *QuestionHandler) Create(c echo.Context) error {
	var data domain.Question

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create question")
}

func (h *QuestionHandler) CreateWithAnswerOptions(c echo.Context) error {
	var data domain.Question

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.CreateWithAnswerOptions(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	return helper_http.SuccessResponse(c, data, "success create question with answer_options")
}

type FileUpload struct {
	File string `json:"file"`
}

// UploadFile handles the file upload
func (h *QuestionHandler) UploadFile(c echo.Context) error {
	// Mendapatkan file dari form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to get file: %v", err))
	}

	// Tentukan lokasi penyimpanan file
	uploadDir := "../gapaiskor_fe/build/uploads"
	// Tentukan lokasi penyimpanan file
	// uploadDir := "/home/gapaisko/gapaiskorweb/gapaiskor_fe/build/uploads"

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to create upload directory: %v", err))
	}

	// Menambahkan timestamp atau UUID untuk menghindari penimpaan file yang sama
	uniqueFilename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	dst := filepath.Join(uploadDir, uniqueFilename)

	// Membuka file tujuan untuk menulis
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to open file: %v", err))
	}
	defer src.Close()

	// Membuat file di tujuan dan menulis file
	dstFile, err := os.Create(dst)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to create file: %v", err))
	}
	defer dstFile.Close()

	// Menyalin file ke lokasi tujuan
	_, err = dstFile.ReadFrom(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to save file: %v", err))
	}

	dst = "https://gapaiskor.web.id/uploads/" + uniqueFilename
	// Mengembalikan response sukses dengan lokasi file yang di-upload
	// return c.JSON(http.StatusOK, helper_http.SuccessResponse(c, FileUpload{File: dst}, "file uploaded successfully"))
	return helper_http.SuccessResponse(c, FileUpload{File: dst}, "success file upload")

}
