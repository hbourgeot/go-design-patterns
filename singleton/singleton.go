package main

import (
	"fmt"
	"sync"
	"time"
)

// estructura para simular una base de datos

type Database struct{}

//metodo para simular una conexión a la base de
//datos falsa

func (d Database) createSingleConnection() {
	fmt.Println("Creating singleton for Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")
}

var db *Database

// instancia de la base de datos

var lock sync.Mutex

// instancia para bloquear y desbloquear la
// ejecución de un fragmento de código de las
// goroutines para que una vez que un
// goroutine acceda a ese fragmento las demás
// esperen a que esta goroutine termine

func getDatabaseInstance() *Database {
	lock.Lock()
	// cuando una de las goroutines acceda a este
	// punto,las demás esperarán a que esta finalice

	defer lock.Unlock()
	// al finalizar la función, las demás goroutines
	// podrán acceder al cuerpo de la función

	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.createSingleConnection()
	} else {
		fmt.Println("DB Already created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup // nos ayuda a manejar la concurrencia
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}
