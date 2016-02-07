parses the Fide rating lists in text format and converts them to an R data tables

#######################################

installation

#######################################

under Windows 64 bit

1) dowload the repository as a zip
2) unzip
3) parsetxt.exe is in the root directory of the unzipped repository ready for use

on any platform:

1) install the Go language
2) create a workspace directory for Go
3) set the PATH environment variable to the full path of the workspace directory
4) install git and make sure it is in your path
5) open a console ( command prompt ) window and type:

go get github.com/chessstats/parsetxt

6) the executable will be created in the bin directory of the workspace

#######################################

usage:

#######################################

the program should be in a directory containing a subdirectory called 'hist_download'
the 'hist_download' directory should contain Fide standard rating lists in text format with .txt or .TXT extension
running the program will create a directory called 'hist_smartconv' and will create R data tables in this directory
from the rating lists in the 'hist_download' directory