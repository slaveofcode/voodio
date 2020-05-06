<img src="https://raw.github.com/slaveofcode/voodio/master/assets/Voodio.png" align="right" />
# Voodio (Private Media Server)
Voodio is a Simple Private Media Server based on your local Movie Directories. It just a simple program doing tracking on your "Movies" folder, create an index and showing the movies via Web UI to play streamly through the browser.

## Background
I always wanted to watch my old movie collection that saved on my external hardisk or on my PC drive. But unfortunately I'am too lazy to open my computer, starting to crawl and watch those movies. I always wonder if I could see those movies as a Website, click the detail and play, like I'm watching on the Netfl*x, but the movie is on local Hard Drive. 

## Application Behavior
This application will need extra space like **8-10x** of the played video which extracted from FFmpeg transcoding/transmuxing process of **HLS** files, the space needed could be lower if there a less available resolution generated (like only for 480p or 720o), then you can play it streamly through your local network on your favourite device via **Browser**. The space will be cleand up after the server is turned of (killed), it will be immediatelly deletes all the generated HLS files so you get the space again.

## Installation

### Have FFMPEG installed on your OS

Visit [FFmpeg Official Download](https://www.ffmpeg.org/download.html) page to install based on your current OS. FFmpeg is available for **Windows**, **Mac** and **Linux**. Please use FFmpeg with version below **4.0** or better with version **3.***.

FFmpeg with version >= 4.1 had an issue of immediatelly generate playlist file (**.m3u8**), then do update  periodically at transcoding time, I don't know why that's happen, but I think there's a solution for that, just need a more time to find out.

### Using Precompiled Binary

Go to the prebuilt binary on the [release section here](https://github.com/slaveofcode/voodio/releases), depending on what your OS 

### Using Go Get

    // install it via go get
    go get github.com/slaveofcode/voodio

    // run the binary with path to the parent of video directories from your module binary path
    // options for port and ffmpeg-bin are optional
    ~/go/bin/voodio -port 8080 -tmdb-key <your-tmdb-key> -ffmpeg-bin /usr/bin/ffmpeg -path /path/to/videos/dir 

If the configuration and steps above is complete, you can heads up to http://[your-ip-host-or-localhost]:8080 on your browser to start watching.

### Running Options

- `-path` The full path of the video directory
- `-tmdb-key` API key of TMDB, you can grab one at [Official TMDB API](https://www.themoviedb.org/documentation/api)
- `-port` (optional) The port number for the server to run
- `-ffmpeg-bin` (optional) The path of FFmpeg binary, if you have a different path of FFmpeg

### Screenshot
<img src="https://raw.github.com/slaveofcode/voodio/master/assets/home.png" align="center" />
<img src="https://raw.github.com/slaveofcode/voodio/master/assets/detail.png" align="center" />
<img src="https://raw.github.com/slaveofcode/voodio/master/assets/play.png" align="center" />

# LICENSE
MIT

Copyright 2020 Aditya Kresna Permana

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
