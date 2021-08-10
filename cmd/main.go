package main

import (
	"CretRT/pkg/graphics"
	"CretRT/pkg/matht"
	"fmt"
)

type Projectile struct {
	position matht.Tuple
	velocity matht.Tuple
}

type Environment struct {
	gravity matht.Tuple
	wind matht.Tuple
}

func tick(environment Environment, projectile Projectile) Projectile {
	newPosition := matht.Add(projectile.position, projectile.velocity)
	newVelocity := matht.Add(projectile.velocity, matht.Add(environment.gravity, environment.wind))
	return Projectile{newPosition, newVelocity}
}

func main() {
	width := 900
	height := 550
	c := graphics.CreateCanvas(width, height)
	p := Projectile{matht.Point(0, 1, 0), matht.Normalize(matht.Vector(1, 100, 0))}
	e := Environment{matht.Vector(0, -0.01, 0), matht.Vector(0.1, 0, 0)}

	for p.position.Y > 0 {
		cond1 := p.position.X >= 0 && p.position.X < float64(width)
		cond2 := p.position.Y >= 0 && p.position.Y < float64(height)

		if cond1 && cond2 {
			c.WritePixel(int(p.position.X), int(p.position.Y), matht.Color(1, 1, 1))
			fmt.Printf("X : %f, Y : %f, Z : %f\n", p.position.X, p.position.Y, p.position.Z)
		}

		p = tick(e, p)
	}

	c.WriteCanvasToFile("test.ppm")
}
