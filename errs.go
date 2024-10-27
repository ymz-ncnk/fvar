package fvar

import "errors"

// ErrSubFolder happens when the folder contains subfolder.
var ErrSubfolder = errors.New("folder should not contain subfolder")
