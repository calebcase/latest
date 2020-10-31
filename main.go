package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
)

func cannot(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	prefix, pathMajor, ok := module.SplitPathVersion(os.Args[1])
	if !ok {
		panic(errors.New("Failed to parse path."))
	}

	major := module.PathMajorPrefix(pathMajor)

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://" + prefix,
	})
	cannot(err)

	tagrefs, err := r.Tags()
	cannot(err)

	latest := ""
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		v := semver.Canonical(t.Name().Short())

		if semver.Compare(latest, v) < 0 {
			if major == "" {
				if strings.HasPrefix(v, "v0") || strings.HasPrefix(v, "v1") {
					latest = v
				}
			} else if strings.HasPrefix(v, major) {
				latest = v
			}
		}

		return nil
	})
	cannot(err)

	fmt.Println(latest)
}
