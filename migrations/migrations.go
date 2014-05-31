package migrations

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Migration struct {
	Timestamp string
	Name      string
	Path      string
}

func (m *Migration) RunAgainst(target string) (string, string, error) {
	if _, err := os.Stat(m.Path); err != nil {
		return "", "", fmt.Errorf("Error opening the file %s: %s", m.Path, err.Error())
	}

	// Run the migration
	cmd := exec.Command(m.Path, target)

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

	log.Printf("Running migration %s...", m.Name)

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

	return sout, serr, cmd.Wait()
}

func UpdateMark(targetPath string, current *Migration) error {
	abspath := filepath.Join(targetPath, ".karver")

	if err := ioutil.WriteFile(abspath, []byte(current.Timestamp), 0755); err != nil {
		return err
	}

	return nil
}

func Run(pending []*Migration, targetPath string) error {
	if len(pending) == 0 {
		return nil
	}

	log.Printf("Karving %s...", targetPath)

	for _, current := range pending {
		sout, serr, err := current.RunAgainst(targetPath)

		if err != nil {
			log.Fatalf("Error running migration %s: %s", current.Name, err.Error())
			return err
		}

		if sout != "" {
			fmt.Println(strings.Trim(sout, "\n"))
		}
		if serr != "" {
			fmt.Println(strings.Trim(serr, "\n"))
		}

		if err = UpdateMark(targetPath, current); err != nil {
			log.Fatalf("Error updating migration mark: %s", err.Error())
			return err
		}
	}

	log.Printf("%s has been karved. :D", targetPath)

	return nil
}

func NeedsToRun(timestamp string, current *Migration) bool {
	if current != nil {
		return timestamp < current.Timestamp
	}

	return false
}

func AbsMigrationsPath(migrationsPath string) (string, error) {
	if migrationsPath == "" {
		dir := filepath.Dir(os.Args[0])
		migrationsPath = filepath.Join(dir, "migrations")
	}

	return filepath.Abs(migrationsPath)
}

func List(migrationsPath string) ([]*Migration, error) {
	files, err := ioutil.ReadDir(migrationsPath)

	if err != nil {
		return nil, err
	}

	migrations := make([]*Migration, 0, len(files))
	r, _ := regexp.Compile("^([0-9]{14})_.*")

	for _, f := range files {
		m := r.FindStringSubmatch(f.Name())
		if len(m) > 0 && !f.IsDir() {
			migration := Load(filepath.Join(migrationsPath, f.Name()))
			migrations = append(migrations, migration)
		}
	}

	return migrations, nil
}

// If no karver mark in the target we assume that we have to
// run all the migrations.
func CurrentTimestamp(targetPath string) (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(targetPath, ".karver"))

	if err != nil {
		if os.IsNotExist(err) {
			return "0", nil
		} else {
			return "", err
		}
	}

	return string(data[:14]), nil
}

func Last(list []*Migration, timestamp string) *Migration {
	var previous *Migration

	for _, migration := range list {
		if NeedsToRun(timestamp, migration) {
			return previous
		}

		previous = migration
	}

	return previous
}

func Pending(list []*Migration, timestamp string) []*Migration {
	pending := make([]*Migration, 0, len(list))

	for _, migration := range list {
		if NeedsToRun(timestamp, migration) {
			pending = append(pending, migration)
		}
	}

	return pending
}

func Create(title string, migrationsPath string) (*Migration, error) {
	name := strings.Replace(title, " ", "_", -1)
	timestamp := time.Now().Local().Format("20060102150405")
	migrationName := timestamp + "_" + name + ".sh"
	abspath := filepath.Join(migrationsPath, migrationName)
	template := `#!/bin/sh
# Welcome to the Karver migration template
#
# $1 is the migration's target path
echo "Karving from $0..."
`

	err := ioutil.WriteFile(abspath, []byte(template), 0755)
	if err != nil {
		return nil, err
	}

	migration := Load(abspath)

	return migration, err
}

// path is an absolute path to the migration
func Load(path string) *Migration {
	name := filepath.Base(path)
	r, _ := regexp.Compile("^([0-9]{14})_.*")
	m := r.FindStringSubmatch(name)
	migration := &Migration{m[1], name, path}
	return migration
}
