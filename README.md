# Last Window ROM Toolkit

A Go package to work with the files inside the ROM for Last Window: The Secret of Cape West.

## Usage

You can run this through `go` with `go run src/lastwindow-tk/main.go`

Or you can build the package with `go build -o lastwindow-tk src/lastwindow-tk/main.go` and run the compiled binary. 

## Supported Files

- Pack (.pack) - these are custom `zlib` archives that can be extracted to a folder using `lastwindow-tk packfile <input.pack> <output dir>`

## Todo

- All the others, mainly:
    - .bra - animation files
    - .bpg - image format (this one doesn't seem custom, but the only public format out there is from 2014 - 4 years after the game came out. Yes I have tried that format anyway and it does not work; surprise)
