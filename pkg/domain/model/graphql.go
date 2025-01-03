// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/secmon-lab/alertchain/pkg/domain/types"
)

type ActionRecord struct {
	ID         string            `json:"id"`
	Seq        int               `json:"seq"`
	Uses       string            `json:"uses"`
	Args       []*ArgumentRecord `json:"args"`
	Result     *string           `json:"result,omitempty"`
	Next       []*NextRecord     `json:"next"`
	Error      *string           `json:"error,omitempty"`
	StartedAt  time.Time         `json:"startedAt"`
	FinishedAt time.Time         `json:"finishedAt"`
}

type AlertRecord struct {
	ID          types.AlertID      `json:"id"`
	Schema      string             `json:"schema"`
	Data        string             `json:"data"`
	CreatedAt   time.Time          `json:"createdAt"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Source      string             `json:"source"`
	Namespace   *string            `json:"namespace,omitempty"`
	InitAttrs   []*AttributeRecord `json:"initAttrs"`
	LastAttrs   []*AttributeRecord `json:"lastAttrs"`
	Refs        []*ReferenceRecord `json:"refs"`
}

type ArgumentRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AttributeRecord struct {
	ID      string  `json:"id"`
	Key     string  `json:"key"`
	Value   string  `json:"value"`
	Type    *string `json:"type,omitempty"`
	Persist bool    `json:"persist"`
	TTL     int     `json:"ttl"`
}

type NextRecord struct {
	Abort bool               `json:"abort"`
	Attrs []*AttributeRecord `json:"attrs"`
}

type Query struct {
}

type ReferenceRecord struct {
	Title *string `json:"title,omitempty"`
	URL   *string `json:"url,omitempty"`
}

type WorkflowRecord struct {
	ID        types.WorkflowID `json:"id"`
	CreatedAt time.Time        `json:"createdAt"`
	Alert     *AlertRecord     `json:"alert"`
	Actions   []*ActionRecord  `json:"actions"`
}
