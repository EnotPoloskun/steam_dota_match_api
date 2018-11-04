package steamdotaapi

import (
  "fmt"
  "net/url"
)

type HistoryResult struct {
  Status uint
  StatusDetails string `json:",omitempty"`
  NumResult uint64 `json:"num_results"`
  TotalResults uint `json:"total_results"`
  ResultsRemaining uint `json:"results_remaining"`
  Matches []Match
}

type Match struct {
  MatchId uint64 `json:"match_id"`
  MatchSequenceNo uint `json:"match_seq_num"`
  MatchStart uint `json:"start_time"`
  LobbyType LobbyType `json:"lobby_type"`
  Players []PlayerInfo
}

type MatchHistoryJson struct {
  Result HistoryResult `json:"result"`
}

type PlayerInfo struct {
  AccountId uint32 `json:"account_id"`
  PlayerSlot PlayerSlot `json:"player_slot"`
  HeroId uint `json:"hero_id"`
}

func (api DotaApi) GetMatchHistory(params url.Values, target *MatchHistoryJson) error {
  url := fmt.Sprintf("%v/IDOTA2Match_%v/GetMatchHistory/v1", BaseUrl, DotaAppID)
  err := api.MakeRequest(url, params, target)

  if err != nil {
    return err
  }

  return nil
}
