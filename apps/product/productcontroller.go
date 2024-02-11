package product

import (
	"go-restapi/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data, err := GetAll(c, []Product{})
	helper.SendData(c, data, err)
	return
}

func Show(c *gin.Context) {
	data, err := GetDetail(c)

	request := DetailResponse{}
	request.NamaProduct = data.NamaProduct
	request.Deskripsi = data.Deskripsi

	helper.SendData(c, request, err)
	return
}

func Create(c *gin.Context) {
	product := Product{}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data, err := DoStore(c, product)

	helper.SendData(c, data, err)
	return
}

func Update(c *gin.Context) {
	product := Product{}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data, err := DoUpdate(c, product)
	helper.SendData(c, data, err)
	return
}

func Delete(c *gin.Context) {
	err := DoDelete(c)
	helper.SendStatus(c, err)
	return
}
