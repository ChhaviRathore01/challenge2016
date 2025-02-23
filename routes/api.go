package routes

import (
	"challenge2016/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckDistribution API Handler
func CheckDistribution(c *gin.Context) {
	distributorName := c.Query("distributor")
	city := c.Query("city")
	state := c.Query("state")
	country := c.Query("country")

	distributor, exists := services.Distributors[distributorName]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Distributor not found"})
		return
	}

	result := distributor.CanDistribute(city, state, country)
	c.JSON(http.StatusOK, gin.H{"distributor": distributorName, "canDistribute": result})
}

// AddDistributor API Handler
func AddDistributor(c *gin.Context) {
	name := c.Query("name")
	parent := c.Query("parent")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Distributor name is required"})
		return
	}

	distributor, err := services.CreateDistributor(name, parent)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) // 409 Conflict for duplicate distributor
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Distributor created successfully",
		"distributor": distributor.Name,
	})
}

// AddPermission API Handler
func AddPermission(c *gin.Context) {
	distributorName := c.Query("distributor")
	permissionType := c.Query("type")
	region := c.Query("region")

	distributor, exists := services.Distributors[distributorName]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Distributor not found"})
		return
	}

	err := distributor.AddPermission(permissionType, region)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permission added successfully"})
}
