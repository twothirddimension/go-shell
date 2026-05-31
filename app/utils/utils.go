package utils

import (
	"os"
)

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0111 != 0
}

func CheckFileInDirectory(path string, fileName string) (bool, []os.DirEntry) {
	files, err := os.ReadDir(path)
	var dir []os.DirEntry
	if err != nil {
		return false, dir
	}

	for _, file := range files {
		if file.Name() == fileName && isExecutable(path+"/"+file.Name()) {
			return true, nil
		} else {
			if file.IsDir() {
				dir = append(dir, file)
			}
		}
	}

	return false, dir
}

func getNewDirNames(presenDir string, files []os.DirEntry) []string {
	newPaths := []string{}
	for _, file := range files {
		newPaths = append(newPaths, file.Name()+"/"+presenDir)
	}

	return newPaths
}

func recursivelyCheckFiles(paths []string, fileName string) (bool, string) {
	for i := 0; i < len(paths); i++ {
		isPresent, dir := CheckFileInDirectory(paths[i], fileName)
		if isPresent {
			return true, paths[i] + "/" + fileName
		} else {
			return recursivelyCheckFiles(getNewDirNames(paths[i], dir), fileName)
		}
	}

	return false, ""
}
