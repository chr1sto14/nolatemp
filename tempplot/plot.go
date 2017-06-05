package tempplot

import (
	"bytes"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"io"
	"time"
)

func makePoints(tss []time.Time, temps []float64) plotter.XYs {
	pts := make(plotter.XYs, len(temps))
	for i := range pts {
		pts[i].X = float64(tss[i].Unix())
		pts[i].Y = temps[i]
	}
	return pts
}

func MakePlot(tss []time.Time, intemps []float64, outtemps []float64) (img bytes.Buffer, err error) {
	p, err := plot.New()
	if err != nil {
		return
	}

	xticks := plot.TimeTicks{Format: "2006-01\n15:04"}

	p.Title.Text = "NOLA Temp"
	p.X.Label.Text = "Time"
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = "Temperature"

	err = plotutil.AddLinePoints(p,
		"Inside", makePoints(tss, intemps),
		"Outside", makePoints(tss, outtemps))
	if err != nil {
		return
	}

	w1 := io.MultiWriter(&img)

	w2, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		return
	}
	w2.WriteTo(w1)
	return

}
