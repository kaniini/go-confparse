package confparse

// #include "confparse.h"
import (
	"C"
)

type ConfigFile struct {
	Filename string
	Entries *ConfigEntry
	Next *ConfigFile
}

type ConfigEntry struct {
	// variable and section line numbers
	VarLineNum int
	SectLineNum int

	// variable name and optional data section
	VarName string
	VarData string

	// child entries
	Entries *ConfigEntry

	// next entry at same level
	Next *ConfigEntry
}

func LoadConfigFile (filename string) *ConfigFile {
	return WalkCTree(C.config_file_load(C.CString(filename)))
}

// Walk roots
func WalkCTree(cfptr *C.ConfigFile) *ConfigFile {
	var entries *ConfigEntry = nil
	if cfptr.entries != nil {
		entries = WalkCETree(cfptr.entries)
	}

	var child *ConfigFile = nil
	if cfptr.next != nil {
		child = WalkCTree(cfptr.next)
	}

	return &ConfigFile {
		C.GoString(cfptr.filename),
		entries,
		child,
	}
}

// Walk entries
func WalkCETree(ceptr *C.ConfigEntry) *ConfigEntry {
	var entries *ConfigEntry = nil
	var next *ConfigEntry = nil
	var vardata string = "<undefined>"

	if ceptr.vardata != nil {
		vardata = C.GoString(ceptr.vardata);
	}

	if ceptr.entries != nil {
		entries = WalkCETree(ceptr.entries)
	}

	if ceptr.next != nil {
		next = WalkCETree(ceptr.next)
	}

	return &ConfigEntry {
		int(ceptr.varlinenum),
		int(ceptr.sectlinenum),
		C.GoString(ceptr.varname),
		vardata,
		entries,
		next,
	}
}
