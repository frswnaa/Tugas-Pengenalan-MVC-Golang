package controller

import (
	"net/http"
	"pengenalan-mvc/app/model"

	"github.com/gin-gonic/gin"
)

func AddAntrianHandler(c *gin.Context) {
	flag, err := model.AddAntrian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failed",
		})

		return
	}
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failed",
		})

		return
	}
}

func GetAntrianHandler(c *gin.Context) {
	flag, err, resp := model.GetAntrian()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})

		return
	}

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   resp,
		})

		return

	} else {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "unknown error",
		})

		return
	}
}

func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	err := model.UpdateAntrian(idAntrian)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	err := model.DeleteAntrian(idAntrian)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func PageAntrianHandler(c *gin.Context) {
	flag, err, result := model.GetAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}
