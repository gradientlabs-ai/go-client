package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ExperimentalConfig struct {
	MaxDailyConversations int
}

type ProcedureVersion struct {
	// Name is the user-given name of the procedure at the time of this version.
	Name string
	// Description of the procedure at the time of this version.
	Description string
	// Version is a numeric identifier for the version. It is incremented every
	// time a new version of the procedure is saved.
	Version int
	// Author is the ID of the user who created this version of the procedure.
	Author string
	// Created is the time at which this version of the procedure was created.
	Created time.Time
	// Experimental indicates whether this is an experimental version that is used before "live",
	// within the daily limit defined in ExperimentalConfig.
	Experimental bool
	// ExperimentalConfig defines how the experimental version is limited. Relevant only if Experimental == true.
	ExperimentalConfig *ExperimentalConfig
	// Live indicates whether this is the "production" version that is used by the agent by default,
	// if there are no experimental versions or all of them have exceeded their limit.
	Live bool
}

type ListProcedureVersionsResponse struct {
	Versions []*ProcedureVersion `json:"versions"`
}

func (c *Client) ListProcedureVersions(ctx context.Context, procedureID string) (*ListProcedureVersionsResponse, error) {
	rsp, err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("procedures/%s/versions", procedureID), nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var result ListProcedureVersionsResponse
	if err := json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
