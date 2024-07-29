package another_package // import "github.com/jsokel/counterfeiter/v6/fixtures/another_package"

type SomeType int

//go:generate go run github.com/jsokel/counterfeiter/v6 . AnotherInterface
type AnotherInterface interface {
	AnotherMethod([]SomeType, map[SomeType]SomeType, *SomeType, SomeType, chan SomeType)
}
