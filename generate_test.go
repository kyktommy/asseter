package asseter_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/kyktommy/asseter"
)

func TestMain(m *testing.M) {
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestGenerate(t *testing.T) {

	err := asseter.Generate(asseter.GenerateConfig{
		AssetConfig: "./test.json",
		InputDir:    "./assets",
		InputAssets: []string{
			"/images/100.jpeg",
			"/images/404.jpeg",
		},
		OutputAssets: []string{
			"/statics/images/69382363afbc19916a477f7acab11023.jpeg",
			"/statics/images/c7cecb798ecd0b0d467c019bcc362ba1.jpeg",
		},
		OutputDir: "dist",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if !isFileExist("./dist/statics/images/69382363afbc19916a477f7acab11023.jpeg") ||
		!isFileExist("./dist/statics/images/c7cecb798ecd0b0d467c019bcc362ba1.jpeg") {
		t.Errorf("dist assets is not exist")
		return
	}

	if !isFileExist("./dist/test.json") {
		t.Errorf("dist config file is not exist")
		return
	}

	ba, err := ioutil.ReadFile("./dist/test.json")
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(ba), "/statics/images/69382363afbc19916a477f7acab11023.jpeg") ||
		!strings.Contains(string(ba), "/statics/images/c7cecb798ecd0b0d467c019bcc362ba1.jpeg") {
		t.Errorf("invalid asset config file")
		return
	}

}

func tearDown() {
	if err := os.RemoveAll("./dist"); err != nil {
		fmt.Println("fail to remove dist")
	}
}

func isFileExist(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
