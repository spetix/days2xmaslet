package popup

import (
	"log/slog"
	"time"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type Popup struct {
}

func New() *Popup {
	return &Popup{}
}

func (p *Popup) Execute() {

	slog.Info(("PopUP enabled"))
	X, err := xgb.NewConn()
	if err != nil {
		slog.Warn(err.Error())
		return
	}
	defer X.Close()
	setup := xproto.Setup(X)

	screen := setup.DefaultScreen(X)
	win, err := xproto.NewWindowId(X)
	if err != nil {
		slog.Warn(err.Error())
		return
	}

	screenWidth := int(setup.Roots[0].WidthInPixels)
	screenHeight := int(setup.Roots[0].HeightInPixels)

	// Calculate the x and y coordinates to center the window
	x := (screenWidth - 400) / 2  // assuming a window width of 400
	y := (screenHeight - 300) / 2 // assuming a window height of 300

	xproto.CreateWindow(X, screen.RootDepth, win, screen.Root,
		int16(x), int16(y), 400, 300, 0,
		xproto.WindowClassInputOutput, screen.RootVisual, 0, []uint32{})

	xproto.ChangeWindowAttributes(X, win,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{ // values must be in the order defined by the protocol
			0xffffffff,
			xproto.EventMaskStructureNotify |
				xproto.EventMaskKeyPress |
				xproto.EventMaskKeyRelease})

	// Set the window's properties to make it floating
	xproto.ChangeWindowAttributes(X, win, xproto.CwOverrideRedirect, []uint32{1})

	xproto.MapWindow(X, win)

	gc, err := xproto.NewGcontextId(X)
	if err != nil {
		slog.Warn(err.Error())
	}
	xproto.CreateGC(
		X,
		gc,
		xproto.Drawable(win),
		xproto.GcForeground|xproto.GcBackground,
		[]uint32{screen.BlackPixel, screen.WhitePixel},
	)

	// Draw initial text
	drawText("Hello, X11 with custom theme!", X, win, gc)

	for counter := 0; counter < 20; counter++ {
		actual := dateutil.Format(dateutil.HowManyDaysToXmas(time.Now(), time.Second), time.Second)
		drawText(actual, X, win, gc)
		time.Sleep(time.Second)
	}

}

func drawText(newText string, conn *xgb.Conn, window xproto.Window, gc xproto.Gcontext) { // Clear the window
	xproto.ClearArea(conn, false, window, 0, 0, 800, 600)
	xproto.ImageText8(conn, byte(len(newText)), xproto.Drawable(window), gc, 50, 50, newText)
}

// Here's an hint:

// package main

// import (
//     "github.com/jazek/xgb"
//     "github.com/jazek/xgb/xproto"
//     "github.com/freedesktop-go/xft"
//     "log"
// )

// func main() {
//     // Connect to the X server
//     conn, err := xgb.NewConn()
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer conn.Close()

//     // Get the screen
//     screen := xproto.Setup(conn).DefaultScreen(conn)

//     // Create a window
//     window, err := xproto.NewWindowId(conn)
//     if err != nil {
//         log.Fatal(err)
//     }

//     xproto.CreateWindow(
//         conn,
//         screen.RootDepth,
//         window,
//         screen.Root,
//         0, 0, 800, 600, 0,
//         xproto.WindowClassInputOutput,
//         screen.RootVisual,
//         xproto.CwBackPixel|xproto.CwEventMask,
//         []uint32{screen.WhitePixel, xproto.EventMaskExposure},
//     )

//     // Map the window (make it visible)
//     xproto.MapWindow(conn, window)

//     // Create a graphics context
//     gc, err := xproto.NewGcontextId(conn)
//     if err != nil {
//         log.Fatal(err)
//     }

//     xproto.CreateGC(
//         conn,
//         gc,
//         xproto.Drawable(window),
//         xproto.GcForeground|xproto.GcBackground,
//         []uint32{screen.BlackPixel, screen.WhitePixel},
//     )

//     // Create Xft draw context
//     draw := xft.DrawCreate(
//         conn,
//         xproto.Drawable(window),
//         screen.RootVisual,
//         xproto.ColormapCopyFromParent,
//     )

//     // Load a font
//     font := xft.FontOpenName(conn, screen.Number, "Arial-16")
//     if font == nil {
//         log.Fatal("Failed to load font")
//     }

//     // Set font color
//     color := xft.ColorAllocName(conn, screen.RootVisual, screen.DefaultColormap, "black")

//     // Draw text
//     drawText := func(text string) {
//         // Clear the window
//         xproto.ClearArea(conn, false, window, 0, 0, 800, 600)

//         // Draw the text using Xft
//         xft.DrawStringUtf8(
//             draw,
//             font,
//             color,
//             50, 50, // Coordinates
//             []byte(text),
//         )
//     }

//     // Draw initial text
//     drawText("Hello, X11 with custom theme!")

//     // Event loop
//     for {
//         ev, err := conn.WaitForEvent()
//         if err != nil {
//             log.Fatal(err)
//         }

//         // Handle expose event
//         if expose, ok := ev.(xproto.ExposeEvent); ok {
//             if expose.Count == 0 {
//                 drawText("Updated Text!")
//             }
//         }
//     }
// }
