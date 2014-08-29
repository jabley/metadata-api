package main

import (
	"encoding/json"
)

type Organisation struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Status       string   `json:"govuk_status"`
	Abbreviation string   `json:"abbreviation"`
	ParentIDs    []string `json:"parent_ids"`
	ChildIDs     []string `json:"child_ids"`
}

type Need struct {
	ID                 int            `json:"id"`
	Role               string         `json:"role"`
	Goal               string         `json:"goal"`
	Benefit            string         `json:"benefit"`
	OrganisationIDs    []string       `json:"organisation_ids"`
	Organisations      []Organisation `json:"organisations"`
	Justifications     []string       `json:"justifications"`
	Impact             string         `json:"impact"`
	MetWhen            []string       `json:"met_when"`
	YearlyUserContacts int            `json:"yearly_user_contacts"`
	YearlySiteViews    int            `json:"yearly_site_views"`
	YearlyNeedViews    int            `json:"yearly_need_views"`
	YearlySearches     int            `json:"yearly_searches"`
	OtherEvidence      string         `json:"other_evidence"`
	Legislation        string         `json:"legislation"`
	AllOrganisations   bool           `json:"applies_to_all_organisations"`
	InScope            bool           `json:"in_scope"`
	OutOfScopeReason   string         `json:"out_of_scope_reason"`
	DuplicateOf        int            `json:"duplicate_of"`
}

func ParseNeedResponse(response []byte) (*Need, error) {
	need := &Need{}
	if err := json.Unmarshal(response, &need); err != nil {
		return nil, err
	}

	return need, nil
}