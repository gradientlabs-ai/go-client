package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProcedureListParams struct {
	// Cursor is used to retrieve the next/previous page of the list.
	Cursor string `json:"cursor,omitempty"`

	// Status is used to filter the list of procedures by status.
	Status ProcedureStatus `json:"status,omitempty"`
}

type ProcedureListResponse struct {
	// Procedures contains the list of procedures.
	Procedures []*Procedure `json:"procedures"`

	// Pagination contains the pagination-related information.
	Pagination *PaginationInfo `json:"pagination"`
}

// ListProcedures lists procedures.
//
// Note: requires a `Management` API key.
func (c *Client) ListProcedures(ctx context.Context, p *ProcedureListParams) (*ProcedureListResponse, error) {
	path := "procedures"
	if p.Cursor != "" {
		path = fmt.Sprintf("%v?cursor=%v", path, p.Cursor)
	}
	if p.Status != "" {
		path = fmt.Sprintf("%v?status=%v", path, p.Status)
	}

	rsp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if err := responseError(rsp); err != nil {
		return nil, err
	}

	var procs ProcedureListResponse
	if err := json.NewDecoder(rsp.Body).Decode(&procs); err != nil {
		return nil, err
	}
	return &procs, nil
}
