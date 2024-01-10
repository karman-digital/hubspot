package hubspotmodels

import "time"

type OwnerTeam struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}

type Owner struct {
	ID        string      `json:"id"`
	Email     string      `json:"email"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	UserID    int         `json:"userId"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Archived  bool        `json:"archived"`
	Teams     []OwnerTeam `json:"teams"`
}

type OwnerResponse struct {
	Results []Owner `json:"results"`
	Paging  Paging  `json:"paging"`
}
