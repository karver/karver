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
}

func Run(target string) (*Migration, error) {
}

func (m *Migration) IsRun(current *Migration) (bool, error) {

}

func List() ([]*Migration, error) {
	//files, _ := ioutil.ReadDir("./migrations")
	//for _, f := range files {
	//	m, _ := regexp.MatchString("[0-9]{14}.*", f.Name())
	//	if m && !f.IsDir() {
	//		fmt.Println(f.Name())
	//		Load(f.Path())
	//	}
	//}
}

func Last(target string) (string, error) {

}

func Create(name string) (*Migration, error) {

}

func Load(path string) (*Migration, error) {

}
