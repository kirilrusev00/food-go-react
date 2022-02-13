package mocks

import (
	"database/sql"
)

// MockDB is the mock database
type MockDB struct {
	callParams []interface{}
}

// Query implements the MockDB interface
func (mockDB *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	mockDB.callParams = []interface{}{query}
	mockDB.callParams = append(mockDB.callParams, args...)

	return nil, nil
}

// CalledWith is a helper method to inspect the `callParams` field
func (mockDB *MockDB) CalledWith() []interface{} {
	return mockDB.callParams
}
