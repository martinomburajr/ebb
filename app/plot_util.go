package app

import (
	"fmt"
	"github.com/martinomburajr/ebb/evolution"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var (
	StatsTypeAverage = "StatsTypeAverage"
	StatsTypeBest    = "StatsTypeBest"
)

func plotTopologyForBothIndividualStats(p *plot.Plot, name, statsType string, kind, index int, topologyPlots []evolution.TopologyPlot) PlotDetails {
	for i, pl := range topologyPlots {
		var antPlotXY plotter.XYer
		var proPlotXY plotter.XYer
		kindStr := ""

		switch kind {
		case evolution.IndividualAntagonist:
			kindStr = "Bug"
		case evolution.IndividualProtagonist:
			kindStr = "Test"
		}

		switch statsType {
		case StatsTypeAverage:

			antPlotXY = NewPlotXY(pl.XAxisGen, pl.AvgOfAllAntagonistsInGen)
			proPlotXY = NewPlotXY(pl.XAxisGen, pl.AvgOfAllProtagonistsInGen)
		case StatsTypeBest:
			antPlotXY = NewPlotXY(pl.XAxisGen, pl.TopAntagonistsBestInGen)
			proPlotXY = NewPlotXY(pl.XAxisGen, pl.TopProtagonistsBestInGen)
		}

		err := AddLinePoints(p, i+index, kind, fmt.Sprintf("%s-%s", pl.Topology, kindStr), antPlotXY, proPlotXY)
		if err != nil {
			panic(err)
		}
	}

	return PlotDetails{
		Name: name,
		Plot: p,
	}
}

func plotIndividualTopologyStats(p *plot.Plot, name, statsType string, kind, index int, topologyPlots []evolution.TopologyPlot) PlotDetails {

	for i, pl := range topologyPlots {
		var plotXY plotter.XYer
		kindStr := ""

		switch kind {
		case evolution.IndividualAntagonist:
			kindStr = "Bug"
			switch statsType {
			case StatsTypeAverage:
				plotXY = NewPlotXY(pl.XAxisGen, pl.AvgOfAllAntagonistsInGen)
			case StatsTypeBest:
				plotXY = NewPlotXY(pl.XAxisGen, pl.TopAntagonistsBestInGen)
			}
		case evolution.IndividualProtagonist:
			kindStr = "Test"
			switch statsType {
			case StatsTypeAverage:
				plotXY = NewPlotXY(pl.XAxisGen, pl.AvgOfAllProtagonistsInGen)
			case StatsTypeBest:
				plotXY = NewPlotXY(pl.XAxisGen, pl.TopProtagonistsBestInGen)
			}
		}

		err := AddLinePoints(p, i+index, kind, fmt.Sprintf("%s-%s", pl.Topology, kindStr), plotXY)
		if err != nil {
			panic(err)
		}
	}

	return PlotDetails{
		Name: name,
		Plot: p,
	}
}

// AddLinePoints adds Line and Scatter plotters to a
// plot.  The variadic arguments must be either strings
// or plotter.XYers.  Each plotter.XYer is added to
// the plot using the next color, dashes, and glyph
// shape via the Color, Dashes, and Shape functions.
// If a plotter.XYer is immediately preceeded by
// a string then a legend entry is added to the plot
// using the string as the name.
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddLinePoints(plt *plot.Plot, index, kind int, vs ...interface{}) error {
	var ps []plot.Plotter
	type item struct {
		name  string
		value [2]plot.Thumbnailer
	}
	var items []item
	name := ""

	for _, v := range vs {
		switch t := v.(type) {
		case string:
			name = t

		case plotter.XYer:
			l, s, err := plotter.NewLinePoints(t)
			if err != nil {
				return err
			}

			switch kind {
			case evolution.IndividualAntagonist:
				l.Color = plotutil.Color(index)
				l.Width = vg.Length(2)
				l.Dashes = plotutil.Dashes(3)
				s.Color = plotutil.Color(index)
				s.Shape = plotutil.Shape(3)
			case evolution.IndividualProtagonist:
				l.Color = plotutil.Color(index)
				l.Width = vg.Length(2)
				l.Dashes = plotutil.Dashes(1)
				s.Color = plotutil.Color(index)
				s.Shape = plotutil.Shape(1)
			default:
				l.Color = plotutil.Color(index)
				l.Width = vg.Length(2)
				l.Dashes = plotutil.Dashes(index)
				s.Color = plotutil.Color(index)
				s.Shape = plotutil.Shape(index)
			}

			index++
			ps = append(ps, l, s)
			if name != "" {
				items = append(items, item{name: name, value: [2]plot.Thumbnailer{l, s}})
				name = ""
			}

		default:
			panic(fmt.Sprintf("plotutil: AddLinePoints handles strings and plotter.XYers, got %T", t))
		}
	}

	plt.Add(ps...)
	for _, item := range items {
		v := item.value[:]
		plt.Legend.Add(item.name, v[0], v[1])
	}
	return nil
}

func NewPlotXY(x []int, y []float64) plotter.XYer {
	pts := make(plotter.XYs, len(x))

	for i := range pts {
		pts[i].X = float64(x[i])
		pts[i].Y = y[i]
	}

	return pts
}

func plotSingle(points []evolution.GenerationResult, stat string) plotter.XYs {
	pts := make(plotter.XYs, len(points))

	switch stat {
	case "AllAntagonistAverageFitness":
		for i := range pts {
			pts[i].X = float64(i)
			pts[i].Y = points[i].AllAntagonistAverageFitness
		}
	case "AllProtagonistAverageFitness":
		for i := range pts {
			pts[i].X = float64(i)
			pts[i].Y = points[i].AllProtagonistAverageFitness
		}
	case "BestProtagonist":
		for i := range pts {
			pts[i].X = float64(i)
			pts[i].Y = points[i].BestProtagonist.AverageFitness
		}
	case "BestAntagonist":
		for i := range pts {
			pts[i].X = float64(i)
			pts[i].Y = points[i].BestAntagonist.AverageFitness
		}
	default:
		panic("invalid plot")
	}

	return pts
}

//
//func plotMetric(points []evolution.TopologicalResult, stat string) plotter.XYs {
//	stratLen := len(points[0].BestAntagonist.Strategy)
//	pts := make(plotter.XYs, stratLen)
//
//	switch stat {
//	case "AllAntagonistAverageFitness":
//		for i := range pts {
//			pts[i].X = float64(i)
//			pts[i].Y = points[i].
//		}
//	case "AllProtagonistAverageFitness":
//		for i := range pts {
//			pts[i].X = float64(i)
//			pts[i].Y = points[i].AllProtagonistAverageFitness
//		}
//	case "BestProtagonist":
//		for i := range pts {
//			pts[i].X = float64(i)
//			pts[i].Y = points[i].BestProtagonist.AverageFitness
//		}
//	case "BestAntagonist":
//		for i := range pts {
//			pts[i].X = float64(i)
//			pts[i].Y = points[0].FinalBestAntagonistOfAllRuns.Strategy[0] .BestAntagonist.Strategy
//		}
//	default:
//		panic("invalid plot")
//	}
//
//	return pts
//}

// valuesPlotter uses data as a form of random distribution to enable the plotting of histograms.
func valuesPlotter(data []float64) plotter.Values {
	v := make(plotter.Values, len(data))
	for i := range v {
		v[i] = data[i]
	}

	return v
}

type HistogramPlot struct {
	Title string
	XAxis string
	YAxis string

	Output string

	Values []float64

	ShouldNormalize bool
	Length          int
	Height          int
}

func (h *HistogramPlot) Plot() *plot.Plot {
	values := valuesPlotter(h.Values)

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = h.Title

	// Create a histogram of our values drawn
	// from the standard normal.
	hist, err := plotter.NewHist(values, len(evolution.AllStrategies))
	if err != nil {
		panic(err)
	}

	if h.ShouldNormalize {
		hist.Normalize(1)
	}
	p.Add(hist)

	// Save the plot to a PNG file.
	//if err := p.Save(vg.Length(h.Length)*vg.Millimeter, vg.Length(h.Height)*vg.Millimeter, h.Output); err != nil {
	//	panic(err)
	//}

	return p
}

type Dimension struct {
	Data  []float64
	Label string
}

func (d *Dimension) ToValues() plotter.Values {
	p := make(plotter.Values, len(d.Data))
	for i := range p {
		p[i] = d.Data[i]
	}

	return p
}

type BoxPlot struct {
	Title string
	XAxis string
	YAxis string

	Values []Dimension
	Width  float64

	ShouldNormalize bool
	Length          uint
	Height          uint
}

func (b *BoxPlot) Plot() *plot.Plot {
	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = b.Title
	p.Y.Label.Text = b.YAxis

	for i := range b.Values {
		box, err := plotter.NewBoxPlot(vg.Length(b.Width)*vg.Millimeter, float64(i), b.Values[0].ToValues())
		if err != nil {
			panic(err)
		}

		p.Add(box)
		p.NominalX(b.Values[i].Label)
	}

	return p
}
