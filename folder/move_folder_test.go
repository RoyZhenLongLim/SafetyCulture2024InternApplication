package folder_test

import (
	"errors"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	/*
		Ideally the test should be defined in another file
		This is ok for small test, but I would spend more time re-working this to work with a test folder
		For the interest of time and because I want to discuss more during the interview, I will leave this here
	*/
	testData := []folder.Folder{
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
	testWant1 := []folder.Folder{
		{
			Name:  "alpha",
			Paths: "alpha",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "bravo",
			Paths: "alpha.delta.bravo",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "charlie",
			Paths: "alpha.delta.bravo.charlie",
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
	testWant2 := []folder.Folder{
		{
			Name:  "alpha",
			Paths: "alpha",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "bravo",
			Paths: "golf.bravo",
			OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
		},
		{
			Name:  "charlie",
			Paths: "golf.bravo.charlie",
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

	tests := [...]struct {
		name     string
		dst      string
		folders  []folder.Folder
		want     []folder.Folder
		err      error
		_comment string
	}{
		{
			name:     "bravo",
			dst:      "delta",
			folders:  testData,
			want:     testWant1,
			err:      nil,
			_comment: "Valid Move (First Test in README.md)",
		},
		{
			name:     "bravo",
			dst:      "golf",
			folders:  testData,
			want:     testWant2,
			err:      nil,
			_comment: "Valid Move (Second Test in README.md)",
		},
		{
			name:     "bravo",
			dst:      "charlie",
			folders:  testData,
			want:     []folder.Folder{},
			err:      errors.New("cannot move a folder to a child of itself"),
			_comment: "Invalid Move (cannot move a folder to a child of itself)",
		},
		{
			name:     "bravo",
			dst:      "bravo",
			folders:  testData,
			want:     []folder.Folder{},
			err:      errors.New("cannot move a folder to itself"),
			_comment: "Invalid Move (cannot move a folder to itself)",
		},
		{
			name:     "bravo",
			dst:      "foxtrot",
			folders:  testData,
			want:     []folder.Folder{},
			err:      errors.New("cannot move a folder to a different organization"),
			_comment: "Invalid Move (cannot move a folder to a different organization)",
		},
		{
			name:     "invalid_folder",
			dst:      "delta",
			folders:  testData,
			want:     []folder.Folder{},
			err:      errors.New("cannot move a folder to a different organization"),
			_comment: "Invalid Move (cannot move a folder to a different organization)",
		},
		{
			name:     "bravo",
			dst:      "invalid_folder",
			folders:  testData,
			want:     []folder.Folder{},
			err:      errors.New("destination folder does not exist"),
			_comment: "Invalid Move (destination folder does not exist)",
		},
	}

	for _, tt := range tests {
		t.Run(tt._comment, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			move, err := f.MoveFolder(tt.name, tt.dst)
			assert.Equal(t, move, tt.want)
			errors.Is(err, tt.err)
		})
	}
}
