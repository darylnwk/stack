# stack
--
    import "github.com/darylnwk/stack"


## Usage

#### type Stack

```go
type Stack struct {
}
```

Stack defines a thread safe stack

#### func  New

```go
func New() *Stack
```
New initialises a new Stack

#### func (*Stack) Pop

```go
func (s *Stack) Pop() interface{}
```
Pop returns the top element and nil if the stack is empty

#### func (*Stack) Push

```go
func (s *Stack) Push(elements ...interface{})
```
Push adds elements to the stack

#### func (*Stack) Size

```go
func (s *Stack) Size() int
```
Size returns the total number of elements in the stack
