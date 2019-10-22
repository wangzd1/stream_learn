package main

// tutorial01.c
// Code based on a tutorial at http://dranger.com/ffmpeg/tutorial01.html

// A small sample program that shows how to use libavformat and libavcodec to
// read video from a file.
//
// Use
//
// gcc -o tutorial01 tutorial01.c -lavformat -lavcodec -lswscale -lz
//
// to build (assuming libavformat and libavcodec are correctly installed
// your system).
//
// Run using
//
// tutorial01 myvideofile.mpg
//
// to write the first five frames from "myvideofile.mpg" to disk in PPM
// format.
import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetFrame(index int) *bytes.Buffer {
	filename := "test.mp4"
	width := 2752
	height := 2208
	// cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", strconv.Itoa(index), "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")

	buf := new(bytes.Buffer)

	cmd.Stdout = buf

	if cmd.Run() != nil {
		panic("could not generate frame")
	}

	return buf
}

func main() {
	buffer := GetFrame(1)
	fmt.Print(buffer.Len())
}
