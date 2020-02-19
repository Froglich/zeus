# Zeus
This is a very simple application for launching (and restarting in case of faliure) multiple daemons, intended for docker containers. I wrote this entirely for my own use, put putting it out here in case anyone else finds it usefull.

## How do I use it?
Write a JSON array containing one object per process you want Zeus to run. Each object should contain values for **name** (string), **cmd** (the command to execute as a string), **args** (string array with command line arguments), **restart** (boolean, if the process should be restarted when/if it closes). Example below:

```json
[
    {
        "name": "Apache",
        "cmd": "apachectl",
        "args": ["-D", "FOREGROUND"],
        "restart": true
    },{
        "name": "Shibboleth",
        "cmd": "sudo",
        "args": ["-u", "_shibd", "shibd", "-F"],
        "restart": true
    }
]
```

## Installing
Install with `go get github.com/Froglich/zeus` .
