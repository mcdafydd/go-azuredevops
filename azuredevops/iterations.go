package azuredevops

import (
	"context"
	"fmt"
	"net/url"
)

// IterationsService handles communication with the work items methods on the API
// utilising https://docs.microsoft.com/en-gb/rest/api/vsts/work/iterations
type IterationsService struct {
	client *Client
}

// IterationsResponse describes the iterations response
type IterationsResponse struct {
	Count      int         `json:"count,omitempty"`
	Iterations []Iteration `json:"value,omitempty"`
}

// Iteration describes an iteration
type Iteration struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Path      string          `json:"path,omitempty"`
	URL       string          `json:"url,omitempty"`
	StartDate string          `json:"startDate,omitempty"`
	EndDate   string          `json:"finishDate,omitempty"`
	WorkItems [][]interface{} `json:"workItems,omitempty"`
}

// List returns list of the iterations available to the user
// utilising https://docs.microsoft.com/en-gb/rest/api/vsts/work/iterations/list
func (s *IterationsService) List(ctx context.Context, team string) ([]Iteration, error) {
	URL := fmt.Sprintf(
		"/%s/_apis/work/teamsettings/iterations?api-version=5.1-preview.1",
		url.PathEscape(team),
	)

	request, err := s.client.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	var response IterationsResponse
	_, err = s.client.Execute(ctx, request, &response)

	return response.Iterations, err
}

// GetByName will search the iterations for the account and project
// and return a single iteration if the names match
func (s *IterationsService) GetByName(ctx context.Context, team string, name string) (*Iteration, error) {
	iterations, err := s.List(ctx, team)
	if err != nil {
		return nil, err
	}

	for index := 0; index < len(iterations); index++ {
		if name == iterations[index].Name {
			iteration := iterations[index]
			return &iteration, nil
		}
	}

	return nil, nil
}
