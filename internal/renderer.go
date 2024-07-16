package internal

import (
	"fmt"
	"strings"
)

type ConsolePrintable interface {
	GetConsoleSprite() string
}

const columnWidth = 2

var emptySpace = strings.Repeat(" ", columnWidth)

func Render(m Map) {
	fmt.Println("┌" + strings.Repeat("─", int(m.Width)*columnWidth) + "┐")

	for i := 0; i < int(m.Height); i++ {
		sb := strings.Builder{}

		for j := 0; j < int(m.Width); j++ {
			coordinates := Coordinates{X: uint8(j), Y: uint8(i)}
			entity, found := Find[Entity](m, coordinates)

			if found {
				if entity, ok := entity.(ConsolePrintable); ok {
					sb.WriteString(entity.GetConsoleSprite())
					continue
				}
			}
			sb.WriteString(emptySpace)
		}
		fmt.Println("│" + sb.String() + "│")
	}

	fmt.Println("└" + strings.Repeat("─", int(m.Width)*columnWidth) + "┘")
}
