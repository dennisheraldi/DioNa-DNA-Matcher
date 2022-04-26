package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"api/penyakit"
)

type penyakitHandler struct {
	penyakitService penyakit.Service
}

func NewPenyakitHandler(penyakitService penyakit.Service) *penyakitHandler {
	return &penyakitHandler{ penyakitService }
}

func (h *penyakitHandler) GetAllPenyakitHandler(c *gin.Context) {
	penyakits, err := h.penyakitService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var penyakitResponses []penyakit.PenyakitResponse

	for _, p := range penyakits {
		penyakitResponse := convertToPenyakitResponse(p)
		penyakitResponses = append(penyakitResponses, penyakitResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": penyakitResponses,
	})
}

func (h *penyakitHandler) GetPenyakitHandler(c *gin.Context){
	getid := c.Param("id")
	id, _ := strconv.Atoi(getid)

	p, err := h.penyakitService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	penyakitResponse := convertToPenyakitResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"data": penyakitResponse,
	})
}

func (h *penyakitHandler) CreatePenyakitHandler(c *gin.Context) {
	var penyakitRequest penyakit.PenyakitRequest

	err := c.ShouldBindJSON(&penyakitRequest)
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

	penyakit, err := h.penyakitService.FindByName(penyakitRequest.NamaPenyakit)

	if len(penyakit.NamaPenyakit) != 0 {
		penyakitResponse := convertToPenyakitResponse(penyakit)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Penyakit already exist",
			"data": penyakitResponse,
		})
		return
	}

	penyakit, err = h.penyakitService.Create(penyakitRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	penyakitResponse := convertToPenyakitResponse(penyakit)

	c.JSON(http.StatusOK, gin.H{
		"data": penyakitResponse,
	})
}

func (h *penyakitHandler) UpdatePenyakitHandler(c *gin.Context) {
	var penyakitRequest penyakit.PenyakitRequest

	err := c.ShouldBindJSON(&penyakitRequest)
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

	getID := c.Param("id")
	id, _ := strconv.Atoi(getID)

	penyakit, err := h.penyakitService.Update(id, penyakitRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": penyakit,
	})

}

func (h *penyakitHandler) DeletePenyakitHandler(c *gin.Context) {
	getID := c.Param("id")
	id, _ := strconv.Atoi(getID)

	p, err := h.penyakitService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	penyakitResponse := convertToPenyakitResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"data": penyakitResponse,
		"status": "deleted",
	})
}

func convertToPenyakitResponse(p penyakit.Penyakit) penyakit.PenyakitResponse {
	return penyakit.PenyakitResponse{
		NamaPenyakit:      p.NamaPenyakit,
		DNAPenyakit: p.DNAPenyakit,
	}
}
