# Go Talks

This repository holds the Go talks I have given. To run these talks locally, make sure that [Go][1] is installed, then clone this repository and execute the following commands

	go get -u -v golang.org/x/tools/cmd/present
    cd [TALK_DIR] # CD to directory of the talk
    present -notes

Now open your browser at `127.0.0.1:[PORT]` with the port given by `present` to view the presentation.

## License
Released under the [MIT license](LICENSE.md).

[1]: http://golang.org/ "Go Language"
