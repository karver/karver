package migrations

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"syscall"
)

type Migration struct {
	Timestamp string
	Name      string
	Path      string
}

func (m *Migration) Run(target string) (string, string, error) {
	// Copy migration to the target path
	f, err := os.Open(m.Path)
	if err != nil {
		return "", "", fmt.Errorf("Error opening the file %s: %s", m.Path, err.Error())
	}
	defer f.Close()

	tmpf, err := ioutil.TempFile(target, "")
	if err != nil {
		return "", "", fmt.Errorf("Error creating tmp file in %s: %s", target, err.Error())
	}
	defer tmpf.Close()

	_, err = io.Copy(tmpf, f)
	if err != nil {
		return "", "", fmt.Errorf("Error copying the file %s: %s", m.Path, err.Error())
	}

	// Prepare the target environment
	if err := syscall.Chroot(target); err != nil {
		return "", "", fmt.Errorf("Error trying to chroot into %s when running migration %s: %s", target, m.Name, err.Error())
	}

	tmpf.Chmod(777)

	// Run the migration
	cmd := exec.Command(target + filepath.Base(tmpf.Name()))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", "", err
	}

	if err := cmd.Start(); err != nil {
		return "", "", fmt.Errorf("Error starting migration %s: %s", target, m.Name, err.Error())
	}

	var (
		sout string
		serr string
	)

	if b, err := ioutil.ReadAll(stdout); err == nil {
		sout = string(b)
	}

	if b, err := ioutil.ReadAll(stderr); err == nil {
		serr = string(b)
	}

	log.Printf("Running migration %s...", m.Name)
	cmd.Wait()

	os.Remove(target + filepath.Base(tmpf.Name()))

	return sout, serr, nil
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
			migration, _ := Load(f.Name())
			migrations = append(migrations, migration)
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
