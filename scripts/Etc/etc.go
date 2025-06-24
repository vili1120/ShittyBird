package etc

import rl "github.com/gen2brain/raylib-go/raylib"

func NewBoundingBox(x, y, w, h float64) *BoundingBox {
	bb := &BoundingBox{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
	bb.Initiate()
	return bb
}

type BoundingBox struct {
	X, Y float64
	W, H float64
	Rect rl.Rectangle
}

func (b *BoundingBox) Initiate() {
	b.Rect = rl.Rectangle{X: float32(b.X), Y: float32(b.Y), Width: float32(b.W), Height: float32(b.H)}
}

func (b *BoundingBox) IsColliding(rect rl.Rectangle) bool {
	return rl.CheckCollisionRecs(b.Rect, rect)
}

type Sprite struct {
	Image rl.Texture2D
	BB *BoundingBox
	V2d rl.Vector2
}
