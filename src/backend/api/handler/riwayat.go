package handler

import (
	"api/library"
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
			"status_code": http.StatusBadRequest,
		})
		return
	}

	// Mencari data Penyakit
	penyakit , _ := h.penyakitService.FindByName(riwayatSubmit.NamaPenyakit)

	if penyakit.NamaPenyakit == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Penyakit tidak ditemukan",
			"status_code" : http.StatusBadRequest,
		})
		return
	}

	// Mengecek validasi DNA Pasien
	isDNAValid := library.Sanitasi(riwayatSubmit.DNAPasien)

	if !isDNAValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DNA Pasien tidak valid",
			"status_code": http.StatusBadRequest,
		})
		return
	}

	// Melakukan pencocokan DNA
	status, similarity := library.CheckDNA(riwayatSubmit.DNAPasien, penyakit.DNAPenyakit)

	dateToday := time.Now().Format("2006-01-02")

	riwayatRequest := riwayat.RiwayatRequest{
		TanggalPred: dateToday,
		NamaPasien: riwayatSubmit.NamaPasien,
		NamaPenyakit: riwayatSubmit.NamaPenyakit,
		Similarity: similarity,
		Status: status,
	}

	riwayat, err := h.riwayatService.Create(riwayatRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			"status_code" : http.StatusBadRequest,
		})
		return
	}

	riwayatResponse := convertToRiwayatResponse(riwayat)

	c.JSON(http.StatusOK, gin.H{
		"data": riwayatResponse,
		"status_code" : http.StatusOK,
	})
}

func (h *riwayatHandler) QueryRiwayatHandler(c *gin.Context){
	query := c.Query("query")

	// parsing query
	tanggal, namaPenyakit := library.QueryCheck(query)

	riwayats, err := h.riwayatService.FindAll()

	if tanggal == "" && namaPenyakit == "" { 
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query kosong",
			"status_code" : http.StatusBadRequest,
		})
		return
	} else if tanggal == "all" && namaPenyakit == "all" { // kalau tidak ada tanggal maupun nama penyakit
		riwayats, err = h.riwayatService.FindAll() 
	} else if tanggal != "" && namaPenyakit == "" { // pencarian berdasarkan tanggal
		riwayats, err = h.riwayatService.FindByTanggal(tanggal)
	} else if tanggal == "" && namaPenyakit != "" { // pencarian berdasarkan nama penyakit
		riwayats, err = h.riwayatService.FindByPenyakit(namaPenyakit)
	} else { // pencarian berdasarkan tanggal dan nama penyakit
		riwayats, err = h.riwayatService.FindByTanggalPenyakit(tanggal, namaPenyakit)
	}

	// jika query tidak valid, maka akan menampilkan error
	if len(riwayats) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
			"status_code" : http.StatusBadRequest,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			"status_code" : http.StatusBadRequest,
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
		"status_code" : http.StatusOK,
	})
}

func convertToRiwayatResponse(r riwayat.Riwayat) riwayat.RiwayatResponse {
	tanggalPred := r.TanggalPred.Format("2006-01-02")
	
	return riwayat.RiwayatResponse{
		TanggalPred: tanggalPred,
		NamaPasien: r.NamaPasien,
		NamaPenyakit: r.NamaPenyakit,
		Similarity: r.Similarity,
		Status: r.Status,
	}
}

