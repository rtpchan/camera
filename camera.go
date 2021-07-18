// Camera functions
// NewCamera(screenW, screenH int)  screen size
// LootAt(x, y float64) world location
// Place(x, y float64, imgW, imgH int) place image (world) onto screen

package camera

import "fmt"

// Camera provide coordinate transformation
type Camera struct {
	screenW, screenH int

	// World coordinates
	lookAtX, lookAtY             float64 // pointing camera at
	sTop, sBottom, sLeft, sRight float64 // area that is on screen
	zoom                         float64
	zoomInv                      float64
}

// NewCamera setup the camera
func NewCamera(width, height int) *Camera {
	c := Camera{}
	c.screenW = width
	c.screenH = height
	c.zoom = 1.0
	c.zoomInv = 1.0
	c.LookAt(0, 0)

	return &c
}

// LookingAt returns position the camera is pointing to
func (c *Camera) LookingAt() (float64, float64) {
	return c.lookAtX, c.lookAtY
}

// LookAt points camera to World location x, y
func (c *Camera) LookAt(x, y float64) {
	c.lookAtX = x
	c.lookAtY = y
	c.sTop = c.lookAtY + float64(c.screenH/2)*c.zoomInv
	c.sBottom = c.lookAtY - float64(c.screenH/2)*c.zoomInv
	c.sLeft = c.lookAtX - float64(c.screenW/2)*c.zoomInv
	c.sRight = c.lookAtX + float64(c.screenW/2)*c.zoomInv

}

// SetZoom set camera zoom level
func (c *Camera) SetZoom(z float64) {
	if z == 0.0 {
		return
	}
	c.zoom = z
	c.zoomInv = 1 / z
	c.sTop = c.lookAtY + float64(c.screenH/2)*c.zoomInv
	c.sBottom = c.lookAtY - float64(c.screenH/2)*c.zoomInv
	c.sLeft = c.lookAtX - float64(c.screenW/2)*c.zoomInv
	c.sRight = c.lookAtX + float64(c.screenW/2)*c.zoomInv
}

// Debug print debug info to terminal
func (c *Camera) Debug(isDebug bool) {
	if isDebug {
		fmt.Println(c.lookAtX, c.lookAtY, c.sTop, c.sBottom, c.sLeft, c.sRight)
	}
}

// WtoS convert world coordinates to screen coordinates
func (c *Camera) WtoS(x, y float64) (float64, float64) {
	var sx, sy float64
	sx = (x-c.lookAtX)*c.zoom + float64(c.screenW)/2.0
	sy = (c.lookAtY-y)*c.zoom + float64(c.screenH)/2.0
	return sx, sy
}

// StoW convert screen coordinates to world coordinates
// e.g. mouse position used to detect items
func (c *Camera) StoW(sx, sy int) (float64, float64) {
	var x, y float64
	x = (float64(sx)-float64(c.screenW)/2.0)/c.zoom + c.lookAtX
	y = c.lookAtY - (float64(sy)-float64(c.screenH)/2.0)/c.zoom
	return x, y
}

// Zoom return camera zoom level
func (c *Camera) Zoom() float64 {
	return c.zoom
}

// ScreenSize return screen size
func (c *Camera) ScreenSize() (int, int) {
	return c.screenW, c.screenH
}

// SetScreen size
func (c *Camera) SetScreenSize(w, h int) {
	c.screenW, c.screenH = w, h
}

// CentreImage return the top left co-ordinates for drawing
// (use if image are drawn with top left at 0,0)
func (c *Camera) CentreImage(sx, sy, boxW, boxH int) (int, int) {
	return (sx - boxW/2), (sy - boxH/2)
}

// Place returns centre coordinates on screen for a given image size
func (c *Camera) Place(x, y float64, imgW, imgH int) (int, int, bool) {
	var sx, sy int
	var isOnScreen bool
	sx, sy, isOnScreen = c.OnScreenPoint(x, y)
	if isOnScreen {
		isOnScreen = c.OnScreenBox(sx, sy, imgW, imgH)
	}
	if isOnScreen {
		sx, sy = c.CentreImage(sx, sy, imgW, imgH)
	}
	return sx, sy, isOnScreen
}

// OnScreenPoint return the on screen position from world position, and image is on screen
// Screen right -> x, down -> y
// World right -> x, up -> y
func (c *Camera) OnScreenPoint(x, y float64) (int, int, bool) {
	var sx, sy int
	var onScreenX, onScreenY bool
	sx = int((x-c.lookAtX)*c.zoom + float64(c.screenW)/2.0)
	sy = int((c.lookAtY-y)*c.zoom + float64(c.screenH)/2.0)

	if x < c.sLeft {
		onScreenX = false
	} else if x > c.sRight {
		onScreenX = false
	} else {
		onScreenX = true
	}

	if y > c.sTop {
		onScreenY = false
	} else if y < c.sBottom {
		onScreenY = false
	} else {
		onScreenY = true
	}
	return sx, sy, onScreenX && onScreenY
}

// OnScreenBox checks if bounding box is on screen
// TODO, fix bug
func (c *Camera) OnScreenBox(sx, sy, boxW, boxH int) bool {
	if (sx - boxW/2) > c.screenW {
		return false
	}
	if (sx + boxW/2) < 0 {
		return false
	}
	if (sy - boxH/2) > c.screenH {
		return false
	}
	if (sy + boxH/2) < 0 {
		return false
	}
	return true
}
