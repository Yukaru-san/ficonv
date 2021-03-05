# ficonv
A tool to convert files into images representing the file's bytes.

# Features
The Tool is currently supporting the following features:
- Convert Files to an Image
- Convert that image back into a file

# Why would you need this?
Why wouldn't you?
<br>
<br>
Here is a converted version of my [ageGUI](https://github.com/Yukaru-san/ageGUI):<br>
https://files.jojii.de/preview/HlFkGWOruPy3TOvIO36NI3ghX<br>
(It's too big for Github)

# How does it work?
The program takes chunks of 4 bytes from the original file to represent 1 pixel (r,g,b,a).<br>
This way the resulting image is always about the same size as the original file.

# Installation
If you want to compile the program yourself:
<br>
```git clone https://github.com/Yukaru-san/FI-Converter``` <br>
```cd FI-Converter``` <br>
```go build -o ficonv``` <br>

You can also head over to the [release tab](https://github.com/Yukaru-san/ageGUI/releases/tag/v1.1) and download a precompiled version.

# Usage
```
usage: ficonv [<flags>] <in> <out>

A tool to parse files as an image and vice versa

Flags:
      --help     Show context-sensitive help (also try --help-long and
                 --help-man).
  -t, --trim     Trims trailing NULL bytes. Useful for txt files, can damage
                 other types. Only used together with "reverse"
  -r, --reverse  Reverses an image to a file. By default the given file will be
                 converted to an image.

Args:
  <in>   The file you want to convert
  <out>  The output file
```
 
Example: <br>
```ficonv main.go output.png```&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;- creates the image<br>
```ficonv output.png main.go -r -t``` - parses the image back into a file
