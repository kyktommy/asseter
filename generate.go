package asseter

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type GenerateConfig struct {
	AssetConfig  string
	InputAssets  []string
	InputDir     string
	OutputAssets []string
	OutputDir    string
}

func Generate(c GenerateConfig) error {
	// read config
	ba, err := ioutil.ReadFile(c.AssetConfig)
	if err != nil {
		return err
	}
	cfg := string(ba)

	for i := range c.InputAssets {
		// update config
		cfg = strings.Replace(cfg, c.InputAssets[i], c.OutputAssets[i], -1)

		// read file
		f, err := ioutil.ReadFile(path.Join(".", c.InputDir, c.InputAssets[i]))
		if err != nil {
			return err
		}
		// mkdir
		if err = os.MkdirAll(path.Join(c.OutputDir, path.Dir(c.OutputAssets[i])), 0777); err != nil {
			return err
		}
		// output file
		if err = ioutil.WriteFile(path.Join(c.OutputDir, c.OutputAssets[i]), f, 0644); err != nil {
			return err
		}
	}

	// output config
	cfgFilename := path.Base(c.AssetConfig)
	err = ioutil.WriteFile(path.Join(c.OutputDir, cfgFilename), []byte(cfg), 0644)
	if err != nil {
		return err
	}

	return nil
}
