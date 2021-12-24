package session

import (
	"testing"

	"github.com/go-ll/llorm"
)

type User struct {
	Name string `llorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s := llorm.NewSession.Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}

// func TestSession_Model(t *testing.T) {
// 	s := NewSession().Model(&User{})
// 	table := s.RefTable()
// 	s.Model(&Session{})
// 	if table.Name != "User" || s.RefTable().Name != "Session" {
// 		t.Fatal("Failed to change model")
// 	}
// }
