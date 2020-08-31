package client

var graffitiMap = []string{
	"graffitiwall:502:507:#000000",
	"graffitiwall:503:508:#000000",
	"graffitiwall:505:510:#000000",
	"graffitiwall:506:511:#000000",
	"graffitiwall:507:506:#000000",
	"graffitiwall:508:506:#000000",
	"graffitiwall:509:510:#82d318",
	"graffitiwall:510:508:#82d318",
	"graffitiwall:510:509:#000000",
	"graffitiwall:511:512:#82d318",
	"graffitiwall:512:506:#82d318",
	"graffitiwall:512:507:#82d318",
	"graffitiwall:512:508:#000000",
	"graffitiwall:512:510:#000000",
	"graffitiwall:512:511:#82d318",
	"graffitiwall:512:512:#82d318",
	"graffitiwall:513:505:#82d318",
	"graffitiwall:513:506:#82d318",
	"graffitiwall:513:507:#82d318",
	"graffitiwall:513:510:#82d318",
	"graffitiwall:514:507:#82d318",
	"graffitiwall:514:508:#82d318",
	"graffitiwall:515:511:#82d318",
	"graffitiwall:515:512:#82d318",
	"graffitiwall:517:510:#82d318",
}

var graffitiWall = make(chan []byte)

func init() {
	go func() {
		for _, graffitiHex := range graffitiMap {
			graffiti := []byte(graffitiHex)
			graffitiWall <- graffiti
		}
		close(graffitiWall)
	}()
}
