package app

import (
	"fmt"
	"github.com/martinomburajr/ebb/evolog"
	"github.com/martinomburajr/ebb/evolution"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"sync"
)

var (
	AllAntagonistAverageFitness  = "AllAntagonistAverageFitness"
	AllProtagonistAverageFitness = "AllAntagonistAverageFitness"
	BestProtagonist              = "AllAntagonistAverageFitness"
	BestAntagonist               = "AllAntagonistAverageFitness"
)

type topologyStrategyPlot struct {
	topology string
	// YAxis will be frequency auto-generated es
	// XAxis will be strategies.
	XAxis []evolution.Strategy
}

func (s *Simulation) Plot(result evolution.SimulationResult) {

	setPlotter := result.SET.ToPlotterFormat("SET")
	rrPlotter := result.RR.ToPlotterFormat("RR")
	hofPlotter := result.HoF.ToPlotterFormat("HoF")
	krtPlotter := result.KRT.ToPlotterFormat("KRT")

	topologyPlots := []evolution.TopologyPlot{krtPlotter, hofPlotter, rrPlotter, setPlotter}

	antagonistAveragesPlot := s.NewPlotAveragesLine(evolution.IndividualAntagonist, topologyPlots)
	protagonistAveragesPlot := s.NewPlotAveragesLine(evolution.IndividualProtagonist, topologyPlots)
	allAveragesPlot := s.NewPlotAveragesLine(-1, topologyPlots)

	topAntagonistAveragesPlot := s.NewPlotTopIndividualAveragesLine(evolution.IndividualAntagonist, topologyPlots)
	topProtagonistAveragesPlot := s.NewPlotTopIndividualAveragesLine(evolution.IndividualProtagonist, topologyPlots)
	topAllAveragesPlot := s.NewPlotTopIndividualAveragesLine(-1, topologyPlots)

	topAntVProkrt := s.NewPlotTopAntVsProInTopologiesLine("top-AntVsProKRT", krtPlotter)
	topAntVProrr := s.NewPlotTopAntVsProInTopologiesLine("top-AntVsProRR", rrPlotter)
	topAntVProhof := s.NewPlotTopAntVsProInTopologiesLine("top-AntVsProHOF", hofPlotter)
	topAntVProset := s.NewPlotTopAntVsProInTopologiesLine("top-AntVsProSET", setPlotter)

	avgAntVProkrt := s.NewPlotAvgAntVsProInTopologiesLine("avg-AntVsProKRT", krtPlotter)
	avgAntVProrr := s.NewPlotAvgAntVsProInTopologiesLine("avg-AntVsProRR", rrPlotter)
	avgAntVProhof := s.NewPlotAvgAntVsProInTopologiesLine("avg-AntVsProHOF", hofPlotter)
	avgAntVProset := s.NewPlotAvgAntVsProInTopologiesLine("avg-AntVsProSET", setPlotter)

	h := s.Config.Plots.Height
	l := s.Config.Plots.Length
	output := s.Config.GeneratedPlotsOutputPath

	plots := []PlotDetails{
		antagonistAveragesPlot,
		protagonistAveragesPlot,
		allAveragesPlot,
		topAntagonistAveragesPlot,
		topProtagonistAveragesPlot,
		topAllAveragesPlot,
		topAntVProkrt,
		topAntVProrr,
		topAntVProhof,
		topAntVProset,
		avgAntVProkrt,
		avgAntVProrr,
		avgAntVProhof,
		avgAntVProset,
	}

	wg := sync.WaitGroup{}
	for _, pl := range plots {
		wg.Add(1)

		go func(plotDetail PlotDetails) {
			defer wg.Done()

			err := plotDetail.Plot.Save(
				vg.Length(l)*vg.Millimeter,
				vg.Length(h)*vg.Millimeter,
				fmt.Sprintf("%s/%s.png", output, plotDetail.Name))

			if err != nil {
				s.Config.Params.ErrorChan <- err
			}

		}(pl)
	}

	msg := fmt.Sprintf("completed plots")
	s.LogMessage(msg, evolog.LoggerPlot)

	wg.Wait()
}

type PlotDetails struct {
	Name string
	Plot *plot.Plot
}

func DefaultPlot(s *Simulation) *plot.Plot {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.X.Padding = vg.Millimeter * 10
	p.Y.Padding = vg.Millimeter * 10
	p.X.Max = float64(s.Config.Params.GenerationsCount + 10)
	p.Y.Max = 1.2
	p.Y.Min = -1.2
	p.Legend.Top = true
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	p.Legend.Padding = vg.Millimeter * 5

	return p
}

// This plots generational averages across all runs for a given set of topologies using a line graph
func (s *Simulation) NewPlotAveragesLine(kind int, topologyPlots []evolution.TopologyPlot) PlotDetails {
	p := DefaultPlot(s)

	p.Add(plotter.NewGrid())

	switch kind {
	case evolution.IndividualAntagonist:
		p.Title.Text = "Bug Fitness across Generations"
		p.X.Label.Text = "Generations"
		p.Y.Label.Text = "Fitness"

		return plotIndividualTopologyStats(p, "bug-averages", StatsTypeAverage, evolution.IndividualAntagonist, 0, topologyPlots)

	case evolution.IndividualProtagonist:
		p.Title.Text = "Tests Fitness across Generations"
		p.X.Label.Text = "Generations"
		p.Y.Label.Text = "Fitness"

		return plotIndividualTopologyStats(p, "test-averages", StatsTypeAverage, evolution.IndividualProtagonist, 0, topologyPlots)

	default:
		p.Title.Text = "Bug & Tests Fitness Across Generations"
		p.X.Label.Text = "Generations"
		p.Y.Label.Text = "Fitness"

		plotIndividualTopologyStats(p, "bug-averages", StatsTypeAverage, evolution.IndividualAntagonist, 0, topologyPlots)
		plotIndividualTopologyStats(p, "bug-averages", StatsTypeAverage, evolution.IndividualProtagonist, 4, topologyPlots)

		return PlotDetails{
			Name: "bug-and-test-averages",
			Plot: p,
		}
	}
}

