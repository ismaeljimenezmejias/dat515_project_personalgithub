package jsonhandler

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var testUser = User{
	ID:        1,
	Username:  "johndoe",
	Email:     "john@example.com",
	CreatedAt: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
	Active:    true,
}

var testUsers = []User{
	testUser,
	{
		ID:        2,
		Username:  "janedoe",
		Email:     "jane@example.com",
		CreatedAt: time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC),
		Active:    false,
	},
}

func testMarshalUser(t *testing.T, scoreDec func()) {
	data, err := MarshalUser(testUser)
	if err != nil {
		t.Errorf("MarshalUser() error = %v", err)
		scoreDec()
		return
	}

	if len(data) == 0 {
		t.Error("MarshalUser() returned empty data")
		scoreDec()
		return
	}

	// Test that we can unmarshal it back
	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		t.Errorf("MarshalUser() produced invalid JSON: %v", err)
		scoreDec()
		return
	}

	if !cmp.Equal(user, testUser) {
		t.Errorf("MarshalUser() roundtrip failed: got %+v, want %+v\ndiff: %s", user, testUser, cmp.Diff(testUser, user))
		scoreDec()
	}
}

func testUnmarshalUser(t *testing.T, scoreDec func()) {
	jsonData := `{"id":1,"username":"johndoe","email":"john@example.com","created_at":"2023-01-01T12:00:00Z","active":true}`

	user, err := UnmarshalUser([]byte(jsonData))
	if err != nil {
		t.Errorf("UnmarshalUser() error = %v", err)
		scoreDec()
		return
	}

	if !cmp.Equal(user, testUser) {
		t.Errorf("UnmarshalUser() = %+v, want %+v\ndiff: %s", user, testUser, cmp.Diff(testUser, user))
		scoreDec()
	}

	// Test invalid JSON
	_, err = UnmarshalUser([]byte("invalid json"))
	if err == nil {
		t.Error("UnmarshalUser() should return error for invalid JSON")
		scoreDec()
	}
}

func testMarshalUsers(t *testing.T, scoreDec func()) {
	data, err := MarshalUsers(testUsers)
	if err != nil {
		t.Errorf("MarshalUsers() error = %v", err)
		scoreDec()
		return
	}

	if len(data) == 0 {
		t.Error("MarshalUsers() returned empty data")
		scoreDec()
		return
	}

	// Test that we can unmarshal it back
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		t.Errorf("MarshalUsers() produced invalid JSON: %v", err)
		scoreDec()
		return
	}

	if !cmp.Equal(users, testUsers) {
		t.Errorf("MarshalUsers() roundtrip failed: got %+v, want %+v\ndiff: %s", users, testUsers, cmp.Diff(testUsers, users))
		scoreDec()
	}
}

func testUnmarshalUsers(t *testing.T, scoreDec func()) {
	jsonData := `[{"id":1,"username":"johndoe","email":"john@example.com","created_at":"2023-01-01T12:00:00Z","active":true},{"id":2,"username":"janedoe","email":"jane@example.com","created_at":"2023-01-02T12:00:00Z","active":false}]`

	users, err := UnmarshalUsers([]byte(jsonData))
	if err != nil {
		t.Errorf("UnmarshalUsers() error = %v", err)
		scoreDec()
		return
	}

	if !cmp.Equal(users, testUsers) {
		t.Errorf("UnmarshalUsers() = %+v, want %+v\ndiff: %s", users, testUsers, cmp.Diff(testUsers, users))
		scoreDec()
	}

	// Test invalid JSON
	_, err = UnmarshalUsers([]byte("invalid json"))
	if err == nil {
		t.Error("UnmarshalUsers() should return error for invalid JSON")
		scoreDec()
	}
}
