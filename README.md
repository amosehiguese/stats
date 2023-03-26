# Statistics Package

This is a Go package for performing statistical calculations, such as calculating the mean and standard deviation of a data set. It is intended to be used by other Go projects that need statistical functionality.

## Installation

```
go get github.com/amosehiguese/stats
```

## Usage

Here's an example of how to use the `mean` function from this package:

```
package main

import (
  "fmt"
  "github.com/amosehiguese/stats"
)

func main() {
  dataSet := []float64{1.0, 3.0, 5.0}
  result := stats.Mean(dataSet)
  fmt.Printf("The mean is :%.2f\n", result)
}
```

This would output `4.50`, which is the mean of the data set.

## Contributing

If you'd like to contribute to this package, please see the [CONTRIBUTING.md](https://github.com/amosehiguese/stats/CONTRIBUTING.md) file for guidelines

## License

This package is released under the MIT License. See the LICENSE file for details.
