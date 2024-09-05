package main

import (
	"fmt"
)

// 调用GetConfigData，获取yml文件中的配置信息
// res 用户名
//var res = /

// 用户名
//var username = res.Locals

func main() {
	fmt.Println("username")
}

//func main() {
//	watcher, err := fsnotify.NewWatcher()
//	if err != nil {
//		panic(err)
//	}
//	defer watcher.Close()
//
//	err = watcher.Add("your_directory_path")
//	if err != nil {
//		panic(err)
//	}
//
//	config := &ssh.ClientConfig{
//		User: "your_username",
//		Auth: []ssh.AuthMethod{
//			ssh.Password("your_password"),
//		},
//		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//	}
//
//	sshClient, err := ssh.Dial("tcp", "remote_host:port", config)
//	if err != nil {
//		panic("Failed to dial: " + err.Error())
//	}
//
//	sftpClient, err := sftp.NewClient(sshClient)
//	if err != nil {
//		panic("Failed to create sftp client: " + err.Error())
//	}
//	defer sftpClient.Close()
//
//	for {
//		select {
//		case event := <-watcher.Events:
//			if event.Op&fsnotify.Create == fsnotify.Create {
//				if filepath.Dir(event.Name) == "your_directory_path" {
//					file, err := os.Open(event.Name)
//					if err != nil {
//						fmt.Println("Error opening file:", err)
//						continue
//					}
//					defer file.Close()
//
//					dstFile, err := sftpClient.Create(filepath.Join("remote_directory_path", filepath.Base(event.Name)))
//					if err != nil {
//						fmt.Println("Error creating file on remote:", err)
//						continue
//					}
//					defer dstFile.Close()
//
//					if _, err := io.Copy(dstFile, file); err != nil {
//						fmt.Println("Error uploading file:", err)
//						continue
//					}
//
//					fmt.Printf("Uploaded %s to remote\n", event.Name)
//				}
//			}
//		case err := <-watcher.Errors:
//			fmt.Println("Error:", err)
//		}
//	}
//}
