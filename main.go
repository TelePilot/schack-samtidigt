package main

import (
	"fmt"
	"image/color"
	"unicode"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Board        [][]Square
	Position     string
}
type Square struct {
	Color     color.RGBA
	PositionX int
	PositionY int
	Piece     *ebiten.Image
}

const SquareSize int = 64
const StartingPosition string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

// cant be constant why??? is there a better struct for this?
var PieceValues = map[string]int{"k": 1, "p": 2,
	"n": 3, "b": 4, "r": 5, "q": 6, "black": 0, "white": 1}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	//after move
	// loadFenPosition
	// g.Position = newFenPosition
	return nil
}

func createBoard() [][]Square {

	board := make([][]Square, 8)
	//Rank is chess for horizontal lines
	//Files is chess for vertical lines
	for rank := 0; rank < 8; rank++ {
		board[rank] = make([]Square, 8)
		for file := range board[rank] {

			square := Square{
				PositionX: file * SquareSize,
				PositionY: rank * SquareSize,
			}

			if (rank+file)%2 == 0 {
				//white square
				square.Color = color.RGBA{237, 215, 175, 255}
			} else {
				// black square
				square.Color = color.RGBA{185, 135, 97, 255}
			}
			board[rank][file] = square
		}
	}
	return board
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).

//Ebitengine matrix is x horizontal and y is vertical

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{72, 71, 66, 255})
	// Write your game's rendering.
	for _, rank := range g.Board {
		for _, file := range rank {
			s := ebiten.NewImage(SquareSize, SquareSize)
			s.Fill(file.Color)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(file.PositionX), float64(file.PositionY))

			screen.DrawImage(s, op)

		}
	}
	// When we add the pieces after the board the z-index is higher
	for _, rank := range g.Board {
		for _, file := range rank {
			if file.Piece != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(file.PositionX), float64(file.PositionY))
				screen.DrawImage(file.Piece, op)
			}

		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}

func NewGame() *Game {
	g := &Game{}
	g.ScreenHeight = 512
	g.ScreenWidth = 512
	g.Board = createBoard()
	g.Position = StartingPosition
	loadFenPosition(g)
	return g
}

func loadFenPosition(g *Game) {
	rank := 0
	file := 0

	for _, v := range g.Position {
		if v == '/' {
			//New line
			rank++
			file = 0
			continue
		}

		if unicode.IsDigit(v) {
			//Skip as many files as digit says
			file += int(v)
			continue
		}
		//get the piece value from the map
		p := PieceValues[string(unicode.ToLower(v))]
		//black as baseline
		c := PieceValues["black"]
		//Check if actual rune is white
		if unicode.IsUpper(v) {
			// change to White
			c = PieceValues["white"]

		}
		// Pointer to board square
		s := &g.Board[rank][file]
		path := fmt.Sprintf("%d%d.png", p, c)
		pImg, _, err := ebitenutil.NewImageFromFile("icons/" + path)
		if err != nil {
			panic(err)
		}
		// add piece image to square
		s.Piece = pImg
		file++
	}
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Schack Samtidigt")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
