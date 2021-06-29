package  tar

import (
	"archive/tar"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var errWrongSize = errors.New("err: write wrong size")
var errNoFileZip =  errors.New("please add at least one file to zip")
var errFileType = errors.New("wrong file type")


func TarOneFile(filename string) error{
	var f  *os.File
	var err error
	var content []byte
	content,err = ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	basename := path.Base(filename)
	suffix := path.Ext(filename)
	prefixName := basename[0:len(basename)-len(suffix)]
	zipFileName := prefixName + ".tar"
	f,err = os.OpenFile(zipFileName,os.O_CREATE|os.O_RDWR,0644)
	if  err != nil {
		return err
	}
	defer f.Close()
	var zf  = tar.NewWriter(f)
	zf.WriteHeader(&tar.Header{
		Name:basename,
		Size: int64(len(content)),
	})
	var n int
	n, err = zf.Write(content)
	if n != len(content) {
		return errWrongSize
	}
	if err != nil {return err}

	return nil
}

func TarMultiFile(zipFileName string, filenames ...string) error{
	if len(filenames) == 0{
		return errNoFileZip
	}
	if !strings.HasSuffix(zipFileName,".tar") {
		zipFileName = zipFileName + ".tar"
	}
	var (
		f *os.File
		err error
	)
	f, err = os.OpenFile(zipFileName, os.O_CREATE|os.O_RDWR,0644)
	if err != nil {
		return err
	}
	defer f.Close()
	var zf  = tar.NewWriter(f)
	defer zf.Close()
	for _,filename := range filenames {
		basename := path.Base(filename)
		fileContent,err := ioutil.ReadFile(filename)
		if err != nil{
			return err
		}
		zf.WriteHeader(&tar.Header{
			Name: basename,
			Size: int64(len(fileContent)),
		})
		zf.Write(fileContent)
	}
	return nil
}

func UnTarFile(filename string , dirpath string) error{
	var (
		f *os.File
		err error
		hdr *tar.Header
	)
	f, err = os.Open(filename)
	if err != nil {
		return err
	}
	var fi os.FileInfo
	fi, err = os.Stat(dirpath)
	if err != nil  || !fi.IsDir(){
		if !fi.IsDir() {
			return errFileType
		}
		return err
	}
	if strings.LastIndex(dirpath,"/") != len(dirpath)-1{
		dirpath = dirpath + "/"
	}
	var zf = tar.NewReader(f)
	for{
		hdr, err = zf.Next()
		if err == io.EOF {
			break //正常退出
		}
		if err != nil{
			return err
		}
		var temp *os.File
		temp,err = os.OpenFile(dirpath + hdr.Name, os.O_CREATE|os.O_RDWR, 0664)
		if err != nil {
			return err
		}
		_, err := io.Copy(temp,zf)
		if err != nil {
			temp.Close()
			return err
		}
		temp.Close()
	}
	return nil
}