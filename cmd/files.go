package cmd

import (
	"embed"
	"os"
	"path"
	"path/filepath"
)

func getAllFilenames(fs embed.FS, dir string) (out []string, err error) {
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}

func extractFile(fp string, fs embed.FS, workingDir string) error {
	fileContent, err := fs.ReadFile(fp)
	if err != nil {
		return err
	}
	filename := filepath.Join(workingDir, fp)
	err = createIfNotExist(filename)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, fileContent, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func createIfNotExist(fp string) error {
	dir, _ := path.Split(fp)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = os.Create(fp)
	if err != nil {
		return err
	}
	return nil
}

func extractAll(fs embed.FS, dir string, workingDir string) error {
	filenames, err := getAllFilenames(fs, dir)
	if err != nil {
		return err
	}
	for _, filename := range filenames {
		err = extractFile(filename, fs, workingDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func ExpandTemporarily(tempUnpackPath string, files embed.FS) {
	_ = extractAll(files, ".", tempUnpackPath)
}
