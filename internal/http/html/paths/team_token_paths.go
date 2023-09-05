// Code generated by "go generate"; DO NOT EDIT.

package paths

import "fmt"

func TeamTokens(team string) string {
	return fmt.Sprintf("/app/teams/%s/team-tokens", team)
}

func CreateTeamToken(team string) string {
	return fmt.Sprintf("/app/teams/%s/team-tokens/create", team)
}

func NewTeamToken(team string) string {
	return fmt.Sprintf("/app/teams/%s/team-tokens/new", team)
}

func TeamToken(teamToken string) string {
	return fmt.Sprintf("/app/team-tokens/%s", teamToken)
}

func EditTeamToken(teamToken string) string {
	return fmt.Sprintf("/app/team-tokens/%s/edit", teamToken)
}

func UpdateTeamToken(teamToken string) string {
	return fmt.Sprintf("/app/team-tokens/%s/update", teamToken)
}

func DeleteTeamToken(teamToken string) string {
	return fmt.Sprintf("/app/team-tokens/%s/delete", teamToken)
}
