package main
import (
	"subtitlefixer/fixer"
	"path/filepath"
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

func main(){
	app := cli.App{
		Name: "Subtitle fixer",
		Usage: "Fix Persian Subtitles",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "file",
				Usage: "Subtitle path",
				Aliases: []string{"f"} ,
			},

			&cli.StringFlag{
				Name: "output",
				Usage: "Output path",
				Aliases: []string{"o"} ,
			},
		},
		Action: func( c *cli.Context) error{
			inputFileName := filepath.Clean(c.String("file"))
			outputFileName := filepath.Clean(c.String("output"))
			
			if 1 == len(inputFileName) {
				return fmt.Errorf("input filename was wrong")
			}
			
			if 1 == len(outputFileName) {
				outputFileName = inputFileName[:len(inputFileName) - len(filepath.Ext(inputFileName))] + "_fixed.srt"
			}

			return fixer.Fixer(inputFileName, outputFileName)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}