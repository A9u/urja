# urja ![](https://d29fhpw069ctt2.cloudfront.net/icon/image/49635/preview.svg)
A general purpose Go wrapper over Mac's pmset command

### Installation
```sh
go get github.com/A9u/urja
```

### Usage
```go
    // returns status of batteries and UPSs.
    output, err := urja.GetBatteryStatus()

    // run custom command, i.e, pass args to pmset
    output, err := urja.RunCustomCommand("-g", "therm")
    fmt.Println(output.String())

    // run man command on pmset
    output, err := urja.GetHelp()
    fmt.Println(output.String())
```

### Functions available

+ `GetBatteryStatus()` returns string and error and is a wrapper for `pmset -g ps`
+ `RunCustomCommand(args ...string)` returns bytes.Buffer and error. Variable number of arguments can be passed.
+ `GetHelp()` returns bytes.Buffer and error. It is a wrapper for `man pmset`

### Contributing
Feature requests are always welcome with accompanying PR(if possible 😉). Please file an issue if you are facing any challenges. Labels used are _bug/question/enhancement_

If there's anything you'd like to chat about, please feel free to reach me via [twitter](https://twitter.com/alien_billi)!

### [License](https://github.com/A9u/urja/blob/main/LICENSE)
Copyright (c) 2021 Anusha Bhat
Licensed under the MIT license.

