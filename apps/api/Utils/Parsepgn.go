package utils

import (
	"fmt"
	"strings"

	// "encoding/json"
	// demo "github.com/notnil/chess"
	// "github.com/RajanDhamala/puzzleSmith/ProcessPipline"
	"github.com/RajanDhamala/puzzleSmith/Types"
)

func SplitPgn(pgn string) (*types.Pgn, []types.Move) {
	splitted := strings.SplitN(pgn, "\n\n", 2)
	header := ParsePngHeader(splitted[0])
	moves := ParsePgnBody(splitted[1])

	// pgnReader := strings.NewReader(pgn)
	// pgnData, err := demo.PGN(pgnReader) // returns *chess.PGN
	// if err != nil {
	// 	fmt.Println("failed to read the PGN")
	// }
	// game := demo.NewGame(pgnData)
	// for _, move := range game.Moves() {
	// 	fmt.Println("SAN:", move.String()) // or move.String()
	// }
	return header, moves
}

func ParsePngHeader(header string) *types.Pgn {
	pg := types.Pgn{}
	lines := strings.Split(header, "\n")
	for _, line := range lines {
		trimmed := strings.Trim(line, "[]")
		parts := strings.SplitN(trimmed, " ", 2)
		if len(parts) < 2 {
			continue
		}
		key := parts[0]
		val := strings.Trim(parts[1], `"`)

		switch key {
		case "Event":
			pg.Event = val
		case "Site":
			pg.Site = val
		case "Date":
			pg.Date = val
		case "White":
			pg.White = val
		case "Black":
			pg.Black = val
		case "Result":
			pg.Result = val
		case "CurrentPosition":
			pg.CurrentPosition = val
		case "Timezone":
			pg.Timezone = val
		case "ECO":
			pg.ECO = val
		case "ECOUrl":
			pg.ECOUrl = val
		case "UTCDate":
			pg.UTCDate = val
		case "UTCTime":
			pg.UTCTime = val
		case "WhiteElo":
			pg.WhiteElo = val
		case "BlackElo":
			pg.BlackElo = val
		case "TimeControl":
			pg.TimeControl = val
		case "Termination":
			pg.Termination = val
		case "StartTime":
			pg.StartTime = val
		case "EndDate":
			pg.EndDate = val
		case "EndTime":
			pg.EndTime = val
		case "Link":
			pg.Link = val
		}
	}
	return &pg
}

func ParsePgnBody(body string) []types.Move {
	bodyArray := strings.Fields(body)
	var moves []types.Move
	// Result := ""

	for i, item := range bodyArray {
		if strings.HasSuffix(item, ".") {
			continue
		}

		if strings.HasPrefix(item, "{[%clk") {
			if i+1 < len(bodyArray) {
				clock := bodyArray[i+1]
				sanitized := strings.Trim(clock, "]}")
				if len(moves) > 0 {
					moves[len(moves)-1].Clock = sanitized
				}
			}
			continue
		}

		if strings.HasPrefix(item, "0:") {
			continue
		}
		if isPGNResultToken(item) {
			continue
		}

		moves = append(moves, types.Move{
			San: item,
		})
	}
	// b, _ := json.MarshalIndent(moves, "", " ")
	// fmt.Println(string(b))
	// fmt.Println("result:", Result)
	return moves
}

func isPGNResultToken(token string) bool {
	switch strings.TrimSpace(token) {
	case "1-0", "0-1", "1/2-1/2", "*":
		return true
	default:
		return false
	}
}

func MaterialCount(fen string) (int, int) {
	white := 0
	black := 0
	fmt.Println("we got the posion:", fen)
	board := strings.SplitN(fen, " ", 2)[0]

	for _, ch := range board {
		switch ch {
		case 'P':
			white = white + 1
		case 'Q':
			white = white + 9
		case 'N', 'B':
			white = white + 3
		case 'R':
			white = white + 5
		case 'p':
			black = black + 1
		case 'q':
			black = black + 9
		case 'n', 'b':
			black = black + 3
		case 'r':
			black = black + 5
		}
	}
	return white, black
}
