<div style="text-align: center">
<h1 style="text-decoration:">almost-blockchain-two</h1>
<img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/zhooda/almost-blockchain-two/Go?style=flat-square">
<img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/zhooda/almost-blockchain-two?color=crimson&label=lines%20of%20code&style=flat-square">
</div>

<div style="margin-bottom: 2%"></div>

almost-blockchain-two is the sequel to [almost-blockchain](https://github.com/zhooda/almost-blockchain), a half-baked blockchain and cryptocurrency I created as my final project for high school computer science. 

<div style="margin-bottom: 2%"></div>

---
<div style="margin-bottom: 2%"></div>

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

Build and run the very basic chain that currently exists! Trying to use `go run src/main.go` or `go build src/main.go` will not work becuase of idiosyncrasies with go modules and folder structures that I dont understand yet.

#### Linux/macOS

```bash
$ cd ./path/to/almost-blockchain-two

# Build and run manually
$ make
$ ./bin/abc2 print
$ ./bin/abc2 add -block "this is some block data"
$ make clean

# OR build and run with makefile
$ make run ARGS="print"
```

#### Windows

```powershell
PS> choco install make

PS> cd .\path\to\almost-blockchain-two

# Build and run manually
PS> make
PS> .\bin\abc2.exe print
PS> .\bin\abc2.exe add -block "this is some block data"
PS> make clean

# OR build and run with makefile
PS> make run ARGS="print"
```

## License
almost-blockchain-two uses the [MIT](https://choosealicense.com/licenses/mit/) license `:~)`
