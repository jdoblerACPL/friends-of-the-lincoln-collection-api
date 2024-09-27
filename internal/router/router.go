package router

import (
	"net/http"

	db "acpl.lib.in.us/m/internal"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the gin router
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/news", getNews)
	return r
}

type News struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

func getNews(c *gin.Context) {
	rows, err := db.DB.Query("SELECT * FROM news")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var NewsArr []News
	for rows.Next() {
		var news News
		err = rows.Scan(&news.ID, &news.Title, &news.Content, &news.Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		NewsArr = append(NewsArr, news)
	}
	c.JSON(http.StatusOK, gin.H{
		"news": NewsArr,
	})
}
