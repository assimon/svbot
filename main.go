package main

import (
	"fmt"
	"github.com/assimon/svbot/internal/config"
	"github.com/assimon/svbot/telegram"
)

func main() {
	fmt.Println("   _____      _           _   \n  / ____|    | |         | |  \n | (_____   _| |__   ___ | |_ \n  \\___ \\ \\ / / '_ \\ / _ \\| __|\n  ____) \\ V /| |_) | (_) | |_ \n |_____/ \\_/ |_.__/ \\___/ \\__|\n                              \n                              ")
	fmt.Println(fmt.Sprintf("Svbot version(%s) Powered by %s %s \n", config.GetAppVersion(), "assimon", "https://github.com/assimon/svbot"))
	telegram.Start()
}
