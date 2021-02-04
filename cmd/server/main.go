package main

import (
	_ "github.com/sirupsen/logrus"
	_ "github.com/spf13/cobra"
)

func main() {
	rootCmd.Execute()
}
