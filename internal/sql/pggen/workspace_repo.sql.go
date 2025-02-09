// Code generated by pggen. DO NOT EDIT.

package pggen

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const insertWorkspaceRepoSQL = `INSERT INTO workspace_repos (
    branch,
    webhook_id,
    vcs_provider_id,
    workspace_id
) VALUES (
    $1,
    $2,
    $3,
    $4
);`

type InsertWorkspaceRepoParams struct {
	Branch        pgtype.Text
	WebhookID     pgtype.UUID
	VCSProviderID pgtype.Text
	WorkspaceID   pgtype.Text
}

// InsertWorkspaceRepo implements Querier.InsertWorkspaceRepo.
func (q *DBQuerier) InsertWorkspaceRepo(ctx context.Context, params InsertWorkspaceRepoParams) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "InsertWorkspaceRepo")
	cmdTag, err := q.conn.Exec(ctx, insertWorkspaceRepoSQL, params.Branch, params.WebhookID, params.VCSProviderID, params.WorkspaceID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query InsertWorkspaceRepo: %w", err)
	}
	return cmdTag, err
}

// InsertWorkspaceRepoBatch implements Querier.InsertWorkspaceRepoBatch.
func (q *DBQuerier) InsertWorkspaceRepoBatch(batch genericBatch, params InsertWorkspaceRepoParams) {
	batch.Queue(insertWorkspaceRepoSQL, params.Branch, params.WebhookID, params.VCSProviderID, params.WorkspaceID)
}

// InsertWorkspaceRepoScan implements Querier.InsertWorkspaceRepoScan.
func (q *DBQuerier) InsertWorkspaceRepoScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec InsertWorkspaceRepoBatch: %w", err)
	}
	return cmdTag, err
}

const updateWorkspaceRepoByIDSQL = `UPDATE workspace_repos
SET
    branch = $1
WHERE workspace_id = $2
RETURNING workspace_id;`

// UpdateWorkspaceRepoByID implements Querier.UpdateWorkspaceRepoByID.
func (q *DBQuerier) UpdateWorkspaceRepoByID(ctx context.Context, branch pgtype.Text, workspaceID pgtype.Text) (pgtype.Text, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateWorkspaceRepoByID")
	row := q.conn.QueryRow(ctx, updateWorkspaceRepoByIDSQL, branch, workspaceID)
	var item pgtype.Text
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query UpdateWorkspaceRepoByID: %w", err)
	}
	return item, nil
}

// UpdateWorkspaceRepoByIDBatch implements Querier.UpdateWorkspaceRepoByIDBatch.
func (q *DBQuerier) UpdateWorkspaceRepoByIDBatch(batch genericBatch, branch pgtype.Text, workspaceID pgtype.Text) {
	batch.Queue(updateWorkspaceRepoByIDSQL, branch, workspaceID)
}

// UpdateWorkspaceRepoByIDScan implements Querier.UpdateWorkspaceRepoByIDScan.
func (q *DBQuerier) UpdateWorkspaceRepoByIDScan(results pgx.BatchResults) (pgtype.Text, error) {
	row := results.QueryRow()
	var item pgtype.Text
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan UpdateWorkspaceRepoByIDBatch row: %w", err)
	}
	return item, nil
}

const deleteWorkspaceRepoByIDSQL = `DELETE
FROM workspace_repos
WHERE workspace_id = $1
RETURNING *
;`

type DeleteWorkspaceRepoByIDRow struct {
	Branch        pgtype.Text `json:"branch"`
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
	WorkspaceID   pgtype.Text `json:"workspace_id"`
}

// DeleteWorkspaceRepoByID implements Querier.DeleteWorkspaceRepoByID.
func (q *DBQuerier) DeleteWorkspaceRepoByID(ctx context.Context, workspaceID pgtype.Text) (DeleteWorkspaceRepoByIDRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteWorkspaceRepoByID")
	row := q.conn.QueryRow(ctx, deleteWorkspaceRepoByIDSQL, workspaceID)
	var item DeleteWorkspaceRepoByIDRow
	if err := row.Scan(&item.Branch, &item.WebhookID, &item.VCSProviderID, &item.WorkspaceID); err != nil {
		return item, fmt.Errorf("query DeleteWorkspaceRepoByID: %w", err)
	}
	return item, nil
}

// DeleteWorkspaceRepoByIDBatch implements Querier.DeleteWorkspaceRepoByIDBatch.
func (q *DBQuerier) DeleteWorkspaceRepoByIDBatch(batch genericBatch, workspaceID pgtype.Text) {
	batch.Queue(deleteWorkspaceRepoByIDSQL, workspaceID)
}

// DeleteWorkspaceRepoByIDScan implements Querier.DeleteWorkspaceRepoByIDScan.
func (q *DBQuerier) DeleteWorkspaceRepoByIDScan(results pgx.BatchResults) (DeleteWorkspaceRepoByIDRow, error) {
	row := results.QueryRow()
	var item DeleteWorkspaceRepoByIDRow
	if err := row.Scan(&item.Branch, &item.WebhookID, &item.VCSProviderID, &item.WorkspaceID); err != nil {
		return item, fmt.Errorf("scan DeleteWorkspaceRepoByIDBatch row: %w", err)
	}
	return item, nil
}
