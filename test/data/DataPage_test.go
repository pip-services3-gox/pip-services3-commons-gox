package test_data

import (
	"encoding/json"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

type user struct {
	Name string
	Age  int
}

func TestNewEmptyDataPage(t *testing.T) {
	dataPage := data.NewEmptyDataPage[user]()

	ok := dataPage.HasData()
	assert.False(t, ok)
	dt, ok := dataPage.Data()
	assert.False(t, ok)
	assert.Nil(t, dt)

	ok = dataPage.HasTotal()
	assert.False(t, ok)
	total, ok := dataPage.Total()
	assert.False(t, ok)
	assert.Equal(t, data.EmptyTotalValue, total)
}

func TestNewDataPage(t *testing.T) {
	arr := []user{{
		Name: "User1",
		Age:  26,
	}, {
		Name: "User2",
		Age:  45,
	}}
	dataPage := data.NewDataPage[user](arr, data.EmptyTotalValue)

	ok := dataPage.HasData()
	assert.True(t, ok)
	dt, ok := dataPage.Data()
	assert.True(t, ok)
	assert.Equal(t, 2, len(dt))

	ok = dataPage.HasTotal()
	assert.False(t, ok)
	total, ok := dataPage.Total()
	assert.False(t, ok)
	assert.Equal(t, data.EmptyTotalValue, total)

	buf, err := json.Marshal(dataPage)
	assert.Nil(t, err)
	assert.True(t, len(buf) > 0)

	err = json.Unmarshal(buf, &dataPage)
	assert.Nil(t, err)

	ok = dataPage.HasData()
	assert.True(t, ok)
	dt, ok = dataPage.Data()
	assert.True(t, ok)
	assert.Equal(t, 2, len(dt))
}
