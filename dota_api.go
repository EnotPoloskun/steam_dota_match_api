package steamdotaapi

const BaseUrl string = "http://api.steampowered.com"
const DotaAppID int = 570

type DotaApi struct {
  Key string
}

type LobbyType int

const (
  Invalid LobbyType = -1
  PublicMatchMaking LobbyType = iota
  Practice
  Tournament
  Tutorial
  Coop
  TeamMatch
  SoloQueue
  Ranked
  OneVsOneMid
)

type PlayerSlot uint8

func (p PlayerSlot) IsDire() bool {
  return (p & (1 << 7) > 0)
}

func (p PlayerSlot) Position() (uint) {
  return uint(p & 7)
}
