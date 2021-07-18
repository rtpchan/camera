package camera

import (
	"testing"
)

var c *Camera

func TestScreenPos(t *testing.T) {

	c = NewCamera(400, 200)
	c.LookAt(600.0, 1200.0)
	itemX := 500.0
	itemY := 1150.0
	screenY := 150
	screenX := 100

	sx, sy, isOnScreen := c.OnScreenPoint(itemX, itemY)
	if sx != screenX {
		t.Errorf("X wrong, want %d, got %d", screenX, sx)
	}
	if sy != screenY {
		t.Errorf("Y wrong, want %d, got %d", screenY, sy)
	}
	if !isOnScreen {
		t.Errorf("Is on screen, want true, got %v", false)
	}
	isBoxOnScreen := c.OnScreenBox(sx, sy, 100, 100)
	if !isBoxOnScreen {

		t.Errorf("Is on screen, want true, got %v", false)
	}

}

func TestScreenBoxInside(t *testing.T) {

	c = NewCamera(200, 100)
	c.LookAt(600.0, 1200.0)
	itemX := 450.0
	itemY := 1280.0
	screenY := -30
	screenX := -50

	sx, sy, isOnScreen := c.OnScreenPoint(itemX, itemY)
	if sx != screenX {
		t.Errorf("X wrong, want %d, got %d", screenX, sx)
	}
	if sy != screenY {
		t.Errorf("Y wrong, want %d, got %d", screenY, sy)
	}
	if isOnScreen {
		t.Errorf("Is on screen, want false, got %v", true)
	}
	isBoxOnScreen := c.OnScreenBox(sx, sy, 100, 100)
	if !isBoxOnScreen {

		t.Errorf("Is on screen, want true, got %v", false)
	}
}

func TestZoom(t *testing.T) {
	c = NewCamera(300, 200)
	c.SetZoom(2.0)
	c.LookAt(0, 0)
	c.LookAt(10, 10)
	x, y, _ := c.OnScreenPoint(-20, 10)
	if !(x == 90 && y == 100) {
		t.Errorf("Zooming, want x=90, y=100, got x=%d, y=%d", x, y)
	}
}
