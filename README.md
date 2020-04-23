# Voodio (Private Media Server)
Voodio is a Simple Private Media Server based on your local Movie Directories. It just a simple program doing tracking on your "Movies" folder, create an index and showing the movies as a web UI to play streamly through the browser.

## Background
I always wanted to watch my old movie collection that saved on my external hardisk or on my PC drive. But unfortunately I'am too lazy to open my computer and starting to crawl and watch those movies. I always wonder if I could see those movies as a webpage, click the detail and play, like I'm watching on the Netfl*x, but the movie is based on my collection. zSo this Voodio is the result.

## Installation

### Have FFMPEG installed on your OS

Visit [FFmpeg Official Download](https://www.ffmpeg.org/download.html page to install based on your current OS. FFmpeg is available for **Windows**, **Mac** and **Linux**

### Using Go Get

    // install it via go get
    go get github.com/slaveofcode/voodio


    // run the binary with path to the parent of video directories
    ./voodio -path /path/to/videos/dir

If the configuration and steps above is complete, you can heads up to http://localhost:8080 on your browser to start watching. By default the server will spawn 2 service on port `:8080` for Web UI and `:8181` for Server API.

### Screenshot
<img src="https://raw.github.com/slaveofcode/voodio/master/assets/home.jpg" align="center" />
<img src="https://raw.github.com/slaveofcode/voodio/master/assets/detail.png" align="center" />
<img src="https://raw.github.com/slaveofcode/voodio/master/assets/play.png" align="center" />

# LICENSE
MIT

Copyright 2020 Aditya Kresna Permana

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
