# classifier
A naive bayes text classifier.

[ ![Codeship Status for n3integration/classifier](https://app.codeship.com/projects/a9a8adf0-d14a-0135-6c51-26e28af241d2/status?branch=master)](https://app.codeship.com/projects/262403)
[![codecov](https://codecov.io/gh/n3integration/classifier/branch/master/graph/badge.svg)](https://codecov.io/gh/n3integration/classifier)
[![Go Report Card](https://goreportcard.com/badge/github.com/kokizzu/classifier)](https://goreportcard.com/report/github.com/kokizzu/classifier)
[![Documentation](https://godoc.org/github.com/kokizzu/classifier?status.svg)](http://godoc.org/github.com/kokizzu/classifier)

## Installation

```bash
go get github.com/kokizzu/classifier
```

## Usage

### Classification

There are two methods of classification: `io.Reader` or `string`. To classify strings, use the `TrainString` or `ClassifyString` functions. To classify larger sources, use the `Train` and `Classify` functions that take an `io.Reader` as input.

```go
import "github.com/kokizzu/classifier/naive"

classifier := naive.New()
classifier.TrainString("The quick brown fox jumped over the lazy dog", "ham")
classifier.TrainString("Earn a degree online", "ham")
classifier.TrainString("Earn cash quick online", "spam")

if classification, err := classifier.ClassifyString("Earn your masters degree online"); err == nil {
    fmt.Println("Classification => ", classification) // ham
} else {
    fmt.Println("error: ", err)
}

probabilities, best := classifier.Probabilities("Quick cash online")
fmt.Println(probabilities)
// map[ham:0.041666666666666664 spam:0.1736111111111111]
fmt.Println(best)
// spam
```

## Contributing

- Fork the repository
- Create a local feature branch
- Run `gofmt`
- Bump the `VERSION` file using [semantic versioning](https://semver.org/)
- Submit a pull request

## License

Copyright 2020 n3integration@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
