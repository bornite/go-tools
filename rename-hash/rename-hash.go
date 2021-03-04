package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	dir, _ := os.Getwd()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var fnames []string
	for _, file := range files {
		if !file.IsDir() && !strings.Contains(file.Name(), "rename-hash") {
			fnames = append(fnames, file.Name())
		}
	}

	for _, fname := range fnames {
		fmt.Println("Processing " + fname + " " + time.Now().String() + " (" + strconv.FormatInt(time.Now().Unix(), 10) + ")")
		if err := os.Rename(fname, renameHash(fname)); err != nil {
			fmt.Println(err)
		}
	}
}

func renameHash(fname string) string {
	hash := sha256.Sum256([]byte(fname + "_SALT_nvnfjkasd" + strconv.FormatInt(time.Now().Unix(), 10)))
	return filepath.Base(fname[:len(fname)-len(filepath.Ext(fname))]) + "_" + hex.EncodeToString(hash[:]) + path.Ext(fname)
}
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	dir, _ := os.Getwd()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var fnames []string
	for _, file := range files {
		if !file.IsDir() && !strings.Contains(file.Name(), "renamehash") {
			fnames = append(fnames, file.Name())
		}
	}

	for _, fname := range fnames {
		fmt.Println("Processing " + fname + " " + time.Now().String() + " (" + strconv.FormatInt(time.Now().Unix(), 10) + ")")
		if err := os.Rename(fname, renameHash(fname)); err != nil {
			fmt.Println(err)
		}
	}
}

func renameHash(fname string) string {
	hash := sha256.Sum256([]byte(fname + "_SALT_nvnfjkasd" + strconv.FormatInt(time.Now().Unix(), 10)))
	return filepath.Base(fname[:len(fname)-len(filepath.Ext(fname))]) + "_" + hex.EncodeToString(hash[:]) + path.Ext(fname)
}
