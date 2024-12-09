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

	// "bytes"
	"io"
	"image"
	"image/jpeg"
	_ "image/png" // Untuk mendukung file PNG

	"github.com/disintegration/imaging"

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
	e.PUT("/question_options/:id", handler.UpdateWithAnswerOptions)
	e.GET("/question/:id", handler.GetByIDHandler)
	e.GET("/question/test_id/:id", handler.GetByTestIDHandler)
	e.DELETE("/question/:id", handler.DeleteHandler)
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

func (h *QuestionHandler) UpdateWithAnswerOptions(c echo.Context) error {
	id := c.Param("id")

	var data domain.Question

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.UpdateWithAnswerOptions(id, &data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	return helper_http.SuccessResponse(c, data, "success update question with answer_options")
}



func (h *QuestionHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	err := h.usecase.Delete(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, nil, "success delete option")
	return resp
}

type FileUpload struct {
	File string `json:"file"`
}
	// // UploadFile handles the file upload
	// func (h *QuestionHandler) UploadFile(c echo.Context) error {
	// 	// Mendapatkan file dari form data
	// 	file, err := c.FormFile("file")
	// 	if err != nil {
	// 		return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to get file: %v", err))
	// 	}
	
	// 	// Tentukan lokasi penyimpanan file
	// 	uploadDir := "../gapaiskor_fe/build/uploads"
	// 	// uploadDir := "../uploads"
	// 	// Tentukan lokasi penyimpanan file
	// 	// uploadDir := "/home/gapaisko/gapaiskorweb/gapaiskor_fe/build/uploads"
	
	// 	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
	// 		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to create upload directory: %v", err))
	// 	}
	
	// 	// Menambahkan timestamp atau UUID untuk menghindari penimpaan file yang sama
	// 	uniqueFilename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	// 	dst := filepath.Join(uploadDir, uniqueFilename)
	
	// 	// Membuka file tujuan untuk menulis
	// 	src, err := file.Open()
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to open file: %v", err))
	// 	}
	// 	defer src.Close()
	
	// 	// Membuat file di tujuan dan menulis file
	// 	dstFile, err := os.Create(dst)
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to create file: %v", err))
	// 	}
	// 	defer dstFile.Close()
	
	// 	// Menyalin file ke lokasi tujuan
	// 	_, err = dstFile.ReadFrom(src)
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to save file: %v", err))
	// 	}
	
	// 	dst = "https://gapaiskor.web.id/uploads/" + uniqueFilename
	// 	// Mengembalikan response sukses dengan lokasi file yang di-upload
	// 	// return c.JSON(http.StatusOK, helper_http.SuccessResponse(c, FileUpload{File: dst}, "file uploaded successfully"))
	// 	return helper_http.SuccessResponse(c, FileUpload{File: dst}, "success file upload")
	
	// }

	
	func (h *QuestionHandler) UploadFile(c echo.Context) error {
		// Mendapatkan file dari form data
		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Failed to get file: %v", err))
		}
	
		// Membuka file sumber untuk membaca
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to open file: %v", err))
		}
		defer src.Close()
	
		// Membaca data file untuk identifikasi jenis file
		buffer := make([]byte, 512)
		if _, err := src.Read(buffer); err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to read file")
		}
		contentType := http.DetectContentType(buffer)
	
		// Reset posisi pembacaan file
		src.Seek(0, 0)
	
		// Tentukan lokasi penyimpanan
		uploadDir := "../gapaiskor_fe/build/uploads"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to create upload directory: %v", err))
		}
	
		uniqueFilename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		dst := filepath.Join(uploadDir, uniqueFilename)
	
		// Kompresi gambar jika file adalah gambar
		if contentType == "image/jpeg" || contentType == "image/png" {
			img, _, err := image.Decode(src)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to decode image: %v", err))
			}
	
			// Resize gambar menggunakan library Imaging
			resizedImg := imaging.Resize(img, 800, 0, imaging.Lanczos) // Resize width menjadi 800px (proportional)
			dstFile, err := os.Create(dst)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to create file: %v", err))
			}
			defer dstFile.Close()
	
			// Simpan gambar dengan kualitas 80 (lebih kecil ukurannya)
			err = jpeg.Encode(dstFile, resizedImg, &jpeg.Options{Quality: 80})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to save compressed image: %v", err))
			}
		} else {
			// Untuk file non-gambar, salin langsung
			dstFile, err := os.Create(dst)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to create file: %v", err))
			}
			defer dstFile.Close()
	
			if _, err := io.Copy(dstFile, src); err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to save file: %v", err))
			}
		}
	
		// URL publik untuk file yang di-upload
		publicURL := "https://gapaiskor.web.id/uploads/" + uniqueFilename
	
		// Mengembalikan response sukses dengan lokasi file yang di-upload
		return helper_http.SuccessResponse(c, FileUpload{File: publicURL}, "File uploaded successfully")
	}
	