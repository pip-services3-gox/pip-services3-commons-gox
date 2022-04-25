package test_commands

import (
	"context"
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	"github.com/pip-services3-gox/pip-services3-commons-gox/run"
	"github.com/stretchr/testify/assert"
)

func commandExec(ctx context.Context, correlationId string, args *run.Parameters) (any, error) {
	if correlationId == "wrongId" {
		panic("Test error")
	}

	return nil, nil
}

func TestGetCommandName(t *testing.T) {
	command := commands.NewCommand("name", nil, commandExec)

	// Check match by individual fields
	assert.NotNil(t, command)
	assert.Equal(t, "name", command.Name())
}

func TestExecuteCommand(t *testing.T) {
	command := commands.NewCommand("name", nil, commandExec)

	_, err := command.Execute(context.Background(), "", nil)
	assert.Nil(t, err)

	_, err = command.Execute(context.Background(), "wrongId", nil)
	assert.NotNil(t, err)
}
