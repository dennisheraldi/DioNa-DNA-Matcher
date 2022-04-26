package handler

import (
	"api/method"
	"api/penyakit"
	"api/riwayat"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type riwayatHandler struct {
	riwayatService riwayat.Service
	penyakitService penyakit.Service
}

func NewRiwayatHandler(riwayatService riwayat.Service, penyakitService penyakit.Service) *riwayatHandler {
	return &riwayatHandler{riwayatService, penyakitService}
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
	var riwayatSubmit riwayat.RiwayatSubmit
	err := c.ShouldBindJSON(&riwayatSubmit) // Mula-mula menerima submit nama pasien, dna pasien, dan nama penyakit
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

	// Mencari data Penyakit
	penyakit , _ := h.penyakitService.FindByName(riwayatSubmit.NamaPenyakit)

	// Melakukan pencocokan DNA
	status := patternFound(riwayatSubmit.DNAPasien, penyakit.DNAPenyakit)

	dateToday := time.Now().Format("2006-01-02")

	riwayatRequest := riwayat.RiwayatRequest{
		TanggalPred: dateToday,
		NamaPasien: riwayatSubmit.NamaPasien,
		DNAPasien: riwayatSubmit.DNAPasien,
		NamaPenyakit: riwayatSubmit.NamaPenyakit,
		Status: status,
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
		DNAPasien: r.DNAPasien,
		NamaPenyakit: r.NamaPenyakit,
		Status: r.Status,
	}
}

func patternFound(t, p string) string {
	if method.KMP(t, p) != -1{
		return "Positif"
	} else {
		return "Negatif"
	}
}