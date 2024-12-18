# TrailVidStitch #

Concatenate video files from a trailcam into 1 hour long videos. Videos totalling less than 15 minutes are added to the previous video. Uses ffmpeg's concat demuxer for stitching videos without re-encoding them.

Must have go and ffmpeg installed.

Only works on linux (only tested on arch).

Go get this:

>$ go get github.com/adamringeisen/trailvidstitch

Install to /usr/local/bin/

>$ sudo make install

Then go to the directory where you have all your videos (preferably not your memory card) and run

>$ trailvidstitch

It will make some txt files and then will concatenate the videos. This will take a few minutes. You can delete the txt files afterwards.

Gmail me @ ringeisen if you actually use this.
