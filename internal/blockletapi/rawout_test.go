package blockletapi

import (
	"os"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
)

func ExampleRawOut_Print_ok() {
	out := NewRawOut(os.Stdout)
	daysToXmas := time.Duration(10 * dateutil.Day)
	out.Print(daysToXmas, &RenderOptions{Label: "test", Format: "text", Unit: dateutil.Day})
	// Output:
	// 10
	// 10
}

// func ExampleRawOut_Print_ok_withColor() {
// 	out := NewRawOut(os.Stdout)
// 	daysToXmas := time.Duration(10 * dateutil.Day)
// 	out.Print(daysToXmas, &RenderOptions{Label: "test", Format: "text", Unit: dateutil.Day, BackgroundColor: "#ff0000", ForegroundColor: "#00ff00"})
// 	// Output:
// 	// test10
// 	// test10
// 	// #00ff00
// 	// #ff0000

// }
