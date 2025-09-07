# The Nature of Code: Go + Ebitengine

### 🩷 Featured in the official [Nature of Code Ports to Other Languages](https://natureofcode.com/resources/) 🩷

This repository contains ports of the [p5.js](https://p5js.org/) sketches from [*The Nature of Code*](https://natureofcode.com/) book by [Daniel Shiffman](https://shiffman.net/), published by No Starch Press® Inc., licensed under [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).

This project is non-commercial and distributed under the same license.

The examples are rewritten in [Go](https://golang.org/) using the [Ebitengine](https://github.com/hajimehoshi/ebiten) game engine.

The original p5.js examples can be found at [natureofcode.com/examples](https://natureofcode.com/examples/) and in the [noc-book-2 GitHub repository](https://github.com/nature-of-code/noc-book-2/tree/main/content/examples).

## How To Run

You will need to have [Go](https://golang.org/) installed (version 1.24.5 or newer).

Each example is in its own directory with a `main.go` file. To run an example, navigate to its directory and run:

```sh
go run .
```

## No Additional Libraries

To keep things simple, I've chosen to use only [Go](https://golang.org/) and [Ebitengine](https://github.com/hajimehoshi/ebiten) for these examples. Where necessary, I've written my own simplified ported versions of [p5.js](https://p5js.org/) math functions and separate physics library functions, rather than relying on additional third-party libraries.

If you adapt these examples, feel free to modify the ported functions or substitute them with functions from another library.
