# almost-blockchain-two

![go workflow](https://github.com/zhooda/almost-blockchain-two/workflows/Go/badge.svg)

almost-blockchain-two is the sequel to [almost-blockchain](https://github.com/zhooda/almost-blockchain), a half-baked blockchain and cryptocurrency I created as my final project for high school computer science. 

## Getting Started

These instructions *should* get you a copy of the project up and running
on your local machine. There are no deployment strategies right now `:~)`

### Prerequisites

To run almost-blockchain-two you will need:

- [golang](https://golang.org/) v1.15+

On Windows you will also need:

- [chocolatey](https://chocolatey.org/)

### Installation

Clone the repository or checkout with subversion to get a copy of almost-blockchain-two

```bash
$ git clone https://github.com/zhooda/almost-blockchain-two
```

## Usage

Build and run the very basic chain that currently exists!

### Linux

```bash
$ cd ./path/to/almost-blockchain-two

# Build and run
$ make
$ ./bin/abc2
$ make clean

# OR run without building
$ make run
```

### Windows

```powershell
PS> choco install make

PS> cd .\path\to\almost-blockchain-two

# Build and run
PS> make
PS> .\bin\abc2.exe
PS> make clean

# OR run without building
PS> make run
```

## License
almost-blockchain-two uses the [MIT](https://choosealicense.com/licenses/mit/) license `:~)`
