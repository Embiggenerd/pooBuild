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

var Run = func(inputFIle, outputFile string) {
	fmt.Println("^^^^")
	fmt.Println(args)
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

	// logger.API = logger.CLIAPI

	osArgs := os.Args[1:]
	// heapFile := ""
	// traceFile := ""
	// cpuprofileFile := ""
	// isRunningService := false
	// sendPings := false
	// isWatch := false
	// isWatchForever := false
	// isServe := false

	// Do an initial scan over the argument list
	argsEnd := 0
	for _, arg := range osArgs {
		switch {
		// Show help if a common help flag is provided
		case arg == "-h", arg == "-help", arg == "--help", arg == "/?":
			fmt.Println(helpText(TerminalColors))
			os.Exit(0)

		// Special-case the version flag here
		case arg == "--version":
			fmt.Printf("%s\n", ".0.01")
			os.Exit(0)
		case strings.HasPrefix(arg, "--input="):
			inputFile = arg[len("--input="):]
		default:
			fmt.Println("***default***")
			// case arg == "-s", arg == "--source", arg == "-src", arg == "/?":
			// 	fmt.Println(helpText(nil))
			// 	os.Exit(0)

			// case strings.HasPrefix(arg, "--heap="), strings.HasPrefix(arg, "--heap="):
			// 	heapFile = arg[len("--heap="):]

			// case strings.HasPrefix(arg, "--trace="):
			// 	traceFile = arg[len("--trace="):]

			// case strings.HasPrefix(arg, "--timing"):
			// 	// This is a hidden flag because it's only intended for debugging esbuild
			// 	// itself. The output is not documented and not stable.
			// 	api_helpers.UseTimer = true

			// case strings.HasPrefix(arg, "--cpuprofile="):
			// 	cpuprofileFile = arg[len("--cpuprofile="):]

			// // This flag turns the process into a long-running service that uses
			// // message passing with the host process over stdin/stdout
			// case strings.HasPrefix(arg, "--service="):
			// 	hostVersion := arg[len("--service="):]
			// 	isRunningService = true

			// 	// Validate the host's version number to make sure esbuild was installed
			// 	// correctly. This check was added because some people have reported
			// 	// errors that appear to indicate an incorrect installation.
			// 	if hostVersion != esbuildVersion {
			// 		logger.PrintErrorToStderr(osArgs,
			// 			fmt.Sprintf("Cannot start service: Host version %q does not match binary version %q",
			// 				hostVersion, esbuildVersion))
			// 		os.Exit(1)
			// 	}

			// case strings.HasPrefix(arg, "--ping"):
			// 	sendPings = true

			// default:
			// 	// Some people want to be able to run esbuild's watch mode such that it
			// 	// never exits. However, esbuild ends watch mode when stdin is closed
			// 	// because stdin is always closed when the parent process terminates, so
			// 	// ending watch mode when stdin is closed is a good way to avoid
			// 	// accidentally creating esbuild processes that live forever.
			// 	//
			// 	// Explicitly allow processes that live forever with "--watch=forever".
			// 	// This may be a reasonable thing to do in a short-lived VM where all
			// 	// processes in the VM are only started once and then the VM is killed
			// 	// when the processes are no longer needed.
			// 	if arg == "--watch" || arg == "--watch=true" {
			// 		isWatch = true
			// 	} else if arg == "--watch=forever" {
			// 		arg = "--watch"
			// 		isWatch = true
			// 		isWatchForever = true
			// 	} else if arg == "--serve" ||
			// 		strings.HasPrefix(arg, "--serve=") ||
			// 		strings.HasPrefix(arg, "--servedir=") ||
			// 		strings.HasPrefix(arg, "--serve-fallback=") {
			// 		isServe = true
			// 	}

			// Strip any arguments that were handled above
			osArgs[argsEnd] = arg
			argsEnd++
		}
	}
	osArgs = osArgs[:argsEnd]

	// Print help text when there are no arguments

	if len(osArgs) == 0 {
		fmt.Println("defaults are gonna be sued")
		os.Exit(0)
	}

	// Print help text when there are no arguments
	// isStdinTTY := logger.GetTerminalInfo(os.Stdin).IsTTY
	// if len(osArgs) == 0 && isStdinTTY {
	// 	logger.PrintText(os.Stdout, logger.LevelSilent, osArgs, helpText)
	// 	os.Exit(0)
	// }

	// Capture the defer statements below so the "done" message comes last
	exitCode := 1
	fmt.Printf("os args: %s", osArgs)
	// func() {
	// 	// To view a CPU trace, use "go tool trace [file]". Note that the trace
	// 	// viewer doesn't work under Windows Subsystem for Linux for some reason.
	// 	if traceFile != "" {
	// 		if done := createTraceFile(osArgs, traceFile); done == nil {
	// 			return
	// 		} else {
	// 			defer done()
	// 		}
	// 	}

	// 	// To view a heap trace, use "go tool pprof [file]" and type "top". You can
	// 	// also drop it into https://speedscope.app and use the "left heavy" or
	// 	// "sandwich" view modes.
	// 	if heapFile != "" {
	// 		if done := createHeapFile(osArgs, heapFile); done == nil {
	// 			return
	// 		} else {
	// 			defer done()
	// 		}
	// 	}

	// 	// To view a CPU profile, drop the file into https://speedscope.app.
	// 	// Note: Running the CPU profiler doesn't work under Windows subsystem for
	// 	// Linux. The profiler has to be built for native Windows and run using the
	// 	// command prompt instead.
	// 	if cpuprofileFile != "" {
	// 		if done := createCpuprofileFile(osArgs, cpuprofileFile); done == nil {
	// 			return
	// 		} else {
	// 			defer done()
	// 		}
	// 	}

	// 	if cpuprofileFile != "" {
	// 		// The CPU profiler in Go only runs at 100 Hz, which is far too slow to
	// 		// return useful information for esbuild, since it's so fast. Let's keep
	// 		// running for 30 seconds straight, which should give us 3,000 samples.
	// 		seconds := 30.0
	// 		start := time.Now()
	// 		for time.Since(start).Seconds() < seconds {
	// 			exitCode = cli.Run(osArgs)
	// 		}
	// 	} else {
	// 		if !isWatch && !isServe {
	// 			// If this is not a long-running process and there is at most a single
	// 			// entry point, then disable the GC since we're just going to allocate
	// 			// a bunch of memory and then exit anyway. This speedup is not
	// 			// insignificant. We don't do this when there are multiple entry points
	// 			// since otherwise esbuild could unnecessarily use much more memory
	// 			// than it might otherwise need to process many entry points.
	// 			nonFlagCount := 0
	// 			for _, arg := range osArgs {
	// 				if !strings.HasPrefix(arg, "-") {
	// 					nonFlagCount++
	// 				}
	// 			}
	// 			if nonFlagCount <= 1 {
	// 				debug.SetGCPercent(-1)
	// 			}
	// 		} else if isServe && isServeUnsupported() {
	// 			// The development server isn't supported on WebAssembly, so we will
	// 			// immediately call "os.Exit(1)" below, which will call "process.exit(1)"
	// 			// in node. However, node has a bug/feature where any pending calls to
	// 			// "fs.read(process.stdin.fd)" hold up "process.exit()" without seemingly
	// 			// any way to stop this from happening. So to avoid this bug/feature,
	// 			// we explicitly avoid listening to stdin in this case (when we know
	// 			// that we are about to exit due to an invalid flag).
	// 		} else if !isStdinTTY && !isWatchForever {
	// 			// If stdin isn't a TTY, watch stdin and abort in case it is closed.
	// 			// This is necessary when the esbuild binary executable is invoked via
	// 			// the Erlang VM, which doesn't provide a way to exit a child process.
	// 			// See: https://github.com/brunch/brunch/issues/920.
	// 			//
	// 			// We don't do this when stdin is a TTY because that interferes with
	// 			// the Unix background job system. If we read from stdin then Ctrl+Z
	// 			// to move the process to the background will incorrectly cause the
	// 			// job to stop. See: https://github.com/brunch/brunch/issues/998.
	// 			go func() {
	// 				// This just discards information from stdin because we don't use
	// 				// it and we can avoid unnecessarily allocating space for it
	// 				buffer := make([]byte, 512)
	// 				for {
	// 					_, err := os.Stdin.Read(buffer)
	// 					if err != nil {
	// 						if options := logger.OutputOptionsForArgs(osArgs); options.LogLevel <= logger.LevelInfo {
	// 							if isWatch {
	// 								// Mention why watch mode was stopped to reduce confusion, and
	// 								// call out "--watch=forever" to get the alternative behavior
	// 								logger.PrintTextWithColor(os.Stderr, options.Color, func(colors logger.Colors) string {
	// 									return fmt.Sprintf("%s[watch] stopped automatically because stdin was closed (use \"--watch=forever\" to keep watching even after stdin is closed)%s\n", colors.Dim, colors.Reset)
	// 								})
	// 							} else if isServe {
	// 								logger.PrintTextWithColor(os.Stderr, options.Color, func(colors logger.Colors) string {
	// 									return fmt.Sprintf("%s[serve] stopped automatically because stdin was closed (keep stdin open to continue serving)%s\n", colors.Dim, colors.Reset)
	// 								})
	// 							}
	// 						}

	// 						// Only exit cleanly if stdin was closed cleanly
	// 						if err == io.EOF {
	// 							os.Exit(0)
	// 						} else {
	// 							os.Exit(1)
	// 						}
	// 					}

	// 					// Some people attempt to keep esbuild's watch mode open by piping
	// 					// an infinite stream of data to stdin such as with "< /dev/zero".
	// 					// This will make esbuild spin at 100% CPU. To avoid this, put a
	// 					// small delay after we read some data from stdin.
	// 					time.Sleep(4 * time.Millisecond)
	// 				}
	// 			}()
	// 		}

	// 		exitCode = cli.Run(osArgs)
	// 	}
	// }()
	Run(osArgs)
	os.Exit(exitCode)
}
