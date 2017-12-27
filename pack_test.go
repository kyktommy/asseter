package asseter_test

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/kyktommy/asseter"
)

type TestFile struct {
	Images map[string][]TestFileImage `json:"images"`
}

type TestFileImage struct {
	Key   string              `json:"key"`
	Image TestImageTranslated `json:"image"`
}

type TestImageTranslated struct {
	EnUS string `json:"en-US"`
	ZhTW string `json:"zh-TW"`
}

func (testFile *TestFile) GetAssets(filename string) ([]string, error) {
	ret := []string{}
	ba, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var tf TestFile
	err = json.Unmarshal(ba, &tf)
	if err != nil {
		return nil, err
	}
	for _, a := range tf.Images {
		for _, img := range a {
			ret = append(ret, img.Image.EnUS)
			ret = append(ret, img.Image.ZhTW)
		}
	}
	return ret, nil
}

func TestPack(t *testing.T) {
	tf := TestFile{}
	assets, err := tf.GetAssets("./test.json")

	if err != nil {
		t.Error(err)
	}

	packed, err := asseter.Pack(asseter.PackConfig{
		BaseURL: "/statics",
		Assets:  assets,
	})

	if err != nil {
		t.Error(err)
	}

	if len(packed.Output) != 2 {
		t.Error("should have 2 images packed with hash")
		return
	}

	for _, o := range packed.Output {
		if o == "" {
			t.Errorf("should not be empty")
		}
		if !strings.HasPrefix(o, "/statics/images/") {
			t.Errorf("should prefix with baseURL and nested folder path")
		}
	}

	for i := range packed.Output {
		if path.Ext(packed.Input[i]) != path.Ext(packed.Output[i]) {
			t.Errorf("should have equal extension")
			return
		}
	}
}
