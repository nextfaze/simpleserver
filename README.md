#Simple Server

The most simple static http server in go to quickly serve files out of a static directory

(although `cd /home/somedir && python -m SimpleHTTPServer` is pretty simple to :wink: )

##Build and Install

You'll need your `GOBIN` to be set (e.g. `export GOBIN=.`)


To Build:

    export GOPATH=`pwd`; go get github.com/codegangsta/cli
    export GOBIN=.; go install simpleserver.go

To Run:

   $GOBIN/simpleserver
   
