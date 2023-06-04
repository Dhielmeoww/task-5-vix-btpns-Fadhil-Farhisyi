package controllers

import (
	"task-5-vix-btpns-Fadhil-Farhisyi/initializers"
	"task-5-vix-btpns-Fadhil-Farhisyi/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//getdata
	var getdata struct {
		Title    string
		Caption  string
		PhotoUrl string
		USC      int32
	}

	c.Bind(&getdata)

	//post
	post := models.Post{Title: getdata.Title, Caption: getdata.Caption, PhotoUrl: getdata.PhotoUrl, USC: getdata.USC}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get post
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})

	//res
}

func PostsOne(c *gin.Context) {
	//get id
	id := c.Param("id")

	//get post
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsUpdate(c *gin.Context) {
	//get id
	id := c.Param("photoid")

	//get data
	var getdata struct {
		Title    string
		Caption  string
		PhotoUrl string
		USC      int32
	}

	c.Bind(&getdata)

	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: getdata.Title, Caption: getdata.Caption, PhotoUrl: getdata.PhotoUrl, USC: getdata.USC})

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	//get id
	id := c.Param("photoid")

	//delete data
	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
