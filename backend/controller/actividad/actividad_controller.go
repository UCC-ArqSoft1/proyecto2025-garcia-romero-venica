package actividad

import (
	"backend/dto"
	service"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// GET /actividades/
func  GetAll(ctx *gin.Context) {
	actividades, err := service.ActividadService.GetAllActividades()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, actividades)
}

// GET /actividades/:id
func GetActividadByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	actividad, err := service.ActividadService.GetActividadByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, actividad)
}

// POST /actividades/
func Create(ctx *gin.Context) {
	var actividadDto dto.ActividadDto
	if err := ctx.ShouldBindJSON(&actividadDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := service.ActividadService.CreateActividad(actividadDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

// PUT /actividades/:id
func Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var actividadDto dto.ActividadDto
	if err := ctx.ShouldBindJSON(&actividadDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := service.ActividadService.UpdateActividad(id, actividadDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

// DELETE /actividades/:id
func Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := service.ActividadService.DeleteActividad(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
