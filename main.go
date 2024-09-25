// Copyright (C) 2024 Robert A. Wallis, All Rights Reserved

package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func main() {
	flag.Parse()

	fmt.Println("removing empty directories")
	for _, arg := range flag.Args() {
		if err := delInDir(arg); err != nil {
			fmt.Printf("%s error: %v\n", arg, err)
		}
	}
}

func delInDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			err := delInDir(path.Join(dir, entry.Name()))
			if err != nil {
				return err
			}
		}
	}
	if len(entries) > 0 {
		// one more time to make sure it's empty now
		entries, err = os.ReadDir(dir)
		if err != nil {
			return err
		}
	}
	if len(entries) > 0 {
		return nil
	}
	fmt.Printf("removed %s\n", dir)
	return os.Remove(dir)
}
