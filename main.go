package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	//orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := []folder.Folder{
		{
			Name:  "alpha",
			Paths: "alpha",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "bravo",
			Paths: "alpha.bravo",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "charlie",
			Paths: "alpha.bravo.charlie",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "delta",
			Paths: "alpha.delta",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "echo",
			Paths: "alpha.delta.echo",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "foxtrot",
			Paths: "foxtrot",
			OrgId: uuid.FromStringOrNil(folder.InvalidOrgID),
		},
		{
			Name:  "golf",
			Paths: "golf",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
	}
	// This bit of code fetches sample data that details a file system
	// res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	//orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	//childFolder := folderDriver.GetAllChildFolders(orgID, "noble-vixen")

	//folder.PrettyPrint(res)
	after, err := folderDriver.MoveFolder("delta", "charlie")
	folder.PrettyPrint(after)
	fmt.Println(err)
	//fmt.Printf("\n Folders for orgID: %s", orgID)
	//folder.PrettyPrint(orgFolder)
	//folder.PrettyPrint(childFolder)
}
