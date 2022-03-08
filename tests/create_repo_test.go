package tests

import (
	"encoding/json"
	"github.com/abbul/operacion-fuego-quasar/src/domain/github_domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := github_domain.CreateRepoRequest{
		Name:        "Golang development",
		Description: "A golang development repository",
		Homepage:    "https://github_domain.com",
		Private:     true,
		HasIssues:   true,
		HasProject:  true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target github_domain.CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)
	assert.EqualValues(t, target.Homepage, request.Homepage)
	assert.EqualValues(t, target.Private, request.Private)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
	assert.EqualValues(t, target.HasProject, target.HasProject)
	assert.EqualValues(t, target.HasWiki, target.HasWiki)
}
