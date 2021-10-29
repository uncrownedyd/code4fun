package main

import "fmt"

type Opener interface{ Open() }
type Closer interface{ Close() }
type PutInsider interface { PutInside() }
type Checker interface{ Check() }

type NormalOpener struct{}
type NormalCloser struct{checker Checker}
type NormalPutInsider struct{}

func (*NormalOpener) Open() {
	fmt.Println("Normal Open")
}
func (*NormalPutInsider) PutInside() {
	fmt.Println("Normal PutInside")
}
func (n *NormalCloser) Close() {
	fmt.Println("Normal Close")
	n.checker.Check()
}

type PinPutInsider struct{}
type PinChecker struct{}

func (*PinPutInsider) PutInside() {
	fmt.Println("Pin PutInside")
}
func (*PinChecker) Check() {
	fmt.Println("Pin Check")
}

func main() {
	opener := NormalOpener{}
	putInsider := PinPutInsider{}
	checker := PinChecker{}
	closer := NormalCloser{checker: &checker}

	putElephant(&opener, &putInsider, &closer)
}

func putElephant(opener Opener, insider PutInsider, closer Closer) {
	opener.Open()
	insider.PutInside()
	closer.Close()
}