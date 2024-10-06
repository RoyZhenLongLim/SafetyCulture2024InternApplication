package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	/*
		Move folder name to destination dst
		Errors:
			Cannot move a folder to a child or itself
			Cannot move a folder to itself
			Source folder does not exist
			Destination folder does not exist
			Cannot move folder to different organization
	*/
	orgSrc, existSrc := f.dirToOrgMap[name]
	orgDst, existDst := f.dirToOrgMap[dst]

	if name == dst {
		return []Folder{}, errors.New("cannot move a folder to itself")
	} else if !existSrc {
		return []Folder{}, errors.New("source folder does not exist")
	} else if !existDst {
		return []Folder{}, errors.New("destination folder does not exist")
	} else if orgSrc != orgDst {
		return []Folder{}, errors.New("cannot move a folder to a different organization")
	}

	// Determine whether dst is a child of name by creating a map containing all the children
	children := f.GetAllChildFolders(orgSrc, name)
	childrenMap := map[string]struct{}{}

	for _, child := range children {
		childrenMap[child.Name] = struct{}{}
	}

	_, dstExistsAsChild := childrenMap[dst]
	if dstExistsAsChild {
		return []Folder{}, errors.New("cannot move a folder to a child of itself")
	}

	// Find dst Folder
	dstFolder, err := f.GetFolder(orgDst, dst)
	if err != nil {
		return []Folder{}, errors.New("destination folder doest not exist")
	}
	dstFolderPrefix := dstFolder.Paths

	res := []Folder{}

	for _, f := range f.folders {
		newFile := Folder{
			Name:  f.Name,
			Paths: f.Paths,
			OrgId: f.OrgId,
		}
		// if we find the source file, then update the path for it
		if newFile.Name == name {
			newFile.Paths = dstFolderPrefix + "." + newFile.Name
		}
		// If we find a child of source, then update that path
		_, exists := childrenMap[newFile.Name]
		if exists {
			// Split the path into two
			// first part is prefix + . + child path
			parts := strings.SplitN(newFile.Paths, name, 3)
			newFile.Paths = dstFolderPrefix + "." + name + parts[1]
		}
		res = append(res, newFile)
	}
	return res, nil
}
