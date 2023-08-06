package utils

import (
	"fmt"
	"os/exec"
)

// 没有测试
func RunCmd(imageFilePath string, videoFilePath string) error {
	cmd := exec.Command("./app/lib/ffmpeg.exe",
		"-i", videoFilePath,
		"-r", "1",
		"-vframes", "1",
		"-q:v", "2",
		"-f", "image2",
		imageFilePath,
	)
	if err := cmd.Run(); err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
