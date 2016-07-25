package server

import (
	"net/http"

	"github.com/drone/drone/model"
	"github.com/drone/drone/router/middleware/session"
	"github.com/drone/drone/store"

	"github.com/gin-gonic/gin"
)

func GetSecrets(c *gin.Context) {
	repo := session.Repo(c)
	secrets, err := store.GetSecretList(c, repo)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var list []*model.Secret

	for _, s := range secrets {
		list = append(list, s.Clone())
	}

	c.JSON(http.StatusOK, list)
}

func GetTeamSecrets(c *gin.Context) {
	var (
		list []*model.Secret
	)

	// TODO(must): Integrate a real implementation
	c.JSON(http.StatusOK, list)
}

func PostSecret(c *gin.Context) {
	repo := session.Repo(c)

	in := &model.Secret{}
	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON input. %s", err.Error())
		return
	}
	in.ID = 0
	in.RepoID = repo.ID

	err = store.SetSecret(c, in)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to persist secret. %s", err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func PostTeamSecret(c *gin.Context) {
	c.String(http.StatusOK, "")
	// TODO(must): Integrate a real implementation
}

func DeleteSecret(c *gin.Context) {
	repo := session.Repo(c)
	name := c.Param("secret")

	secret, err := store.GetSecret(c, repo, name)
	if err != nil {
		c.String(http.StatusNotFound, "Cannot find secret %s.", name)
		return
	}
	err = store.DeleteSecret(c, secret)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to delete secret. %s", err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteTeamSecret(c *gin.Context) {
	c.String(http.StatusOK, "")
	// TODO(must): Integrate a real implementation
}
