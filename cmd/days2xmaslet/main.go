package main

import (
	"os"
	"strconv"

	"github.com/spetix/days2xmasleft/internal/blockletapi"
	"github.com/spetix/days2xmasleft/internal/dateutil"
	"github.com/spf13/cobra"
)

func main() {
	var renderOptions blockletapi.RenderOptions

	var proto string
	rootCmd := &cobra.Command{
		Use:   "dats2xmasleft",
		Short: "dats2xmasleft",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

			b := blockletapi.New(&renderOptions, proto)
			b.Execute()
		},
	}

	flags := rootCmd.Flags()
	flags.StringVarP(&proto, "proto", "p", "raw", "proto")
	flags.DurationVarP(&renderOptions.Unit, "unit", "u", dateutil.Day, "unit")
	lbl := os.Getenv("label")
	if lbl == "" {
		lbl = "🎄"
	}
	flags.StringVarP(&renderOptions.Label, "label", "l", lbl, "label")
	flags.StringVarP(&renderOptions.Format, "format", "f", "text", "format")
	clr := os.Getenv("color")
	if clr == "" {
		clr = "#ff0000"
	}
	flags.StringVarP(&renderOptions.ForegroundColor, "color", "c", clr, "foreground color")
	bgclr := os.Getenv("background")
	if bgclr == "" {
		bgclr = "#000000"
	}
	flags.StringVarP(&renderOptions.BackgroundColor, "background", "b", "#000000", "background color")

	buttonClicked, err := strconv.Atoi(os.Getenv("BLOCK_BUTTON"))
	if err != nil {
		buttonClicked = 0
	}
	flags.Uint8VarP(&renderOptions.ButtonClicked, "button", "t", uint8(buttonClicked), "button clicked")

	baseLogFolder := os.Getenv("BASE_LOG_FOLDER")
	if baseLogFolder == "" {
		baseLogFolder = "."
	}
	flags.StringVarP(&renderOptions.BaseLogFolder, "logFolder", "d", baseLogFolder, "log folder")

	rootCmd.Execute()
}
