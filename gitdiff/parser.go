	header = strings.TrimPrefix(header, fileHeaderPrefix)
	defaultName, err := parseGitHeaderName(header)
	if err != nil {
		return p.Errorf("git file header: %v", err)
	}
		end, err := parseGitHeaderData(f, line, defaultName)
			return p.Errorf("git file header: %v", err)
	if f.OldName == "" && f.NewName == "" {
		if defaultName == "" {
			return p.Errorf("git file header: missing filename information")
		}
		f.OldName = defaultName
		f.NewName = defaultName
	}

	if (f.NewName == "" && !f.IsDelete) || (f.OldName == "" && !f.IsNew) {
		return p.Errorf("git file header: missing filename information")
	}
