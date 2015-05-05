# easyssh

## Description

This fork of [hypersleep/easyssh](https://github.com/hypersleep/easyssh) doesn't use private key files e.g(/.ssh/id_rsa)
instead of providing the file path to 'id_rsa', you provide the key as a string, so you can run it anywhere, check the examples


## WARNING
Your private key will be included in the binary after compiling it, so this executable should be as private as your id_rsa file


## So easy to use!

[Run a command on remote server and get STDOUT output](https://github.com/pyed/easyssh/blob/master/example/run.go)

[Upload a file to remote server](https://github.com/pyed/easyssh/blob/master/example/scp.go)
