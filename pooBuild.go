package main

import (
	"fmt"
	"os"
	"strings"
)

type Colors struct {
	Reset     string
	Bold      string
	Dim       string
	Underline string

	Red   string
	Green string
	Blue  string

	Cyan    string
	Magenta string
	Yellow  string

	RedBgRed     string
	RedBgWhite   string
	GreenBgGreen string
	GreenBgWhite string
	BlueBgBlue   string
	BlueBgWhite  string

	CyanBgCyan       string
	CyanBgBlack      string
	MagentaBgMagenta string
	MagentaBgBlack   string
	YellowBgYellow   string
	YellowBgBlack    string
}

var helpText = func(colors Colors) string {
	return `
` + colors.Bold + `Usage:` + colors.Reset + `
  	pooBuild [input file] [output file]

	If the CLI is given no flags, it will use defaults.


` + colors.Bold + `Simple options:` + colors.Reset + `
  --input, -in          The input HTML file, default is ./index.html
  --output, -out        The output HTML file, default is ./dist/index.html
`
}

var Run = func(inputFile, outputFile string) {
	fmt.Println("^^^^")
	fmt.Println(inputFile + " " + outputFile)
	fmt.Println("^^^^")
}

var TerminalColors = Colors{
	Reset:     "\033[0m",
	Bold:      "\033[1m",
	Dim:       "\033[37m",
	Underline: "\033[4m",

	Red:   "\033[31m",
	Green: "\033[32m",
	Blue:  "\033[34m",

	Cyan:    "\033[36m",
	Magenta: "\033[35m",
	Yellow:  "\033[33m",

	RedBgRed:     "\033[41;31m",
	RedBgWhite:   "\033[41;97m",
	GreenBgGreen: "\033[42;32m",
	GreenBgWhite: "\033[42;97m",
	BlueBgBlue:   "\033[44;34m",
	BlueBgWhite:  "\033[44;97m",

	CyanBgCyan:       "\033[46;36m",
	CyanBgBlack:      "\033[46;30m",
	MagentaBgMagenta: "\033[45;35m",
	MagentaBgBlack:   "\033[45;30m",
	YellowBgYellow:   "\033[43;33m",
	YellowBgBlack:    "\033[43;30m",
}

func main() {
	inputFile := "./index.html"
	outputFile := "./dist/index.html"

	osArgs := os.Args[1:]

	argsEnd := 0
	for _, arg := range osArgs {
		switch {
		case arg == "-h", arg == "-help", arg == "--help", arg == "/?":
			fmt.Println(helpText(TerminalColors))
			os.Exit(0)
		case arg == "--version":
			fmt.Printf("%s\n", ".0.01")
			os.Exit(0)
		case strings.HasPrefix(arg, "--input="):
			inputFile = arg[len("--input="):]
		case strings.HasPrefix(arg, "-in="):
			inputFile = arg[len("-in="):]
		case strings.HasPrefix(arg, "--output="):
			outputFile = arg[len("--output="):]
		case strings.HasPrefix(arg, "-out="):
			outputFile = arg[len("-out="):]
		default:
			fmt.Printf("Unrecognized input: %s \n\n", arg)
			fmt.Println(helpText(TerminalColors))
			os.Exit(0)
		}
	}

	osArgs = osArgs[:argsEnd]
	// Call defaults if not specified
	if len(osArgs) == 0 {
		fmt.Printf("Defaults will be used (%s, %s)\n", inputFile, outputFile)
	}

	Run(inputFile, outputFile)
	os.Exit(1)
}