// This plots generational averages across all runs for a given set of topologies using a line graph
func (s *Simulation) NewPlotTopIndividualAveragesLine(kind int, topologyPlots []evolution.TopologyPlot) PlotDetails {
	p := DefaultPlot(s)

	p.Add(plotter.NewGrid())

	switch kind {
	case evolution.IndividualAntagonist:
		p.Title.Text = fmt.Sprintf("Bug Averages Across Generations | %s", s.Config.Params.StartIndividual.ToMathematicalString())
		p.X.Label.Text = "Generations"
		p.Y.Label.Text = "Fitness"

		return plotIndividualTopologyStats(p, "top-bug-averages", StatsTypeBest, evolution.IndividualAntagonist, 0, topologyPlots)

	case evolution.IndividualProtagonist:
		p.Title.Text = fmt.Sprintf("Tests Averages Across Generations | %s", s.Config.Params.StartIndividual.ToMathematicalString())
		p.X.Label.Text = "Generations"
		p.Y.Label.Text = "Fitness"

		return plotIndividualTopologyStats(p, "top-test-averages", StatsTypeBest, evolution.IndividualProtagonist, 0, topologyPlots)

	default:
		p.Title.Text =  fmt.Sprintf("Bug & Tests Averages Across Generations | %s", s.Config.Params.StartIndividual.ToMathematicalString())
		p.X.Label.Text = "Generations"
		p.Y.Label.Text = "Fitness"

		plotIndividualTopologyStats(p, "top-bug-averages", StatsTypeBest, evolution.IndividualAntagonist, 0, topologyPlots)
		plotIndividualTopologyStats(p, "top-test-averages", StatsTypeBest, evolution.IndividualProtagonist, 4, topologyPlots)

		return PlotDetails{
			Name: "top-bug-and-test-averages",
			Plot: p,
		}
	}
}



// This plots generational averages across all runs for a given set of topologies using a line graph
func (s *Simulation) NewPlotTopAntVsProInTopologiesLine(fileName string, topologyPlots evolution.TopologyPlot) PlotDetails {
	p := DefaultPlot(s)

	grid := plotter.NewGrid()
	grid.Horizontal.Width.Points()
	grid.Vertical.Width.Dots(200)
	p.Add(grid)

	p.Title.Text = fmt.Sprintf("%s: Top Bug vs Top Test | %s", topologyPlots.Topology, s.Config.Params.StartIndividual.ToMathematicalString())
	p.X.Label.Text = "Generations"
	p.Y.Label.Text = "Fitness"

	plotTopologyForBothIndividualStats(p, "topology-vs", StatsTypeBest, evolution.IndividualAntagonist, 0,
		[]evolution.TopologyPlot{topologyPlots})
	plotTopologyForBothIndividualStats(p, "topology-vsp-1", StatsTypeBest, evolution.IndividualProtagonist, 0,
		[]evolution.TopologyPlot{topologyPlots})

	return PlotDetails{
		Name: fileName,
		Plot: p,
	}
}

// This plots generational averages across all runs for a given set of topologies using a line graph
func (s *Simulation) NewPlotAvgAntVsProInTopologiesLine(fileName string, topologyPlots evolution.TopologyPlot) PlotDetails {
	p := DefaultPlot(s)

	p.Add(plotter.NewGrid())

	p.Title.Text = fmt.Sprintf("%s: Average Bug vs Top Test | %s", topologyPlots.Topology, s.Config.Params.StartIndividual.ToMathematicalString())
	p.X.Label.Text = "Generations"
	p.Y.Label.Text = "Fitness"


	plotTopologyForBothIndividualStats(p, "topology-vs", StatsTypeAverage, evolution.IndividualAntagonist, 0,
		[]evolution.TopologyPlot{topologyPlots})
	plotTopologyForBothIndividualStats(p, "topology-vsp-1", StatsTypeAverage, evolution.IndividualProtagonist, 4,
		[]evolution.TopologyPlot{topologyPlots})

	return PlotDetails{
		Name: fileName,
		Plot: p,
	}
}

// This plots generational averages across all runs for a given set of topologies using a box plot
func (s *Simulation) NewPlotAveragesBox(kind int, plots []evolution.TopologyPlot) *plot.Plot {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	switch kind {
	case evolution.IndividualAntagonist:
	case evolution.IndividualProtagonist:
	default:
		panic("invalid individual kind")
	}

	return p
}

// Plots the strategy histogram for bestIndividuals
func (s *Simulation) NewPlotStrategyHistogram(kind int, plots topologyStrategyPlot) *plot.Plot {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	switch kind {
	case evolution.IndividualAntagonist:
	case evolution.IndividualProtagonist:
	default:
		panic("invalid individual kind")
	}

	return p
}

// PlotStrategyOrder plots the order with which the best (kind_ individual implementes the array) for a given set of topologies
func (s *Simulation) NewPlotStrategyOrder(kind int, plots []topologyStrategyPlot) *plot.Plot {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	switch kind {
	case evolution.IndividualAntagonist:
	case evolution.IndividualProtagonist:
	default:
		panic("invalid individual kind")
	}

	return p
}
