package database

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Generated by https://quicktype.io
type Data struct {
	Content         interface{} `json:"content"`
	Pageable        Pageable    `json:"pageable"`
	TotalPage       int64       `json:"total_page"`
	TotalElements   int64       `json:"total_elements"`
	NumberOfElement int64       `json:"number_of_element"`
	First           bool        `json:"first"`
	Last            bool        `json:"last"`
	Size            int64       `json:"size"`
	Number          int64       `json:"number"`
	Empty           bool        `json:"empty"`
	Sort            Sort        `json:"sort"`
}

type Pageable struct {
	Offset  int64 `json:"offset"`
	Paged   bool  `json:"paged"`
	Unpaged bool  `json:"unpaged"`
}

type Sort struct {
	Sorted   bool `json:"sorted"`
	Unsorted bool `json:"unsorted"`
	Empty    bool `json:"empty"`
}
type Result []interface{}

func Paginate(db *gorm.DB, c *gin.Context, out interface{}, appedQuery func(db *gorm.DB) *gorm.DB) (Data, error) {

	pagesize, _ := strconv.ParseInt(c.Query("page_size"), 0, 64)
	if pagesize == 0 {
		pagesize = 10
	}
	page, _ := strconv.ParseInt(c.Query("page"), 0, 64)

	asc := c.Query("asc") == "true"
	sortBy := c.Query("sort_by")

	if sortBy == "" {
		sortBy = "created_at"
	}
	fmt.Println("Sort", asc, sortBy)

	// if page != 0 { ?page = 1
	// 	page = page
	// }

	count := int64(0)
	s := GetFieldNames(out)
	query := db.Model(out)
	if appedQuery != nil {
		query = query.Scopes(appedQuery)
	}
	for _, field := range s {
		dbName := field
		q := c.Query("filter[" + dbName.DBName + "]")
		if q != "" {
			query = query.Where(dbName.DBName+" ilike ?", "%"+q+"%")
		}
	}

	search := c.Query("search")
	for _, field := range s {
		dbName := field
		if search != "" && (dbName.Type == "string") {
			query = query.Or(dbName.DBName+" ilike ?", "%"+search+"%")
		}
	}

	query.Count(&count)

	query = query.Scopes(PaginateOrm(int(page), int(pagesize)))
	query = query.Scopes(OrderBy(asc, sortBy))

	totalPage := math.Ceil(float64(count) / float64(pagesize))
	isFirstPage := false
	if page == 0 {
		isFirstPage = true
	}
	isLastPage := false
	if int64(totalPage) == page {
		isLastPage = true
	}
	err := error(nil)
	if count != 0 {

		err = query.Find(out).Error
	}
	isEmpty := false
	// if len(*myMap) == 0 {
	// 	isEmpty = true
	// }
	paginated := Data{
		Content: out,
		Pageable: Pageable{
			Offset:  pagesize,
			Paged:   true,
			Unpaged: false,
		},
		TotalPage:       int64(totalPage),
		TotalElements:   count,
		NumberOfElement: 0,
		First:           isFirstPage,
		Last:            isLastPage,
		Size:            pagesize,
		Number:          page,
		Empty:           isEmpty,
		Sort: Sort{
			Empty: true,
		},
	}

	return paginated, err
}

func PaginateOrm(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page) * pageSize
		fmt.Println(offset, pageSize)
		return db.Offset(offset).Limit(pageSize)
	}
}
func OrderBy(asc bool, sortBy string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if !asc {
			return db.Order(sortBy + " DESC")
		}

		return db.Order(sortBy + " ASC")
	}
}
