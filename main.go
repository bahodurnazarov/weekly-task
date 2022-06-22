package main

import (
	"log"
	"os"
	"time"
)

func main() {

	// type Timer interface {
	// 	Start()
	// 	Stop()
	// }

	// type StopWath struct {

	// }

	// func (s *StopWath) Start() {

	// }

	// func (s *StopWath) Stop() {

	// }

	file, err := os.Create("file/log.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer func()  {
		if cerr := file.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}
			time.Sleep(time.Second)	
			_, err = file.Write([]byte("tick\t"))
			if err != nil {
				log.Print(err)
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}
			time.Sleep(time.Second * 5)
			_, err = file.Write([]byte("TACK\n"))
			if err != nil {
				log.Print(err)
				return
			}
		}
	}()

	time.Sleep(time.Second * 60)
	done <- struct{}{}

	

	

	// err := ioutil.WriteFile("file/log.txt", []byte("done1"), 0666)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
}

// func WriteFile(filename string, data []byte, perm os.FileMode) error {
// 	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = f.Write(data)
// 	if err1 := f.Close(); err == nil {
// 		err = err1
// 	}
// 	return err
// }
