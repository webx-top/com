package com

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileIsCompleted(t *testing.T) {
	debugFileIsCompleted = true
	testFile := `./testdata/fileappend/file.txt`
	MkdirAll(filepath.Dir(testFile), os.ModePerm)
	file, err := os.Create(testFile)
	assert.NoError(t, err)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			//time.Sleep(time.Millisecond * 600)
			file.WriteString(fmt.Sprintf("ABDFFDAFEEFFGEGFAFEFEFAFEAFEFAFEAFGEGAEGAGEGAEGEGA-%d\n", i))
		}
		file.Close()
	}()
	fp, err := os.Open(testFile)
	assert.NoError(t, err)
	defer fp.Close()
	ok, err := FileIsCompleted(fp, time.Now())
	assert.NoError(t, err)
	assert.True(t, ok)

	fp.Seek(-5, io.SeekEnd)
	b := make([]byte, 5)
	_, err = fp.Read(b)
	assert.NoError(t, err)
	assert.Equal(t, "-999\n", string(b))

	wg.Wait()
}

func TestBaseFileName(t *testing.T) {
	r := BaseFileName(`abc/dd.txt`)
	assert.Equal(t, `dd.txt`, r)
	r = BaseFileName(`abc\dd.txt`)
	assert.Equal(t, `dd.txt`, r)
	r = BaseFileName(`abc\dd.txt/`)
	assert.Equal(t, `dd.txt`, r)
	r = BaseFileName(`/`)
	assert.Equal(t, ``, r)
}
