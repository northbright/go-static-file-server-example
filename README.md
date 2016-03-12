# go-static-file-server-example

A golang static file server example.  

[![Build Status](https://travis-ci.org/northbright/go-static-file-server-example.svg?branch=master)](https://travis-ci.org/northbright/go-static-file-server-example)

## Notes
* Make sure that the path passed in `http.FileServer(http.Dir("XX"))` is **ABSOLUTE**.  
* `http.Dir("./static")` works only if `PWD` is the server root dir.
* Use [GetCurrentExecDir](https://godoc.org/github.com/northbright/pathhelper#GetCurrentExecDir) to get absolute path of server root(current executable).

## References
* [Serve homepage and static content from root](http://stackoverflow.com/questions/14086063/serve-homepage-and-static-content-from-root)
