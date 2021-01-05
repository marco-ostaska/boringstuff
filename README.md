# boringstuff
--
    import "github.com/marco-ostaska/boringstuff"

Package boringstuff wraps some boring tasks.

## Usage

#### func  EpochToUnix

```go
func EpochToUnix(epoch string) (time.Time, error)
```
EpochToUnix converts epoch time to Unix time.

#### func  ReturnError

```go
func ReturnError(errs ...error) error
```
ReturnError checks for multiple errors at once.
