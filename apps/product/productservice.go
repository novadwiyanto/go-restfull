package product

import (
	"go-restapi/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var whereId string = "id = ?"

func GetAll(c *gin.Context, product []Product) (database.Data, error) {
	data, err := database.Paginate(database.DB, c, &product, func(db *gorm.DB) *gorm.DB {
		return db
	})

	return data, err
}

func GetDetail(c *gin.Context) (Product, error) {
	data := Product{}
	id := c.Param("id")

	err := database.DB.Model(&Product{}).Where(whereId, id).First(&data).Error

	return data, err
}

func DoStore(c *gin.Context, product Product) (Product, error) {
	data := Product{}

	err := database.DB.Model(&Product{}).Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, err
}

func DoUpdate(c *gin.Context, product Product) (Product, error) {
	data := Product{}

	id := c.Param("id")
	err := database.DB.Model(&Product{}).Where(whereId, id).First(&data).Error
	if err != nil {
		return data, err
	}

	errUpdate := database.DB.Model(&data).Updates(product).Error
	if errUpdate != nil {
		return data, errUpdate
	}

	return data, err
}

func DoDelete(c *gin.Context) error {
	id := c.Param("id")
	data := Product{}
	err := database.DB.Model(&Product{}).Where(whereId, id).First(&data).Error
	if err != nil {
		return err
	}
	errDelete := database.DB.Delete(&data).Error
	if errDelete != nil {
		return errDelete
	}
	return err
}
