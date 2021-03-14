package classifier

import "io"

// Classifier provides a simple interface for different text classifiers
type Classifier interface {
	// Train allows clients to train the classifier
	Train(io.Reader, string) error
	// TrainString allows clients to train the classifier using a string
	TrainString(string, string) error
	// ClassifyString performs text classification using a string
	ClassifyString(string) (string, error)
	// get all probabilities
	Probabilities(string) (map[string]float64, error)
}

// WordCounts extracts term frequencies from a text corpus
func WordCounts(r io.Reader) (map[string]int, error) {
	instream := NewTokenizer().Tokenize(r)
	wc := make(map[string]int)
	for token := range instream {
		wc[token] = wc[token] + 1
	}
	return wc, nil
}
