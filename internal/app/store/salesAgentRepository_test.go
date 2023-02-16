package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalesAgentRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("sales_agents")

	ticket, err := s.SalesAgentRepository().Create(&models.Sales_agent{
		Sales_agent_key:     1234,
		Flight_key:          123,
		Sold_tickets_number: 23,
	})
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
}

func TestSalesAgentRepository_FindByAgentKey(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("sales_agents")

	key := 1234

	_, err := s.SalesAgentRepository().FindByAgentKey(uint(key))
	assert.Error(t, err)

	s.SalesAgentRepository().Create(&models.Sales_agent{
		Sales_agent_key:     uint(key),
		Sold_tickets_number: 123,
		Flight_key:          12312321,
	})

	agent, err := s.SalesAgentRepository().FindByAgentKey(uint(key))

	assert.NoError(t, err)
	assert.NotNil(t, agent)
}
