<h1 align="center">almost-blockchain-two</h1>
<p align="center"><img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/zhooda/almost-blockchain-two/Go?style=flat">
<img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/zhooda/almost-blockchain-two?color=crimson&label=lines%20of%20code&style=flat">
<img alt="go report card" src="https://goreportcard.com/badge/github.com/zhooda/almost-blockchain-two"></p>


<div style="margin-bottom: 2%"></div>

almost-blockchain-two is the sequel to [almost-blockchain](https://github.com/zhooda/almost-blockchain), a half-baked blockchain and cryptocurrency I created as my final project for high school computer science. 

<div style="margin-bottom: 2%"></div>

## Roadmap

- [x] ~~CLI to interact with chain~~
- [x] ~~Proof of work and validation~~
- [x] ~~Blockchain persistence with key-value store~~
- [x] ~~Transaction inputs and outputs~~
- [x] ~~Wallet module~~
- [x] ~~Digital signing, locking, and verification (via wallet module)~~
- [ ] Unspent Transaction Output (UTXO) persistence layer
- [ ] Merkle tree
- [ ] Networking module
- [ ] Variable PoW difficulty and reward distribution

## Getting Started

These instructions *should* get you a copy of the project up and running
on your local machine. There are no deployment strategies right now `:~)`

### Prerequisites

To run almost-blockchain-two you will need:

- [golang](https://golang.org/) v1.15+

On linux you will also need `build-essential` to be able to use the `make` command

### Installation

_Clone the repository or checkout with subversion to get a copy of almost-blockchain-two_

<details open>
<summary><b>Linux/macOS</b></summary>
<br>

```bash
$ git clone https://github.com/zhooda/almost-blockchain-two
$ cd almost-blockchain-two
$ make
```
</details>

<details>
<summary><b>Windows</b></summary>
<br>

```powershell
PS> git clone https://github.com/zhooda/almost-blockchain-two
PS> cd .\almost-blockchain-two\src
PS> go build -v -o ..\bin\abc2.exe main.go
PS> cd ..
```

When running abc2 on windows using the commands outlined below, replace `./bin/abc2` with `.\bin\abc2.exe` and you'll be good to go :)

</details>

---

## Usage

After you've compiled the program with the instructions above, you can make a basic chain and send some coins with the following commands!

_`create-chain`_ — creates a new chain if one doesnt exist
```bash
$ ./bin/abc2 create-chain --address "address0"
0000302ca94bc07dc2b4001bc4371c2ed6aa217b48a67b6483f494fe6c8e2727
genesis created
finished!

$ ./bin/abc2 create-chain --address "address1"
blockchain already exists
```

_`create-wallet`_ - creates a new wallet
```bash
$ ./bin.abc2 create-wallet
new address is: aBc1E8Npk8Vozw26ga8wARywfMCS7yJW43nUz
```

_`list-addrs`_ - lists all wallet addresses in wallet file
```bash
$ ./bin/abc2 list-addrs
aBc1E8Npk8Vozw26ga8wARywfMCS7yJW43nUz
```

_`get-balance`_ — returns coin balance of address
```bash
$ ./bin/abc2 get-balance --address "address0"
balance of address0: 100 dhicoins
```

_`send`_ — sends tokens from one address to another
```bash
$ ./bin/abc2 send --from "address0" --to "address1" --amount 20
00000191887c18d424ea139f1a8fdfb95348882127b247583084f24333b33aab
success!
```

_`print`_ — prints out full chain sorted by most recent block
```
$ ./bin/abc2 print
Prev. hash:        
Hash: 00000e9028f6a1427bf73f31db8affe9a76a20d5780772ff130fc19219c66d01
PoW: true
--- Transaction 997c4c58d83f329b08fab15f3e678f289b69fb189c30440720abb80eb482a116:
     Input 0:
       TXID:
       Out:       -1
       Signature:
       PubKey:    4669727374205472616e73616374696f6e2066726f6d2047656e65736973
     Output 0:
       Value:  100
       Script: 9c856302277d6eea6af195860093690d43d0c235
```

## License
almost-blockchain-two uses the [MIT](https://choosealicense.com/licenses/mit/) license `:~)`
