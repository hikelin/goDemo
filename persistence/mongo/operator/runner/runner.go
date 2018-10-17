package runner

// Operation is the common operation function type
type Operation func([]interface{}) []interface{}

// RunOperation will open mutipule goroutines to run operation
// jobNumer set to 0 or 1 will use single goroutine to insert values
func RunOperation(docsNor []interface{}, operation Operation, jobNumber int) []interface{} {
	if jobNumber < 2 {
		return operation(docsNor)
	}

	jobl := len(docsNor) / jobNumber

	chs := make([]chan []interface{}, jobNumber)
	var ids []interface{}

	for i := 0; i < jobNumber; i++ {
		chs[i] = make(chan []interface{})

		if i < jobNumber-1 {
			go executor(docsNor[jobl*i:jobl*(i+1)], operation, chs[i])
		} else {
			go executor(docsNor[jobl*i:], operation, chs[i])
		}
	}

	for _, ch := range chs {
		cIds := <-ch
		ids = append(ids, cIds...)
	}

	return ids
}

func executor(pars []interface{}, operation Operation, ch chan []interface{}) {
	results := operation(pars)
	ch <- results
}

// IntOperation is the common operation function accept int type
type IntOperation func([]int) []interface{}

// RunIntOperation will open mutipule goroutines to run operation
// jobNumer set to 0 or 1 will use single goroutine to insert values
func RunIntOperation(docsNor []int, operation IntOperation, jobNumber int) []interface{} {

	if jobNumber < 2 {
		return operation(docsNor)
	}

	jobl := len(docsNor) / jobNumber

	chs := make([]chan []interface{}, jobNumber)
	var ids []interface{}

	for i := 0; i < jobNumber; i++ {
		chs[i] = make(chan []interface{})

		if i < jobNumber-1 {
			go intExecutor(docsNor[jobl*i:jobl*(i+1)], operation, chs[i])
		} else {
			go intExecutor(docsNor[jobl*i:], operation, chs[i])
		}
	}

	for _, ch := range chs {
		cIds := <-ch
		ids = append(ids, cIds...)
	}

	return ids
}

func intExecutor(pars []int, operation IntOperation, ch chan []interface{}) {
	results := operation(pars)
	ch <- results
}
