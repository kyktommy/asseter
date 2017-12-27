package asseter

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"path"
)

type PackConfig struct {
	BaseURL string
	Assets  []string
}

type PackedAssets struct {
	Input  []string
	Output []string
}

func Pack(c PackConfig) (*PackedAssets, error) {
	packed := &PackedAssets{
		Input:  c.Assets,
		Output: []string{},
	}
	for _, f := range packed.Input {
		ba, err := ioutil.ReadFile(path.Join(".", f))
		if err != nil {
			return nil, err
		}
		basepath := path.Dir(f)
		checksum := md5.Sum(ba)
		extension := path.Ext(f)
		filename := fmt.Sprintf("%x"+extension, checksum)
		filepath := path.Join(c.BaseURL, basepath, filename)
		packed.Output = append(packed.Output, filepath)
	}
	return packed, nil
}
