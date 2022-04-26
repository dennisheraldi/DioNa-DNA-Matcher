package handler

import (
	"api/riwayat"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type riwayatHandler struct {
	riwayatService riwayat.Service
}

func NewRiwayatHandler(riwayatService riwayat.Service) *riwayatHandler {
	return &riwayatHandler{riwayatService}
}

func (h *riwayatHandler) GetAllRiwayatHandler(c *gin.Context){
	riwayats, err := h.riwayatService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var riwayatResponses []riwayat.RiwayatResponse

	for _, r := range riwayats {
		riwayatResponse := convertToRiwayatResponse(r)
		riwayatResponses = append(riwayatResponses, riwayatResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": riwayatResponses,
	})
}

func (h *riwayatHandler) CreateRiwayatHandler(c *gin.Context){
	var riwayatRequest riwayat.RiwayatRequest

	err := c.ShouldBindJSON(&riwayatRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s:, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	riwayat, err := h.riwayatService.Create(riwayatRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	riwayatResponse := convertToRiwayatResponse(riwayat)

	c.JSON(http.StatusOK, gin.H{
		"data": riwayatResponse,
	})
}

func convertToRiwayatResponse(r riwayat.Riwayat) riwayat.RiwayatResponse {
	tanggalPred := r.TanggalPred.Format("2006-01-02")
	
	return riwayat.RiwayatResponse{
		TanggalPred: tanggalPred,
		NamaPasien: r.NamaPasien,
		NamaPenyakit: r.NamaPenyakit,
		Status: r.Status,
	}
}