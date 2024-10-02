package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	/*
		 Best Practises dictates that I add the files somewhere
		For the interest of time / clarity of what my tests does, I will put them into this file
	*/

	tests := [...]struct {
		name     string
		orgID    uuid.UUID
		folders  []folder.Folder
		want     []folder.Folder
		_comment string
	}{
		{
			name:     "alpha",
			orgID:    uuid.FromStringOrNil("org1"),
			folders:  []folder.Folder{},
			want:     []folder.Folder{},
			_comment: "Should return empty array if not given any folders",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)

		})
	}
}
