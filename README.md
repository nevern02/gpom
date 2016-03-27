# gpom
gpom is a command-line [Pomodoro](http://pomodorotechnique.com/) timer written in Go. 

## Installation
First, install [Go](https://golang.org/doc/install).

Then, install the package from this repo:
```
go get github.com/nevern02/gpom
```

Finally, make sure your `PATH` includes `$GOPATH/bin`:
```
export PATH=$PATH:/$GOPATH/bin
```

## Usage
Start the timer:
```
gpom start
```

After counting down, `notify-send` will be used to display a popup alert (on supported systems).

## Contributing

1. Fork it ( http://github.com/nevern02/gpom/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
