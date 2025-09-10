package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var m = sync.RWMutex{} // para proteger variables compartidas (ej: results
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	// secuential()
	// gorutines()
	countGorutines()
}

// Ejemplo de uso de gorutinas para tareas que no son I/O
// (llamadas a base de datos, llamadas a APIs, etc)
// sino tareas que consumen CPU
// En este caso, se lanzan 800 gorutinas que hacen un conteo hasta 1 millón
// y se mide el tiempo total que tarda en completarse la tarea
// Se puede comparar con el tiempo que tarda en completarse la misma tarea
// de forma secuencial (sin gorutinas)
func countGorutines() {
	t0 := time.Now()
	for i := 0; i < 800; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("Total time gorutines: %v\n", time.Since(t0))
}

func count() {
	var res int
	for i := 0; i < 1000000; i++ {
		res += i
	}
	wg.Done()
}

func secuential() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		dbCalls(i)
	}
	fmt.Printf("Total time secuential: %v\n", time.Since(t0))
}

func gorutines() {
	fmt.Println("Cores usados por Go:", runtime.GOMAXPROCS(0))
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go grDbCalls(i)
	}
	wg.Wait()
	fmt.Printf("Total time gorutines: %v\n", time.Since(t0))
	fmt.Println("Results: ", results)
}

func dbCalls(i int) {
	// Simular llamada a base de datos
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("the result from the database is: ", dbData[i])
}

func grDbCalls(i int) {
	// Simular llamada a base de datos
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("The current results are: ", results)
	m.RUnlock()
}
