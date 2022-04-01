package controllers

import (
	"errors"
	"fmt"
	"practice/cache"
	"practice/database"
	"practice/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
)

type MemberRepo struct {
	Db *gorm.DB
}

func New() *MemberRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Member{})
	return &MemberRepo{Db: db}
}

//get member by id
func (repository *MemberRepo) GetMember(c *gin.Context) {
	getCount, err := cache.Client.Incr("Get").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Count of %s operations = %d\n", "GET", getCount)

	id, _ := c.Params.Get("id")
	var member models.Member
	err = models.GetMember(repository.Db, &member, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Member doesn't exist"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

//get members
func (repository *MemberRepo) GetMembers(c *gin.Context) {
	getAllCount, err := cache.Client.Incr("GetAll").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Count of %s operations = %d\n", "GETALL", getAllCount)

	var member []models.Member
	err = models.GetMembers(repository.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

//create member
func (repository *MemberRepo) CreateMember(c *gin.Context) {
	createCount, err := cache.Client.Incr("Create").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Count of %s operations = %d\n", "CREATE", createCount)

	var member models.Member
	c.BindJSON(&member)
	err = models.CreateMember(repository.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

// update member
func (repository *MemberRepo) UpdateMember(c *gin.Context) {
	updateCount, err := cache.Client.Incr("Update").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Count of %s operations = %d\n", "UPDATE", updateCount)

	var member models.Member
	id, _ := c.Params.Get("id")
	err = models.GetMember(repository.Db, &member, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Member doesn't exist"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&member)
	err = models.UpdateMember(repository.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

// delete member
func (repository *MemberRepo) DeleteMember(c *gin.Context) {
	deleteCount, err := cache.Client.Incr("Delete").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Count of %s operations = %d\n", "DELETE", deleteCount)

	var member models.Member
	id, _ := c.Params.Get("id")
	err = models.GetMember(repository.Db, &member, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Member doesn't exist"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.DeleteMember(repository.Db, &member, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully"})
}
