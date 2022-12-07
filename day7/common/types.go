package common

type Command struct {
	Text   []string
	Output []string
}

type Directory struct {
	parent         *Directory
	Size           int
	Name           string
	SubDirectories map[string]*Directory
	Files          []*File
}

type File struct {
	Size int
	Name string
}

func Touch(name string, size int) *File {
	file := new(File)
	file.Name = name
	file.Size = size
	return file
}

func Exec(text []string, output []string) Command {
	return Command{Text: text, Output: output}
}

func Mkdir(name string, parent *Directory) *Directory {
	dir := new(Directory)
	dir.Name = name
	dir.parent = parent
	dir.SubDirectories = map[string]*Directory{}
	dir.Size = 0
	return dir
}

func (dir *Directory) AddFile(newFile *File) {
	dir.Files = append(dir.Files, newFile)
}

// recalcelate all sizes, return root size
func (dir *Directory) RecalculateSizes() int {
	size := 0
	for _, subdir := range dir.SubDirectories {
		size += subdir.RecalculateSizes()
	}

	for _, file := range dir.Files {
		size += file.Size
	}

	dir.Size = size
	return size
}

// get size of directory
func (dir *Directory) GetSize() int {
	size := 0
	for _, subdir := range dir.SubDirectories {
		if subdir.Size == 0 {
			// size not set yet
			size += subdir.GetSize()
		} else {
			// assuming no changes since last calculated size
			size += subdir.Size
		}

	}

	for _, file := range dir.Files {
		size += file.Size
	}

	dir.Size = size
	return size
}
