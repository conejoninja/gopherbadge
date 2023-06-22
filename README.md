# Gopher Badge

TinyGo powered badge in the shape of your favorite burrowing rodent, the mighty Go Gopher.

https://gopherbadge.com

# How to install

- Install TinyGo using the instructions from https://tinygo.org

- Clone this repo

- Change directories into the directory with the repo

- Connect your Gopherbadge to your computer using a USB cable


- Run this command to compile and flash the code to your Gopher badge:

If you are running Mac or Linux, or have make installed you can run the following:

```
make flash
```

otherwise run tinygo directly

```
tinygo flash -target gopher-badge .
```

Note: if you get a `permision denied` error; please, consult this [page](https://tinygo.org/docs/guides/tinygo-flash-errors/) for possible solution. You many need to restart the computer; afterward to get the group to stick.

- To display a conference logo on your badge, use one of the following targets (depending on GC for Europe, Ireland, UK, or US):
```
make flash-gopherconeu
make flash-gopherconie
make flash-gopherconuk
make flash-gopherconus
```

- To customize the Gobadge with your own name and information, use the `NAME`, `TITLEA1`, `TITLEA2`, `MARQUEETOP`, `MARQUEEMIDDLE`, `MARQUEEBOTTOM`, `TITLEB1`, `TITLEB2` and `QRTEXT` variables like this:

```
make flash-gopherconeu NAME="@_CONEJO" TITLEA1="Go compiler" TITLEA2="small places"
```

# Add an new logo

- Create an image with a 320x240 pixels size, copy it into `cmd/assets` folder.  
For the moment only jpeg images are supported.  
- In `cmd/main.go` add the path to your file here

```go
const (
tinygoLogo         = "./cmd/assets/tinygo.jpg"
purpleHardwareLogo = "./cmd/assets/purpleHardware.jpg"
gopherconEU        = "./cmd/assets/gopherconeu.jpg"
gopherconIE        = "./cmd/assets/gopherconie.jpg"
gopherconUK        = "./cmd/assets/gopherconuk.jpg"
gopherconUS        = "./cmd/assets/gopherconus.jpg"
wasmio             = "./cmd/assets/wasmio.jpg"
)
```

- Add the corresponding flag to the conf map:

```go
func confs() map[string]string {
    return map[string]string{
        "tinygo":      tinygoLogo,
        "purple":      purpleHardwareLogo,
        "gopherconeu": gopherconEU,
        "gopherconie": gopherconIE,
        "gopherconuk": gopherconUK,
        "gopherconus": gopherconUS,
        "wasmio":      wasmio,
    }
}
```

Add a new target to the Makefile:

```bash
flash-yourconf:
	go run cmd/main.go -conf=flagLogo
	tinygo flash -target gopher-badge .
```

You can run:

```bash
make flash-yourconf
```

It will store the data in `logo.bin` file that will be embedded into the code.

```go
//go:embed logo.bin
var badgeLogo string
```

After the image has been generated, the make command will flash it to the board.


👏 Congratulations! It is now your own customized Gopher Badge.



# Not powering up with battery connected
If your battery is connected and switching your badge to ON doesn't power it up, **disconnect your battery, switch to ON and connect your battery again**. If it doesn't power up, then check the battery charge.