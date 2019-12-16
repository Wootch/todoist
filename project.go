package todoist

import (
	"context"
	"encoding/json"

	"github.com/ides15/todoist/types"
)

type ProjectService struct {
	c *Client
}

func (s *ProjectService) GetProjects() ([]*types.Project, error) {
	s.c.Log("GetProjects called")
	req, err := s.c.NewRequest("*", nil, &[]string{"projects"})
	if err != nil {
		return nil, err
	}

	res, err := s.c.Do(context.Background(), req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &types.Response{}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Projects, nil
}

func (s *ProjectService) GetProjectByID(id int) (*types.Project, error) {
	s.c.Log("GetProjectByID called")
	projects, err := s.GetProjects()
	if err != nil {
		return nil, err
	}

	s.c.Log("projects", projects)

	for _, project := range projects {
		if project.ID == id {
			return project, nil
		}

		return nil, types.ErrNotFound
	}

	return nil, types.ErrNotFound
}

func (s *ProjectService) GetProjectByName(name string) (*types.Project, error) {
	s.c.Log("GetProjectByName called")
	projects, err := s.GetProjects()
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		if project.Name == name {
			return project, nil
		}

		return nil, types.ErrNotFound
	}

	return nil, types.ErrNotFound
}
