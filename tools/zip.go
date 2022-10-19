package tools

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Zip(zipPath string, paths ...string) error {
	// create zip file
	if err := os.MkdirAll(filepath.Dir(zipPath), os.ModePerm); err != nil {
		return err
	}
	archive, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer archive.Close()

	// new zip writer
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	// traverse the file or directory
	for _, srcPath := range paths {
		// remove the trailing path separator if path is a directory
		srcPath = strings.TrimSuffix(srcPath, string(os.PathSeparator))

		// visit all the files or directories in the tree
		err = filepath.Walk(srcPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// create a local file header
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			// set compression
			header.Method = zip.Deflate

			// set relative path of a file as the header name
			header.Name, err = filepath.Rel(filepath.Dir(srcPath), path)
			if err != nil {
				return err
			}
			if info.IsDir() {
				header.Name += string(os.PathSeparator)
			}

			// create writer for the file header and save content of the file
			headerWriter, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(headerWriter, f)
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// 这个被我魔改过了，不能处理嵌套文件夹
func Unzip(zipPath, dstDir string) (file_name []string) {
	// open zip file
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		println("Unzip make something wrong")
		fmt.Printf("err: %v\n", err)
	}
	defer reader.Close()
	for _, file := range reader.File {
		name := unzipFile(file, dstDir)
		file_name = append(file_name, name)
	}
	return file_name
}

func unzipFile(file *zip.File, dstDir string) string {
	// create the directory of file
	filePath := path.Join(dstDir, file.Name)

	// open the file
	rc, err := file.Open()
	if err != nil {
		println("unzipFile make something wrong")
		fmt.Printf("err: %v\n", err)
	}
	defer rc.Close()

	// create the file
	w, err := os.Create(filePath)
	if err != nil {
		println("unzipFile make something wrong")
		fmt.Printf("err: %v\n", err)
	}
	defer w.Close()

	// save the decompressed file content
	_, err = io.Copy(w, rc)
	if err != nil {
		println("unzipFile make something wrong")
		fmt.Printf("err: %v\n", err)
	}
	return file.FileInfo().Name()
}
