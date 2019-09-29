package main

import "math"

type Vec2 struct {
	X, Y float32
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// 使用矢量减去另外一个矢量生成一个新的矢量
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}
// 使用矢量乘以另外一个矢量，生成新的矢量
func (v Vec2) Scale(s float32) Vec2  {
	return Vec2{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// 计算两个矢量的距离
func (v Vec2) DistanceTo(other Vec2) float32  {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (v Vec2) Normalize() Vec2  {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0{
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}

	return Vec2{0, 0}
}
