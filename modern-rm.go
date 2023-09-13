package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func isUnixLike() bool {
	return runtime.GOOS == "linux" || runtime.GOOS == "darwin"
}

func isRipAvailable() bool {
	_, err := exec.LookPath("rip")
	return err == nil
}

func rmRemove(cmd *cobra.Command, args []string) {
	fmt.Println("Warning: Your Command calls 'rm'. Deleted files can't be recovered.")
	rm_path, _ := exec.LookPath("rm")
	rmCmd := exec.Command(rm_path, args...)
	rmCmd.Stdout = os.Stdout
	rmCmd.Stderr = os.Stderr
	_ = rmCmd.Run()
}

func ripRemove(cmd *cobra.Command, args []string) {
	rip_path, _ := exec.LookPath("rip")
	if cmd.Flag("interactive").Changed {
		for _, filename := range args {
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				if cmd.Flag("force").Changed {
					continue
				}
				fmt.Printf("Error: %s does not exist.\n", filename)
				os.Exit(1)
			}
			confirm := ""
			fmt.Printf("Are you sure you want to delete '%s'?  ", filename)
			fmt.Scanln(&confirm)
			if strings.ToLower(confirm) != "y" {
				os.Exit(0)
			}
			ripCmd := exec.Command(rip_path, filename)
			ripCmd.Stdout = os.Stdout
			ripCmd.Stderr = os.Stderr
			_ = ripCmd.Run()
			if cmd.Flag("verbose").Changed {
				fmt.Printf("Removed %s\n", filename)
			}
		}
	} else {
		var filenamesToRemove []string
		for _, filename := range args {
			if _, err := os.Stat(filename); os.IsNotExist(err) && cmd.Flag("force").Value.String() != "true" {
				fmt.Printf("Error: %s does not exist.\n", filename)
				os.Exit(1)
			}
			filenamesToRemove = append(filenamesToRemove, filename)
		}
		ripCmd := exec.Command(rip_path, args...)
		ripCmd.Stdout = os.Stdout
		ripCmd.Stderr = os.Stderr
		_ = ripCmd.Run()
		if cmd.Flag("verbose").Changed {
			for _, filename := range filenamesToRemove {
				fmt.Printf("Removed %s\n", filename)
			}
		}
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "modern-rm [files]",
		Short: "A modern rm command replacing the default rm command.",
		Long:  "🗑️  modern-rm\n🔒 Safely delete files with the option to recover them using a sleek CLI interface 💻\n💯 Fully compatible with `rm` and built on `rip`.",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("undo").Changed {
				rip_path, _ := exec.LookPath("rip")
				ripCmd := exec.Command(rip_path, "-u")
				ripCmd.Stdout = os.Stdout
				ripCmd.Stderr = os.Stderr
				_ = ripCmd.Run()
				return
			}
			if len(args) == 0 {
				cmd.Help()
				return
			}

			if !isUnixLike() {
				fmt.Println("Error: This command is only available on Unix-like systems.")
				os.Exit(1)
			}

			if !isRipAvailable() {
				fmt.Println("Error: 'rip' command is not available. Please install it and try again.")
				fmt.Println("You can install it at 'https://github.com/nivekuil/rip'.")
				os.Exit(1)
			}

			if cmd.Flag("once").Changed && len(args) > 3 {
				var confirm string
				fmt.Println("Are you sure you want to delete these files? ")
				fmt.Scanln(&confirm)
				if strings.ToLower(confirm) != "y" {
					os.Exit(0)
				}
			}

			if cmd.Flag("force").Changed {
				cmd.Flags().Set("interactive", "false")
			}

			isRmRequired := cmd.Flag("overwrite").Changed ||
				cmd.Flag("undelete").Changed ||
				cmd.Flag("same-fs").Changed ||
				cmd.Flag("directory").Changed
			if isRmRequired {
				rmRemove(cmd, args)
				return
			}

			ripRemove(cmd, args)
		},
	}

	// Store the original help function
	originalHelpFunc := rootCmd.HelpFunc()

	// Set the custom help function using an inline function
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, strs []string) {
		originalHelpFunc(cmd, strs)
		fmt.Println("\nWritten by YeonGyu Kim (public.kim.yeon.gyu@gmail.com)")
		fmt.Println("- https://github.com/code-yeongyu")
	})

	rootCmd.Flags().BoolP("force", "f", false, "🚫 Ignore nonexistent files and arguments, never prompt")
	rootCmd.Flags().BoolP("interactive", "i", false, "❓ Prompt before every removal")
	rootCmd.Flags().BoolP("once", "I", false, "❗ Prompt once before removing more than three files, or when removing recursively. ")
	rootCmd.Flags().BoolP("directory", "d", false, "📁 Remove directories (Invokes original rm).")
	rootCmd.Flags().BoolP("recursive", "r", false, "Remove directories and their contents recursively")
	rootCmd.Flags().BoolP("overwrite", "P", false, "📝 Overwrite regular files before deleting them (Invokes original rm).")
	rootCmd.Flags().BoolP("undelete", "W", false, "🔄 Attempt to undelete the named files (Invokes original rm).")
	rootCmd.Flags().BoolP("same-fs", "x", false, "📌 Stay on the same filesystem (Invokes original rm).")
	rootCmd.Flags().BoolP("verbose", "v", false, "📊 Display detailed information about the deletion process.")
	rootCmd.Flags().BoolP("undo", "u", false, "🔙 Undo the last deletion. ")
	rootCmd.Flags().BoolP("help", "h", false, "📖 Show help. ")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
