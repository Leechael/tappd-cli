package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Dstack-TEE/dstack/sdk/go/tappd"
	"github.com/urfave/cli/v2"
)

var (
	Version   string
	BuildTime string
)

func quote(ctx *cli.Context) error {
	var input []byte
	var err error

	// Check if input is coming from stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped in
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read from stdin: %w", err)
		}
	} else if inputFile := ctx.String("input"); inputFile != "" {
		// Read from input file
		input, err = os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("failed to read input file: %w", err)
		}
	} else {
		cli.ShowCommandHelp(ctx, "quote")
		return fmt.Errorf("input is required either from file (-i flag) or stdin")
	}

	// Initialize client
	var client *tappd.TappdClient
	if endpoint := ctx.String("endpoint"); endpoint != "" {
		client = tappd.NewTappdClient(tappd.WithEndpoint(endpoint))
	} else {
		client = tappd.NewTappdClient()
	}

	tdxQuoteResp, err := client.TdxQuote(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to get TDX quote: %w", err)
	}

	// Handle output
	if outputFile := ctx.String("output"); outputFile != "" {
		err = os.WriteFile(outputFile, []byte(tdxQuoteResp.Quote), 0644)
		if err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
	} else {
		fmt.Printf("%s\n", tdxQuoteResp.Quote)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "tappd-cli",
		Usage: "Command line interface for Tappd service",
		Commands: []*cli.Command{
			{
				Name:  "quote",
				Usage: "Get TDX quote from the TEE environment",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Usage:   "Input file path",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Output file path (default: stdout)",
					},
					&cli.StringFlag{
						Name:    "endpoint",
						Aliases: []string{"e"},
						Usage:   "Tappd server endpoint (e.g., http://localhost:8080)",
					},
				},
				Action: quote,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
