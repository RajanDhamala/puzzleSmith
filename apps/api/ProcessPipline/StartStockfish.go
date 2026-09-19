package processpipline

import (
	"context"
	"fmt"

	"github.com/RajanDhamala/go-stockfish"
)

func ConnectStockfish() *stockfish.Client {
	client, err := stockfish.New(context.Background(), stockfish.Config{
		PoolSize:         4,
		QueueSize:        16,
		PerEngineThreads: 1,
		TotalHashMB:      128,
		MaxMultiPV:       3,
	})
	if err != nil { /* handle */
		fmt.Println("error while init stockfish worker")
		return nil
	}
	fmt.Println("stockfish client init succesfully")
	return client
}
