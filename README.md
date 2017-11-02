# joinICS: A small tool to join multiple .ics files into a stream

This is a command-line small tool to merge all `.ics` files within a given set of folders into a stream, which can e.g. be shared and subscribed online.

## Setup

Upon the first usage, the user will be asked to provide the folders containing .ics files to be joined and the output folder.

```
$ ./joinICS
Missing settings
Run joinICS --set <inputfolders> <outputfolder>,
whereas the input folders must be split by a |
$ ./joinICS --set /path/to/a/folder|/path/to/another/folder /path/to/output/folder
```

The settings will be stored in `~/.config/joinICS` and used when running the program after.
