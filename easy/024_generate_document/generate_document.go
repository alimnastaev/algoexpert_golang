package generate_document

/*
! Generate Document
https://www.algoexpert.io/questions/generate-document


! Example:
*/

func GenerateDocument(characters string, document string) (output bool) {
	lookupTable := make(map[string]int)

	for _, char := range characters {
		lookupTable[string(char)]++

	}

	for _, char := range document {
		if lookupTable[string(char)] != 0 {
			output = true
			lookupTable[string(char)]--
		}

	}

	return output
}
