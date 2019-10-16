package main

type Guild struct {
	Score     int    `json:"score"`
	WorldRank int    `json:"world_rank"`
	AreaRank  int    `json:"area_rank"`
	RealmRank int    `json:"realm_rank"`
	Name      string `json:"name"`
	Url       string `json:"url"`
}
