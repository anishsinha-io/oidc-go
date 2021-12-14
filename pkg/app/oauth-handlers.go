package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/anish-sinha1/oidc-go/pkg/api/services/oauth"
	"github.com/anish-sinha1/oidc-go/pkg/repo/models"
	"github.com/gin-gonic/gin"
)

func handleShowUniversalLogin(c *gin.Context) {
	var ar models.AuthorizeQs
	if err := c.Bind(&ar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid_request")})
		return
	}
	valid, err := oauth.ValidateUriParams(ar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid_request")})
		return
	}
	fmt.Println(valid)
	fmt.Println(ar)
}

func handleGetToken(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func handleTokenIntrospect(c *gin.Context) {
	var token models.IntrospectionRequest
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := oauth.ValidateToken(token.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}

func handleLogin(c *gin.Context) {
	var lr models.LoginRequest
	if err := c.BindJSON(&lr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// db login service
	c.JSON(http.StatusOK, "")
}

func handleSignup(c *gin.Context) {
	var sr models.SignupRequest
	if err := c.BindJSON(&sr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// db signup service
	c.JSON(http.StatusOK, "")
}
