# Dependency helper

## Parameters
* `-s` output dependencies as salt code
* `-d` output dependency diagram
* `-o` output directory for installer files 

## VSCode
### SSH Remote
#### .vscode/launch.json
```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Connect to server",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "${workspaceFolder}",
            "port": 2222,
            "host": "127.0.0.1"
        }
    ]
}
```
#### SSH Config
```
Host vagrant2222-Secure
  HostName localhost
  User vagrant
  Port 2222
  IdentityFile C:/Users/Alexander/.ssh/20170926_vagrant_private_key
  StrictHostKeyChecking no
  
Host vagrant2200-Secure
  HostName localhost
  User vagrant
  Port 2200
  IdentityFile C:/Users/Alexander/.ssh/20170926_vagrant_private_key
  StrictHostKeyChecking no
```
### Troubleshooting

#### Cannot run test suite
```
  Failed to start test suite
  fork/exec /vagrant/salt/salt.test: permission denied
```