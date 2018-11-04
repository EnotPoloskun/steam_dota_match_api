package steamdotaapi

import (
  "testing"
)

func TestPlayerSlotIsDire(t *testing.T) {
  var playerSlot1 PlayerSlot = (1 << 7)
  var playerSlot2 PlayerSlot = 0

  if playerSlot1.IsDire() != true {
    t.Errorf("not dire %v", playerSlot1)
  }

  if playerSlot2.IsDire() == true {
    t.Errorf("not radiant %v", playerSlot2)
  }
}

func TestPlayerSlotPosition(t *testing.T) {
  var playerSlot PlayerSlot = 9

  if playerSlot.Position() != 1 {
    t.Errorf("wrong position %v", playerSlot)
  }

  playerSlot = 10
  if playerSlot.Position() != 2 {
    t.Errorf("wrong position %v", playerSlot)
  }
}
