package confparse

// #include "confparse.h"
import (
	"C"
)

type ConfigFile struct {
	filename string
	entries *ConfigEntry
	next *ConfigFile
}

type ConfigEntry struct {
	// variable and section line numbers
	var_linenum int
	sect_linenum int

	// variable name and optional data section
	varname string
	vardata string

	// child entries
	entries *ConfigEntry

	// next entry at same level
	next *ConfigEntry
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
