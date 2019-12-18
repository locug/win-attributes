package wina

// Archive returns if a file at a path has the archive bit set
func Archive(fp string) (bool, error) {
	return isArchive(fp)
}

// Hidden returns if a file at a path has the hidden bit set
func Hidden(fp string) (bool, error) {
	return isHidden(fp)
}
