package main

import (
	"os"
	"os/exec"
)

func main() {
	//filename := "./video/sample.mp4"
	//cmdname := "ffmpeg"
	//转换视频格式
	//cmd := exec.Command(cmdname, "-i", filename, "-c ", "copy", "video/sample.mov")
	//
	//cmd := exec.Command(cmdname, "-ss", "00:00:17","-t", "00:01:30" ,"-accurate_seek","-i", filename, "-c" , "copy" , "./video/corp.mp4")
	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}

	//cmd  := exec.Command(cmdname, "-i", filename,  "-r","1" ,"-f","image2", "video/image-%3d.jpeg" )
	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}
	//截图第一张
	//cmd  := exec.Command(cmdname, "-i", filename,  "-frames:v","1" ,"-f","image2", "video/image-demo.jpeg" )
	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}
	//裁剪视频
	//cmd := exec.Command(cmdname, "-i" ,filename, "-vf", "scale=640:480", "video/sample_640_480.mp4")
	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}
    //获取视频的信息
	//cmd := exec.Command("ffprobe", "-print_format" ,"json", "-show_streams", "-i", "video/sample_640_480.mp4")
	//cmd.Stdout = os.Stdout
	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}
	//图片转视频

	//cmd := exec.Command("ffmpeg", "-f", "image2","-i", "video/1224352843_640x360.jpg","video/1224352843_640x360.mp4")
	//cmd.Stdout = os.Stdout
	//if err := cmd.Run(); err != nil {
	//	panic(err)
	//}
    //图片转gif
	cmd := exec.Command("ffmpeg",  "-t", "3", "-ss", "00:00:01","-i", "video/2766118275.mp4", "video/sample.gif")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
