package xlsx_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"justanother.org/microdungeon/testdata"
	"justanother.org/microdungeon/xlsx"
)

func TestCharacter(t *testing.T) {
	f, err := xlsx.ExcelFromCharacter(testdata.Character())
	require.NoError(t, err)

	err = os.MkdirAll("testdata", 0766)
	require.NoError(t, err)

	err = f.SaveAs(filepath.Join("testdata", t.Name()+".xlsx"))
	require.NoError(t, err)

	err = f.Close()
	require.NoError(t, err)
}
