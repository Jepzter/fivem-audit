package main

import (
	"time"

	"github.com/google/uuid"
)

type Data struct {
	Online int
	Audits []Audit
}

type AuditType string

const (
	ChatMessage    AuditType = "chatMessage"
	PlayerJoining  AuditType = "playerJoining"
	PlayerDropping AuditType = "playerDropping"
)

type Audit struct {
	AuditID   uuid.UUID `json:"auditId"`
	Audit     string    `json:"audit"`
	UserID    int       `json:"userId"`
	Username  string    `json:"username"`
	AuditType AuditType `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}
