package models

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type SlotStatus string

var ErrSlotExisted = errors.New("slots already existed")
var ErrUnknownSlotStatus = errors.New("unknown slot status, slot status should be " +
	"(online, offline, migrate, pre_migrate)")

type SlotMigrateStatus struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type SlotMultiSetParam struct {
	From    int        `json:"from"`
	To      int        `json:"to"`
	Status  SlotStatus `json:"status"`
	GroupId int        `json:"group_id"`
}

type SlotState struct {
	Status        SlotStatus        `json:"status"`
	MigrateStatus SlotMigrateStatus `json:"migrate_status"`
	LastOpTs      string            `json:"last_op_ts"` // operation timestamp
}

type Slot struct {
	ProductName string    `json:"product_name"`
	Id          int       `json:"id"`
	GroupId     int       `json:"group_id"`
	State       SlotState `json:"state"`
}

func (s *Slot) String() string {
	b, _ := json.MarshalIndent(s, "", "  ")
	return string(b)
}

func NewSlot(productName string, id int) *Slot {
	return &Slot{
		ProductName: productName,
		Id:          id,
		GroupId:     INVALID_ID,
		State: SlotState{
			Status:   SLOT_STATUS_OFFLINE,
			LastOpTs: "0",
			MigrateStatus: SlotMigrateStatus{
				From: INVALID_ID,
				To:   INVALID_ID,
			},
		},
	}
}
