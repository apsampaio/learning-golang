```go
// Create pointer to a variable and change its value

	pPointer := &p
	pPointer.updateName("Jim")
	p.print()

func (pPointer *person) updateName(new string) {
	(*pPointer).firstName = new
}
```
