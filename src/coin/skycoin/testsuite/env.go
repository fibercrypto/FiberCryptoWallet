package testsuite

import (
	"go/build"
	"path"
)

const (
	TestIDToken           = "fibercryptotest"
	ManyAddressesFilename = "many-addresses.golden"
)

func GetSkycoinCipherTestDataDir() string {
	return path.Join(build.Default.GOPATH, ".", "src", "github.com", "fibercrypto", "fibercryptowallet", "vendor", "github.com", "SkycoinProject", "skycoin", "src", "cipher", "testsuite", "testdata")
}
