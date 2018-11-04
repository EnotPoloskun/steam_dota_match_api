package steamdotaapi

import (
  "fmt"
  "net/url"
)

type MatchResultJson struct {
  Result MatchResult
}

type MatchResult struct {
  Players []Player
  Season uint `json:"item_0,omitempty"`
  RadiantWin bool `json:"radiant_win"`
  Duration uint
  MatchStart uint `json:"start_time"`
  MatchId uint64 `json:"match_id"`
  MatchSequenceNo uint `json:"match_seq_num"`
  TowerStatusRadiant TowerStatus `json:"tower_status_radiant"`
  TowerStatusDire TowerStatus `json:"tower_status_dire"`
  BarracksStatusRadiant BarracksStatus `json:"barracks_status_radiant"`
  BarracksStatusDire BarracksStatus `json:"barracks_status_dire"`
  Cluster uint
  FirstBloodTime int `json:"first_blood_time"`
  LobbyType LobbyType `json:"lobby_type"`
  HumanPlayers uint `json:"human_players"`
  LeagueId uint
  PositiveVotes uint `json:"positive_votes"`
  NegativeVotes uint `json:"negative_votes"`
  GameMode GameMode `json:"game_mode"`
  PicksBans []PickBan `json:"picks_bans,omitempty"`
  RadiantScore uint `json:"radiant_score"`
  DireScore uint `json:"dire_score"`
}

type Player struct {
  AccountId uint32 `json:"account_id"`
  Slot PlayerSlot `json:"player_slot"`
  HeroId uint `json:"hero_id"`
  Item0 uint `json:"item_0"`
  Item1 uint `json:"item_1"`
  Item2 uint `json:"item_2"`
  Item3 uint `json:"item_3"`
  Item4 uint `json:"item_4"`
  Item5 uint `json:"item_5"`
  Kills uint
  Deaths uint
  Assists uint
  LeaverStatus LeaverStatus `json:"leaver_status"`
  GoldRemaining uint `json:"gold"`
  LastHits uint `json:"last_hits"`
  Denies uint `json:"denies"`
  GPM uint `json:"gold_per_min"`
  XPM uint `json:"xp_per_min"`
  GoldSpent uint `json:"gold_spent"`
  HeroDamage uint `json:"hero_damage"`
  TowerDamage uint `json:"tower_damage"`
  HeroHealing uint `json:"hero_healing"`
  Level uint
  Abilities []Ability `json:"ability_upgrades"`
  Units []Unit `json:"additional_units,omitempty"`
}

type Ability struct {
  Id uint `json:"ability"`
  TimeUpgraded int `json:"time"`
  Level uint
}

type Unit struct {
  Name string `json:"unitname"`
  Item0 uint `json:"item_0"`
  Item1 uint `json:"item_1"`
  Item2 uint `json:"item_2"`
  Item3 uint `json:"item_3"`
  Item4 uint `json:"item_4"`
  Item5 uint `json:"item_5"`
}

type PickBan struct {
  IsPick bool `json:"is_pick"`
  HeroId uint `json:"hero_id"`
  Team Team
  Order uint
}

type GameMode int

const (
  NoneMode GameMode = iota
  AllPick
  CaptainsMode
  RandomDraft
  SingleDraft
  AllRandom
  Intro
  Diretide
  ReverseCaptainsMode
  TheGreeviling
  TutorialMode
  MidOnly
  LeastPlayed
  NewPlayerPool
  CompendiumMatchmaking
  CoopVsBots
  CaptainsDraft
  _
  AbilityDraft
  _
  AllRandomDeathmatch
  OneVsOneMidOnly
  RankedMatchmaking
  TurboMode
)

type LeaverStatus uint

const (
  NoneStatus LeaverStatus = iota
  Disconnected
  DisconnectedTooLong
  Abandoned
  AFK
  NeverConnected
  NeverConnectedTooLong
)

type Team uint

const (
  Radiant Team = iota
  Dire
)

type TowerStatus uint16

const (
  TopTierOne TowerStatus = (1 << iota)
  TopTierTwo
  TopTierThree
  MiddleTierOne
  MiddleTierTwo
  MiddleTierThree
  BottomTierOne
  BottomTierTwo
  BottomTierThree
  AncientTop
  AncientBottom
)

func (s TowerStatus) IsDestroyed(tower TowerStatus) bool {
  return (s & tower) == 0
}

type BarracksStatus uint8

const (
  TopMelee BarracksStatus = (1 << iota)
  TopRanged
  MiddleMelee
  MiddleRanged
  BottomMelee
  BottomRanged
)

func (s BarracksStatus) IsDestroyed(barrack BarracksStatus) bool {
  return (s & barrack) == 0
}


func (api DotaApi) GetMatchDetails(params url.Values, target *MatchResultJson) error {
  url := fmt.Sprintf("%v/IDOTA2Match_%v/GetMatchDetails/v1", BaseUrl, DotaAppID)
  err := api.MakeRequest(url, params, target)

  if err != nil {
    return err
  }

  return nil
}



