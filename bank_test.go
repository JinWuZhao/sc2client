package sc2client

import (
	"testing"
)

func TestBank_Load(t *testing.T) {
	bank, err := NewBank("stararena")
	if err != nil {
		t.Errorf("NewBank() error: %s", err)
		return
	}
	err = bank.Load()
	if err != nil {
		t.Errorf("bank.Load() error: %s", err)
		return
	}
	t.Logf("%s", bank.data)
}

func TestBank_Save(t *testing.T) {
	bank, err := NewBank("stararena")
	if err != nil {
		t.Errorf("NewBank() error: %s", err)
		return
	}
	err = bank.Load()
	if err != nil {
		t.Errorf("bank.Load() error: %s", err)
		return
	}
	err = bank.Save()
	if err != nil {
		t.Errorf("bank.Save() error: %s", err)
		return
	}
}

func TestBank_StoreValue(t *testing.T) {
	bank, err := NewBank("stararena")
	if err != nil {
		t.Errorf("NewBank() error: %s", err)
		return
	}
	err = bank.Load()
	if err != nil {
		t.Errorf("bank.Load() error: %s", err)
		return
	}
	bank.StoreKey("rank", "first", BankValue{
		Type:  BankValueTypeString,
		Value: "星际竞技场",
	})
	err = bank.Save()
	if err != nil {
		t.Errorf("bank.Save() error: %s", err)
		return
	}
}

func TestBank_LoadValue(t *testing.T) {
	bank, err := NewBank("stararena")
	if err != nil {
		t.Errorf("NewBank() error: %s", err)
		return
	}
	err = bank.Load()
	if err != nil {
		t.Errorf("bank.Load() error: %s", err)
		return
	}
	value, ok := bank.LoadKey("rank", "first")
	if ok {
		t.Log(value)
	}
}
