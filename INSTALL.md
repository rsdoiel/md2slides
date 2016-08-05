
# Installation

*md2slides* is a command line program run from a shell like Bash. You can find compiled
version in the [releases](https://github.com/rsdoiel/md2slides/releases/latest) 
in the Github repository in a zip file named *md2slides-binary-release.zip*. Inside
the zip file look for the directory that matches your computer and copy that someplace
defined in your path (e.g. $HOME/bin). 

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (both ARM6 and ARM7)

## Mac OS X

1. Go to [github.com/rsdoiel/md2slides/releases/latest](https://github.com/rsdoiel/md2slides/releases/latest)
2. Click on the green "md2slides-binary-release.zip" link and download
3. Open a finder window and find the downloaded file and unzip it (e.g. md2slides-binary-release.zip)
4. Look in the unziped folder and find dist/macosx-amd64/md2slides
5. Drag (or copy) the *md2slides* to a "bin" directory in your path
6. Open and "Terminal" and run `md2slides -h`

## Windows

1. Go to [github.com/rsdoiel/md2slides/releases/latest](https://github.com/rsdoiel/md2slides/releases/latest)
2. Click on the green "md2slides-binary-release.zip" link and download
3. Open the file manager find the downloaded file and unzip it (e.g. md2slides-binary-release.zip)
4. Look in the unziped folder and find dist/windows-amd64/md2slides.exe
5. Drag (or copy) the *md2slides.exe* to a "bin" directory in your path
6. Open Bash and and run `md2slides -h`

## Linux

1. Go to [github.com/rsdoiel/md2slides/releases/latest](https://github.com/rsdoiel/md2slides/releases/latest)
2. Click on the green "md2slides-binary-release.zip" link and download
3. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/md2slides-binary-release.zip)
4. In the unziped directory and find for dist/linux-amd64/md2slides
5. copy the *md2slides* to a "bin" directory (e.g. cp ~/Downloads/md2slides-binary-release/dist/linux-amd64/md2slides ~/bin/)
6. From the shell prompt run `md2slides -h`

## Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Go to [github.com/rsdoiel/md2slides/releases/latest](https://github.com/rsdoiel/md2slides/releases/latest)
2. Click on the green "md2slides-binary-release.zip" link and download
3. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/md2slides-binary-release.zip)
4. In the unziped directory and find for dist/raspberrypi-arm7/md2slides
5. copy the *md2slides* to a "bin" directory (e.g. cp ~/Downloads/md2slides-binary-release/dist/raspberrypi-arm7/md2slides ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
6. From the shell prompt run `md2slides -h`

