### Forever
Forever is a helper for running forever loops with limited number of go routines.  
You can use Mutex between these go routines, also you can use channels by defining channels as global.

### Usage
```go
import "github.com/panwenbin/forever"

forever.Go(func(mux *sync.Mutex) {
    // do something in a forever loop, with 3 go routines
}, 3)

```