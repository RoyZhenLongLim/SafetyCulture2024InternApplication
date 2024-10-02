package folder

import (
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
	// Note for marker. We could alternatively use regex for this :)
	// We can check that name followed by a dot , e.g. "steady-insect." exists in the string
	// Both have tradeoffs, so I chose the simpler method to verify
	// I suspect this method is slower but easier to maintain and improve,
	// which I believe is far more important for a product that is continuously updated

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
