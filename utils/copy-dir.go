package utils

import "os"

func CopyDir(srcDir, destDir string) error {

	//TODO : check if destDir doesn't exist

	err := os.CopyFS(destDir, os.DirFS(srcDir))
	if err != nil {
		return err
	}
	return nil
}
