package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
    // Check if the required number of command line arguments is provided.
    if len(os.Args) < 3 {
        // Print an error message to standard error and exit with a non-zero status code.
        fmt.Fprintln(os.Stderr, "Not enough arguments")
        os.Exit(-1)
    }

    // Get the old and new strings from command line arguments.
    old, new := os.Args[1], os.Args[2]

    // Create a new Scanner to read from standard input.
    scan := bufio.NewScanner(os.Stdin)

    // Loop through each line of input.
    for scan.Scan() {
        // Split the line into substrings using the old string as the separator.
        s := strings.Split(scan.Text(), old)

        // Join the substrings back together using the new string as the separator.
        t := strings.Join(s, new)

        // Print the modified line to standard output.
        fmt.Println(t)
    }
}