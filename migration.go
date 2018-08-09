package pop

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

// MigrationCreate writes contents for a given migration in normalized files
func MigrationCreate(path, name, ext string, up, down []byte) error {
	g := genny.New()
	n := time.Now().UTC()
	s := n.Format("20060102150405")

	upf := filepath.Join(path, (fmt.Sprintf("%s_%s.up.%s", s, name, ext)))
	g.File(genny.NewFile(upf, strings.NewReader(string(up))))

	downf := filepath.Join(path, (fmt.Sprintf("%s_%s.down.%s", s, name, ext)))
	g.File(genny.NewFile(downf, strings.NewReader(string(down))))

	r := genny.WetRunner(context.Background())
	r.With(g)

	return r.Run()
}

// MigrateUp is deprecated, and will be removed in a future version. Use FileMigrator#Up instead.
func (c *Connection) MigrateUp(path string) error {
	warningMsg := "Connection#MigrateUp is deprecated, and will be removed in a future version. Use FileMigrator#Up instead."
	_, file, no, ok := runtime.Caller(1)
	if ok {
		warningMsg = fmt.Sprintf("%s Called from %s:%d", warningMsg, file, no)
	}
	fmt.Println(warningMsg)

	mig, err := NewFileMigrator(path, c)
	if err != nil {
		return errors.WithStack(err)
	}
	return mig.Up()
}

// MigrateDown is deprecated, and will be removed in a future version. Use FileMigrator#Down instead.
func (c *Connection) MigrateDown(path string, step int) error {
	warningMsg := "Connection#MigrateDown is deprecated, and will be removed in a future version. Use FileMigrator#Down instead."
	_, file, no, ok := runtime.Caller(1)
	if ok {
		warningMsg = fmt.Sprintf("%s Called from %s:%d", warningMsg, file, no)
	}
	fmt.Println(warningMsg)

	mig, err := NewFileMigrator(path, c)
	if err != nil {
		return errors.WithStack(err)
	}
	return mig.Down(step)
}

// MigrateStatus is deprecated, and will be removed in a future version. Use FileMigrator#Status instead.
func (c *Connection) MigrateStatus(path string) error {
	warningMsg := "Connection#MigrateStatus is deprecated, and will be removed in a future version. Use FileMigrator#Status instead."
	_, file, no, ok := runtime.Caller(1)
	if ok {
		warningMsg = fmt.Sprintf("%s Called from %s:%d", warningMsg, file, no)
	}
	fmt.Println(warningMsg)

	mig, err := NewFileMigrator(path, c)
	if err != nil {
		return errors.WithStack(err)
	}
	return mig.Status()
}

// MigrateReset is deprecated, and will be removed in a future version. Use FileMigrator#Reset instead.
func (c *Connection) MigrateReset(path string) error {
	warningMsg := "Connection#MigrateReset is deprecated, and will be removed in a future version. Use FileMigrator#Reset instead."
	_, file, no, ok := runtime.Caller(1)
	if ok {
		warningMsg = fmt.Sprintf("%s Called from %s:%d", warningMsg, file, no)
	}
	fmt.Println(warningMsg)

	mig, err := NewFileMigrator(path, c)
	if err != nil {
		return errors.WithStack(err)
	}
	return mig.Reset()
}
