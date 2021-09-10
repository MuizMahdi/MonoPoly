package helpers

type Stage struct {
	Name        string
	Description string
	Actors      []Actor
}

type Actor struct {
	Name        string
	Description string
}

var WorkspacesTypes = newWorkspaceType()

func newWorkspaceType() *WorkspaceType {
	return &WorkspaceType{
		Stage: "stage",
		Actor: "actor",
	}
}

type WorkspaceType struct {
	Stage string
	Actor string
}
