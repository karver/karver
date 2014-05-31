package migrations

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
	return "", nil
}

func RunMigrations(target string) (*Migration, error) {
	return nil, nil
}

func (m *Migration) IsRun(current *Migration) (bool, error) {
	return false, nil
}

func ListMigrations() ([]*Migration, error) {
	//files, _ := ioutil.ReadDir("./migrations")
	//for _, f := range files {
	//	m, _ := regexp.MatchString("[0-9]{14}.*", f.Name())
	//	if m && !f.IsDir() {
	//		fmt.Println(f.Name())
	//		Load(f.Path())
	//	}
	//}
	return nil, nil
}

func LastMigration(target string) (string, error) {
	return "", nil
}

func CreateMigration(name string) (*Migration, error) {
	return nil, nil
}

func LoadMigration(path string) (*Migration, error) {
	return nil, nil
}
