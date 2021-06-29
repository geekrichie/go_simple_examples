package zip

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
)
var errWrongType  = errors.New("func param filename require a filename but got a folder")

func ZipOneFile(filename string)error {
	var (
		f *os.File
		fi os.FileInfo
		err error
		w io.Writer
	)
	fi , err = os.Stat(filename)
	if err != nil || fi.IsDir() {
		if fi.IsDir(){
			return errWrongType
		}
		return err
	}
	f , err = os.Open(filename)
	if err !=  nil {
		return err
	}
	defer f.Close()
	basename  := path.Base(filename)
	suffix := path.Ext(filename)
	prefixName := basename[0:len(basename)-len(suffix)]
	zipFileName  := prefixName + ".zip"
	zipHandle, err := os.OpenFile(zipFileName, os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		return err
	}
	defer zipHandle.Close()
	var zf = zip.NewWriter(zipHandle)
	defer zf.Close()
	w, err = zf.CreateHeader(&zip.FileHeader{
		Name: basename,
	})
	if err != nil {
		return err
	}
	filecontent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	_,err = w.Write(filecontent)
	if err != nil {
		return err
	}

	return nil
}