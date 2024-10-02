package folder

import "github.com/gofrs/uuid"

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// GetAllChildFolders component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error)

	// MoveFolder component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type driver struct {
	dirToOrgMap map[string]uuid.UUID
	folders     []Folder
}

func NewDriver(folders []Folder) IDriver {
	dirToOrgMap := make(map[string]uuid.UUID)
	/*
		Create a map between folder name and which org the folder belongs to
		This allows for us to quickly
			- Check which org a folder belongs to
			- check whether a folder exists
	*/
	for _, folder := range folders {
		dirToOrgMap[folder.Name] = folder.OrgId
	}

	return &driver{
		dirToOrgMap: dirToOrgMap,
		folders:     folders,
	}
}
