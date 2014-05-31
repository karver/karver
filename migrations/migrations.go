package migrations

import (
	"io/ioutil"
	"regexp"
)

type Migration struct {
	timestamp string
	name      string
	path      string
}

func (m *Migration) Run(target string) (string, error) {
	// cmd := exec.Command("ls", "-la")
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ls: %q\n", out.String())
}

func Run(target string) (*Migration, error) {
	return nil, nil
}

func (m *Migration) IsRun(current *Migration) (bool, error) {
	return false, nil
}

func List() ([]*Migration, error) {
	files, _ := ioutil.ReadDir("./test/migrations")
	migrations := make([]*Migration, 0, len(files))
	r, _ := regexp.Compile("([0-9]{10,14}).*")
	for _, f := range files {
		m := r.FindStringSubmatch(f.Name())
		if len(m) > 0 && !f.IsDir() {
			migration, err := Load(f.Name())
			migrations := append(migrations, migration)
		}
	}

	return migrations, nil
}

func Last(target string) (string, error) {
	return "", nil
}

func Create(name string) (*Migration, error) {
	return nil, nil
}

func Load(path string) (*Migration, error) {
	return nil, nil
}
