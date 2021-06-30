package zip

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
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

func ZipMultiFile(zipfilename string ,filenames []string) error{
	var (
		zipHandle *os.File
		err error
		zf *zip.Writer
	)
	if path.Ext(zipfilename) != ".zip" {
		zipfilename = zipfilename +".zip"
	}
	zipHandle, err = os.OpenFile(zipfilename,os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		return err
	}
	defer zipHandle.Close()
	zf = zip.NewWriter(zipHandle)
	defer zf.Close()
	for _, filename := range filenames {
		content,err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		basename := path.Base(filename)
		w, err := zf.Create(basename)
		if err != nil {
			return err
		}
		_, err = w.Write(content)
		if err != nil {
			return err
		}
	}
	return nil
}

func UnzipFile(filename string, dstpath string) error{
	var (
		zf *zip.ReadCloser
		err error
	)
	zf,err = zip.OpenReader(filename)
	if err != nil {
		return err
	}
	if f , err := os.Stat(dstpath) ; err != nil || !f.IsDir(){
		if !f.IsDir() {
			return errWrongType
		}
		return err
	}
	if strings.LastIndex(dstpath, "/") != len(dstpath)- 1 {
		dstpath = dstpath + "/"
	}
	for _, file := range zf.File {
		fd, err := file.Open()
		if err != nil {
			return err
		}
		f, err := os.OpenFile(dstpath + file.Name, os.O_CREATE|os.O_RDWR, 0664)
		if err != nil {
			return err
		}
		_,err = io.Copy(f,fd)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return nil
}