package nospaceleftondevice_test

import (
	"testing"

	nospaceleftondevice "github.com/mhdiiilham/adventofcode/2022/no-space-left-on-device"
	"github.com/stretchr/testify/assert"
)

const input = "../../input/2022/nslod_test.txt"

func TestFileNotFound(t *testing.T) {
	assertion := assert.New(t)

	_, err := nospaceleftondevice.New("")
	assertion.Error(err)

}

func TestGetTotalSizeOfDirectories(t *testing.T) {
	assertion := assert.New(t)

	noSpaceLeftOnDeviceClient, err := nospaceleftondevice.New(input)
	assertion.NoError(err)

	actual := noSpaceLeftOnDeviceClient.GetTotalSizeOfDirectories()
	assertion.Equal(95437, actual)
}

func TestGetTotalSizeDirectoryNeedToDelete(t *testing.T) {
	assertion := assert.New(t)

	noSpaceLeftOnDeviceClient, err := nospaceleftondevice.New(input)
	assertion.NoError(err)

	actual := noSpaceLeftOnDeviceClient.GetTotalSizeDirectoryNeedToDelete()
	assertion.Equal(24933642, actual)
}
