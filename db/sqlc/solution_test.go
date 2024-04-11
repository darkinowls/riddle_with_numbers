package db

import (
	"context"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSomeSolution(t *testing.T, message *string) Solution {

	solution, err := testQueries.CreateSolution(context.Background(), CreateSolutionParams{
		Condition: json.RawMessage(*message),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, solution)

	return solution
}

func TestCreateSolution(t *testing.T) {
	message := `{"message": "Hello, world!"}`
	solution := createSomeSolution(t, &message)

	count, err := testQueries.GetSolutionCount(context.Background())
	assert.NoError(t, err)
	assert.True(t, count >= 1)

	solution, err = testQueries.GetSolution(context.Background(), solution.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, solution)
	assert.Equal(t, message, string(solution.Condition))

}

func TestDeleteSolutions(t *testing.T) {
	message := `{"message": "Hello, world!"}`
	solution := createSomeSolution(t, &message)

	count, err := testQueries.GetSolutionCount(context.Background())
	assert.NoError(t, err)
	assert.True(t, count >= 1)

	err = testQueries.DeleteAllSolutions(context.Background())
	assert.NoError(t, err)
	solution, err = testQueries.GetSolution(context.Background(), solution.ID)
	assert.Error(t, err)
	assert.Empty(t, solution)

	count, err = testQueries.GetSolutionCount(context.Background())
	assert.NoError(t, err)
	assert.True(t, count == 0)

	solution = createSomeSolution(t, &message)
	assert.True(t, solution.ID == 1)
}
