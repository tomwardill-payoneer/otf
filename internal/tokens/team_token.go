package tokens

import (
	"context"
	"time"

	"github.com/leg100/otf/internal"
	"github.com/leg100/otf/internal/rbac"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type (
	// TeamToken provides information about an API token for a user.
	TeamToken struct {
		ID          string
		CreatedAt   time.Time
		Description string

		// Token belongs to an team
		Team string
		// Optional expiry.
		Expiry *time.Time
	}

	// CreateTeamTokenOptions are options for creating an team token via the service
	// endpoint
	CreateTeamTokenOptions struct {
		Team   string `schema:"team_id,required"`
		Expiry *time.Time
	}

	// NewTeamTokenOptions are options for constructing a user token via the
	// constructor.
	NewTeamTokenOptions struct {
		CreateTeamTokenOptions
		Team string
		key  jwk.Key
	}

	teamTokenService interface {
		// CreateTeamToken creates a user token.
		CreateTeamToken(ctx context.Context, opts CreateTeamTokenOptions) (*TeamToken, []byte, error)
		// GetTeamToken gets the team token. If a token does not
		// exist, then nil is returned without an error.
		GetTeamToken(ctx context.Context, team string) (*TeamToken, error)
		// DeleteTeamToken deletes an team token.
		DeleteTeamToken(ctx context.Context, tokenID string) error
	}
)

func NewTeamToken(opts NewTeamTokenOptions) (*TeamToken, []byte, error) {
	tt := TeamToken{
		ID:        internal.NewID("tt"),
		CreatedAt: internal.CurrentTimestamp(),
		Team:      opts.Team,
		Expiry:    opts.Expiry,
	}
	token, err := NewToken(NewTokenOptions{
		key:     opts.key,
		Subject: tt.ID,
		Kind:    teamTokenKind,
		Expiry:  opts.Expiry,
	})
	if err != nil {
		return nil, nil, err
	}
	return &tt, token, nil
}

func (a *service) CreateTeamToken(ctx context.Context, opts CreateTeamTokenOptions) (*TeamToken, []byte, error) {

	_, err := a.team.CanAccess(ctx, rbac.CreateTeamTokenAction, opts.Team)
	if err != nil {
		return nil, nil, err
	}

	tt, token, err := NewTeamToken(NewTeamTokenOptions{
		CreateTeamTokenOptions: opts,
		Team:                   opts.Team,
		key:                    a.key,
	})
	if err != nil {
		a.Error(err, "constructing team token", "team", opts.Team)
		return nil, nil, err
	}

	if err := a.db.createTeamToken(ctx, tt); err != nil {
		a.Error(err, "creating team token", "team", opts.Team)
		return nil, nil, err
	}

	a.V(0).Info("created team token", "team", opts.Team)

	return tt, token, nil
}

func (a *service) GetTeamToken(ctx context.Context, team string) (*TeamToken, error) {
	return a.db.getTeamTokenByName(ctx, team)
}

func (a *service) DeleteTeamToken(ctx context.Context, team string) error {
	_, err := a.team.CanAccess(ctx, rbac.CreateTeamTokenAction, team)
	if err != nil {
		return err
	}

	if err := a.db.deleteTeamToken(ctx, team); err != nil {
		a.Error(err, "deleting team token", "team", team)
		return err
	}

	a.V(0).Info("deleted team token", "team", team)

	return nil
}
