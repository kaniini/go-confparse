/*
 * Copyright (C) 2005-2008 William Pitcock, et al.
 * Rights to this code are as documented in LICENSE.
 *
 * Config file parser.
 */

#ifndef CONFPARSE_H
#define CONFPARSE_H

#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <errno.h>
#include <string.h>

typedef struct configfile_ ConfigFile;
typedef struct configentry_ ConfigEntry;

struct configfile_
{
	char *filename;
	ConfigEntry *entries;
	ConfigFile *next;
	int curline;
	char *mem;
};

struct configentry_
{
	ConfigFile *fileptr;

	int varlinenum;
	char *varname;
	char *vardata;
	int sectlinenum;/* line containing closing brace */

	ConfigEntry *entries;
	ConfigEntry *prevlevel;
	ConfigEntry *next;
};

/* confp.c */
extern void config_file_free(ConfigFile *cfptr);
extern ConfigFile *config_file_load(const char *filename);

#endif
