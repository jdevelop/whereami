package display

type Display interface {
	Cls() error
	Println(message string) error
}
