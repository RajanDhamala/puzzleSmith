package processpipline

import (
	"fmt"
	_ "time"

	"github.com/RajanDhamala/go-stockfish"
	types "github.com/RajanDhamala/puzzleSmith/Types"
	utils "github.com/RajanDhamala/puzzleSmith/Utils"
	lib "github.com/corentings/chess/v2"
)

func StartProcessingGame(username string, games *types.UserGames, client *stockfish.Client) {
	fmt.Println("we got the games to process do 2 rn")

	// first parse the games extract each move san
	for _, game := range games.Games {
		head, tail := utils.SplitPgn(game.PGN)

		fmt.Println("paresed sucesfully", tail[0], head.Link)
		fmt.Println("usr:", game.White.Rating)
		PlayGame(game, tail, client)
	}
}

func PlayGame(game *types.Game, moves []types.Move, client *stockfish.Client) {
	newGame := lib.NewGame()

	for i, move := range moves {

		// err := newGame.PushNotationMove(move.San, lib.AlgebraicNotation{}, nil)

		// using unsafe push for higher performance and the moves san are parsed safely and might
		// have edge case need to look more into it

		err := newGame.UnsafePushNotationMove(move.San, lib.AlgebraicNotation{}, nil)
		if err != nil {
			fmt.Println("failed to play move", err.Error())
			return
		}

		if i < 10 {
			continue
		}
		fmt.Println("fen:", newGame.FEN())
		// result, err := client.Evaluate(context.Background(), stockfish.EvalRequest{
		// 	FEN: newGame.FEN(),
		// 	// MoveTime: 300 * time.Millisecond,
		// 	Depth:   17,
		// 	MultiPV: 1,
		// })
		// if err != nil { /* handle */
		// 	fmt.Println("rror while processing game", err.Error())
		// }
		//
		// fmt.Println(result.BestMove) // top line (MultiPV=1)
		// for _, line := range result.Lines {
		// 	fmt.Printf("line=%d depth=%d pv=%v cp=%v mate=%v\n", line.MultiPV, line.Depth, line.PV, line.ScoreCP, line.Mate)
		// }
	}
	final := newGame.CurrentPosition()

	fmt.Println("final position:", final)
}

func NormalizeEval(score int, isWhite bool) int {
	if isWhite {
		return score
	}
	return -score
}
