
// return value and error 
func load(fname string) ([]string, error) {
	if fname == "" {
	// Go code uses the built-in type error to signal when an error occurred during execution of a function.
		return nil, errors.New(
		"Dictionary file name cannot be empty.")
	}
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
	}

func main() {
	words, err := load("dict.txt")
	if err != nil {
	fmt.Println("Unable to load file:", err)
	os.Exit(1)
}


// defer function calls

func do(steps ...string) {
	defer fmt.Println("All done!")
	for _, s := range steps {
		defer fmt.Println(s)
	}
	fmt.Println("Starting")
}

func main() {
	do(
		"Find key",
		"Aplly break",
		"Put key in ignition",
		"Start car",
	)
}


// use defer for resource cleanup

// Closing open files
// Releasing network resources
// Closing the Go channel
// Committing database transactions
// And do on



// panic and recovery
// When a function panics, it aborts and executes its deferred calls. Then its caller panics, causing a chain reaction
// The panic sequence continues all the way up the call stack until the main function is
// reached and the program exits (crashes).


// function panic recovery
// a call to function recover returns the value that was passed as an argument to panic

package main

func write(fname string, anagrams map[string][]string) {
	file, err := os.OpenFile(
		fname,
		os.O_WRONLY+os.O_CREATE+os.O_EXCL,
		0644,
	)
	if err != nil {
		msg := fmt.Sprintf(
		"Unable to create output file: %v", err,
		)
	panic(msg)
	}
}

func makeAnagrams(words []string, fname string) {
	defer func() {
	if r := recover(); r != nil {
		fmt.Println("Failed to make anagram:", r)
		}
	}()
	anagrams := mapWords(words)
	write(fname, anagrams)
}

func main() {
	words, err := load("")
	if err != nil {
		fmt.Println("Unable to load file:", err)
		os.Exit(1)
	}
	makeAnagrams(words, "")
}

