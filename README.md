# go-scratch-pad

## salt
Running tests:
1. Ensure counterfeiter is installed for creating mocks
2. Navigate to our test package
3. Run ginkgo
```
go install github.com/maxbrunsfeld/counterfeiter/v6
cd /vagrant/salt
ginkgo
```

## VSCode SSH extension
**NOTE** that ginkgo tests fail with write permissions immediately after using the SSH extension to connect to the `/vagrant` directory
* Vagrant up
* Remote Explorer > SSH > Open in new window
* Install plugins in ssh env
    * Go `golang.go` (extension ID)
    * Ginkgo Test Explorer `joselitofilho.ginkgotestexplorer` (extension ID)
    * Ginkgo Tools `dlipovetsky.ginkgo-tools` (extension ID) - doesn't work
    * Try to install via the [command line](https://code.visualstudio.com/docs/editor/extension-marketplace#_command-line-extension-management)
* Explorer > Open `/vagrant/`