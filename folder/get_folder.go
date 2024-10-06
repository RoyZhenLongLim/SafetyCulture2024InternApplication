package folder

import (
	"errors"
	"github.com/gofrs/uuid"
	"strings"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}
	return res
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	/*
		If folder doesn't exist or has incorrect orgID, then return empty array with error
	*/
	org, exist := f.dirToOrgMap[name]
	if !exist || org != orgID {
		return []Folder{}
	}

	// Get all folders with specific org id
	res := []Folder{}
	for _, f := range f.GetFoldersByOrgID(orgID) {
		// If this is child folder of name, append to list
		// File path must be in the format name1.name2.name3 etc
		pathNames := strings.Split(f.Paths, ".")
		/*
			To check if this is a valid child directory, we need to check the following
				- orgID must match input orgID
				- pathNames has length more than 1 (otherwise we have a root directory)
				- last string in path is equal to the name of the folder (otherwise data is invalid)
				- name exists in path and is not the last index (ensure that it is a child and not a parent)
		*/
		if len(pathNames) > 1 && pathNames[len(pathNames)-1] == f.Name {
			for i, n := range pathNames {
				if n == name && i < len(pathNames)-1 {
					res = append(res, f)
					// If we confirm it is a subdirectory, we can ignore the rest of the paths names
					break
				}
			}
		}
	}

	return res
}

func (f *driver) GetFolder(orgID uuid.UUID, name string) (Folder, error) {
	/*
		Returns folder with name that belongs to org with orgId
		Function should
	*/
	for _, f := range f.folders {
		if f.Name == name && f.OrgId == orgID {
			return f, nil
		}
	}
	return Folder{}, errors.New("folder does not exists")
}
