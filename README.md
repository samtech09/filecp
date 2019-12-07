# filecp
Utility to copy files from different locations to specific folder with hierarchy by auto creating folders of path.

- get input files from text file
- create files with full folder hierarchy at destination folder
- work with absolute paths as well as with relative paths

## Syntax
```
filecp -f [infile] -o [destfolder]

-f string
     in-file containing filepaths to be copied
-o string
     Base out folder where to copy files
```

### Example
```
filecp -f files2copy.txt -o some/folder/
```

### Notes
It is tool that i created to help myself to copy numerous files updated by developers to some other location.

I use `git diff branch1..branch2 --name-only > updated-files.txt` to genrate list of files for it's input.
