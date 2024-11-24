package genome

func Process() ([]string, error) {
	lines, err := readGenomeFile()
	if err != nil {
		return nil, err
	}

	// for l := range lines {
	// 	if strings.CheckRepo(l) {
	// git.DownloadRepo(l)
	// }

	// }

	return lines.Parse(), nil
}
