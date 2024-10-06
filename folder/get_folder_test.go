package folder_test

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	/*
		Best Practises dictates that I add the files somewhere
		For the interest of time / clarity of what my tests does, I will put them into this file
	*/

	tests := [...]struct {
		orgID    uuid.UUID
		folders  []folder.Folder
		want     []folder.Folder
		_comment string
	}{
		{
			orgID:    uuid.FromStringOrNil(folder.DefaultOrgID),
			folders:  []folder.Folder{},
			want:     []folder.Folder{},
			_comment: "Empty folders yields empty result",
		},
		{
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "beta",
					Paths: "beta",
					OrgId: uuid.FromStringOrNil(folder.InvalidOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			_comment: "Only returns one folder with associated org id",
		},
		{
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "beta",
					Paths: "alpha.beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "beta",
					Paths: "beta",
					OrgId: uuid.FromStringOrNil(folder.InvalidOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "beta",
					Paths: "alpha.beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			_comment: "Only returns one folders (more than one) with associated org id",
		},
		{
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "beta",
					Paths: "beta",
					OrgId: uuid.FromStringOrNil(folder.InvalidOrgID),
				},
			},
			want:     []folder.Folder{},
			_comment: "No valid files or folders in input should yield empty result",
		},
	}
	for _, tt := range tests {
		t.Run(tt._comment, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
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
			orgID:    uuid.FromStringOrNil(folder.DefaultOrgID),
			folders:  []folder.Folder{},
			want:     []folder.Folder{},
			_comment: "Empty input folders yields empty result",
		},
		{
			name:  "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "beta",
					Paths: "beta",
					OrgId: uuid.FromStringOrNil(folder.InvalidOrgID),
				},
			},
			want:     []folder.Folder{},
			_comment: "Folders with no correct org id should yield empty result",
		},
		{
			name:  "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:     []folder.Folder{},
			_comment: "Root directory are not returned",
		},
		{
			name:  "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "beta",
					Paths: "alpha.beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "beta",
					Paths: "alpha.beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			_comment: "Direct Children are returned",
		},
		{
			name:  "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "beta",
					Paths: "alpha.beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "gamma",
					Paths: "alpha.beta.gamma",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "beta",
					Paths: "alpha.beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "gamma",
					Paths: "alpha.beta.gamma",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			_comment: "Nested Children are returned (i.e. not directories not at the root)",
		},
		{
			name:  "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "gamma",
					Paths: "alpha.beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:     []folder.Folder{},
			_comment: "Incorrect children are not returned (name is gamma but paths says beta)",
		},
	}
	for _, tt := range tests {
		t.Run(tt._comment, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.name)
			assert.Equal(t, tt.want, get)
		})
	}
}
