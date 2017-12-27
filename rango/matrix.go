package rango

import "math"

type Matrix struct {
	m00, m01, m02, m03 float64
	m10, m11, m12, m13 float64
	m20, m21, m22, m23 float64
	m30, m31, m32, m33 float64
}

func deg2rad(deg float64) float64 {
	return (math.Pi * deg) / 180.0
}

func Identity() Matrix {
	return Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func MatrixMultiply(a Matrix, b Matrix) Matrix {
	m := Matrix{}
	m.m00 = a.m00*b.m00 + a.m01*b.m10 + a.m02*b.m20 + a.m03*b.m30
	m.m10 = a.m10*b.m00 + a.m11*b.m10 + a.m12*b.m20 + a.m13*b.m30
	m.m20 = a.m20*b.m00 + a.m21*b.m10 + a.m22*b.m20 + a.m23*b.m30
	m.m30 = a.m30*b.m00 + a.m31*b.m10 + a.m32*b.m20 + a.m33*b.m30
	m.m01 = a.m00*b.m01 + a.m01*b.m11 + a.m02*b.m21 + a.m03*b.m31
	m.m11 = a.m10*b.m01 + a.m11*b.m11 + a.m12*b.m21 + a.m13*b.m31
	m.m21 = a.m20*b.m01 + a.m21*b.m11 + a.m22*b.m21 + a.m23*b.m31
	m.m31 = a.m30*b.m01 + a.m31*b.m11 + a.m32*b.m21 + a.m33*b.m31
	m.m02 = a.m00*b.m02 + a.m01*b.m12 + a.m02*b.m22 + a.m03*b.m32
	m.m12 = a.m10*b.m02 + a.m11*b.m12 + a.m12*b.m22 + a.m13*b.m32
	m.m22 = a.m20*b.m02 + a.m21*b.m12 + a.m22*b.m22 + a.m23*b.m32
	m.m32 = a.m30*b.m02 + a.m31*b.m12 + a.m32*b.m22 + a.m33*b.m32
	m.m03 = a.m00*b.m03 + a.m01*b.m13 + a.m02*b.m23 + a.m03*b.m33
	m.m13 = a.m10*b.m03 + a.m11*b.m13 + a.m12*b.m23 + a.m13*b.m33
	m.m23 = a.m20*b.m03 + a.m21*b.m13 + a.m22*b.m23 + a.m23*b.m33
	m.m33 = a.m30*b.m03 + a.m31*b.m13 + a.m32*b.m23 + a.m33*b.m33
	return m
}

func RotateX(rx float64) Matrix {
	rxRad := deg2rad(rx)
	cosrx := math.Cos(rxRad)
	sinrx := math.Sin(rxRad)
	return Matrix{
		1, 0, 0, 0,
		0, cosrx, -sinrx, 0,
		0, sinrx, cosrx, 0,
		0, 0, 0, 1,
	}
}

func RotateY(ry float64) Matrix {
	ryRad := deg2rad(ry)
	cosry := math.Cos(ryRad)
	sinry := math.Sin(ryRad)
	return Matrix{
		cosry, 0, sinry, 0,
		0, 1, 0, 0,
		-sinry, 0, cosry, 0,
		0, 0, 0, 1,
	}
}

func RotateZ(rz float64) Matrix {
	rzRad := deg2rad(rz)
	cosrz := math.Cos(rzRad)
	sinrz := math.Sin(rzRad)
	return Matrix{
		cosrz, -sinrz, 0, 0,
		sinrz, cosrz, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Translate(tx, ty, tz float64) Matrix {
	return Matrix{
		1, 0, 0, tx,
		0, 1, 0, ty,
		0, 0, 1, tz,
		0, 0, 0, 1,
	}
}

func Scale(sx, sy, sz float64) Matrix {
	return Matrix{
		sx, 0, 0, 0,
		0, sy, 0, 0,
		0, 0, sz, 0,
		0, 0, 0, 1,
	}
}

func Rotate(rx, ry, rz float64) Matrix {
	return MatrixMultiply(MatrixMultiply(RotateX(rx), RotateY(ry)), RotateZ(rz))
}

func MatrixVecMultiply(m Matrix, vector Vector) Vector {
	vec := Vector{}

	vec.X = m.m00*vector.X + m.m01*vector.Y + m.m02*vector.Z + m.m03
	vec.Y = m.m10*vector.X + m.m11*vector.Y + m.m12*vector.Z + m.m13
	vec.Z = m.m20*vector.X + m.m21*vector.Y + m.m22*vector.Z + m.m23
	return vec
}
