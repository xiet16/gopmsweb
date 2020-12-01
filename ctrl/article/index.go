package article

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.xiet16.com/gopmsweb/models"
	"go.xiet16.com/gopmsweb/modules/response"
	"go.xiet16.com/gopmsweb/public/common"
)

func Index(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	var filters = map[string]string{"status": "", "title": "", "importance": "", "start_time": "", "end_time": ""}
	dataValues := c.QueryArray("dateValue[]")
	if len(dataValues) == 2 {
		filters["start_time"] = dataValues[0]
		filters["end_time"] = dataValues[1]
	}
	filters["status"] = c.Query("status")
	filters["importance"] = c.Query("importance")
	filters["title"] = c.Query("title")
	paging := &common.Paging{Page: page, PageSize: limit}
	articleModel := models.SystemArticle{}
	articleArr, err := articleModel.GetAllPage(paging, filters)
	var articlePageArr []models.SystemArticlePage

	for _, v := range articleArr {
		userModel := models.SystemUser{}
		userModel.Id = v.Author
		has := userModel.GetRow()
		if !has {
			continue
		}
		articlePageArr = append(articlePageArr, models.SystemArticlePage{SystemArticle: v, AuthorName: userModel.Name})
		if err != nil {
			response.ShowError(c, "fail")
			return
		}
		data := make(map[string]interface{})
		data["items"] = articlePageArr
		data["total"] = paging.Total
		response.ShowData(c, data)
		return
	}
}
