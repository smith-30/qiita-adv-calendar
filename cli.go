package main

import (
	"flag"
	"fmt"

	"github.com/smith-30/qiita-adv-calendar/domain/service"
	"go.uber.org/zap"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	Name = "qiita-adv-calendar"

	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

var (
	revision = "unknown"
)

// CLI is the command line object
type CLI struct{}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		name  string
		count int
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)

	flags.StringVar(&name, "name", "", "")
	flags.StringVar(&name, "n", "", "(Short)")
	flags.IntVar(&count, "count", 0, "")
	flags.IntVar(&count, "c", 0, "(Short)")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	logger, _ := zap.NewDevelopment()
	s := logger.Sugar()

	s.Infof("revision %s", revision)

	// ready grid.
	cap := count * 25
	ag := service.NewAggregater(cap, s)
	gridUpdateCh := ag.UpdateGrid(cap)

	// fetch grids each calendar.
	ga := service.NewGridAggregater(name, count, s)
	ga.FetchGrids(gridUpdateCh)

	// wait to send grid.
	ga.Wait()
	s.Info("finished FetchGrids.")
	close(gridUpdateCh)
	// wait aggregate.
	ag.Wait()

	fmt.Println(`
		████████╗██╗  ██╗ █████╗ ███╗   ██╗██╗  ██╗███████╗                           
		╚══██╔══╝██║  ██║██╔══██╗████╗  ██║██║ ██╔╝██╔════╝                           
		   ██║   ███████║███████║██╔██╗ ██║█████╔╝ ███████╗                           
		   ██║   ██╔══██║██╔══██║██║╚██╗██║██╔═██╗ ╚════██║                           
		   ██║   ██║  ██║██║  ██║██║ ╚████║██║  ██╗███████║                           
		   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝                           
																					  
		███████╗ ██████╗ ██████╗                                                      
		██╔════╝██╔═══██╗██╔══██╗                                                     
		█████╗  ██║   ██║██████╔╝                                                     
		██╔══╝  ██║   ██║██╔══██╗                                                     
		██║     ╚██████╔╝██║  ██║                                                     
		╚═╝      ╚═════╝ ╚═╝  ╚═╝                                                     
																					  
		███████╗██╗  ██╗ █████╗ ██████╗ ██╗███╗   ██╗ ██████╗                         
		██╔════╝██║  ██║██╔══██╗██╔══██╗██║████╗  ██║██╔════╝                         
		███████╗███████║███████║██████╔╝██║██╔██╗ ██║██║  ███╗                        
		╚════██║██╔══██║██╔══██║██╔══██╗██║██║╚██╗██║██║   ██║                        
		███████║██║  ██║██║  ██║██║  ██║██║██║ ╚████║╚██████╔╝                        
		╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝                         
																					  
		██╗  ██╗███╗   ██╗ ██████╗ ██╗    ██╗██╗     ███████╗██████╗  ██████╗ ███████╗
		██║ ██╔╝████╗  ██║██╔═══██╗██║    ██║██║     ██╔════╝██╔══██╗██╔════╝ ██╔════╝
		█████╔╝ ██╔██╗ ██║██║   ██║██║ █╗ ██║██║     █████╗  ██║  ██║██║  ███╗█████╗  
		██╔═██╗ ██║╚██╗██║██║   ██║██║███╗██║██║     ██╔══╝  ██║  ██║██║   ██║██╔══╝  
		██║  ██╗██║ ╚████║╚██████╔╝╚███╔███╔╝███████╗███████╗██████╔╝╚██████╔╝███████╗
		╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝  ╚══╝╚══╝ ╚══════╝╚══════╝╚═════╝  ╚═════╝ ╚══════╝
																																																											
`)

	return ExitCodeOK
}
