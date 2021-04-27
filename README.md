# Ren'Py graph vizualiser

[![Go Reference](https://pkg.go.dev/badge/pkg.amethysts.studio/renpy-graphviz.svg)](https://pkg.go.dev/pkg.amethysts.studio/renpy-graphviz)
[![Go Report Card](https://goreportcard.com/badge/pkg.amethysts.studio/renpy-graphviz)](https://goreportcard.com/report/pkg.amethysts.studio/renpy-graphviz)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ewenquim/renpy-graphviz/Distribute%20executable)

This is a tool written in Go that allows you to **visualise the routes** of your story.

![](./data/the_question.jpg)
_Routes of the Question, the classic Ren'Py example_

- [Ren'Py graph vizualiser](#renpy-graph-vizualiser)
  - [Examples](#examples)
  - [How to use](#how-to-use)
    - [Go library](#go-library)
  - [Online demo](#online-demo)
  - [Tags](#tags)
    - [TITLE & GAMEOVER](#title--gameover)
    - [BREAK](#break)
    - [IGNORE](#ignore)
    - [SKIPLINK](#skiplink)
  - [Limitations](#limitations)
  - [LICENSE](#license)

## Examples

![](./data/DDLC_extract.png)
_[Doki Doki Litterature Club](https://ddlc.moe/) will no longer have secrets for you!_

![](./data/CXVL_extract.png)
_An extract from my personnal VN, [Coalescence](https://play.google.com/store/apps/details?id=com.coal). You can't imagine handling a heavy VN like this one without graphic tools... (the labels aren't blurred on the real image)_

## How to use

- [**Download**](https://github.com/EwenQuim/renpy-graphviz/releases) latest version
- **Move** the program in your game folder
- **Run it** from the command line
  - you might have to give yourself the permissions: don't worry my program isn't a virus ! Run `chmod +x renpy-graphviz*` on Unix.
- `renpy-graphviz.png` just appeared, **enjoy** !

### Go library

If you are a Go user and want to integrate this in a lib/program, it is totally possible. The `/parser` module is very powerful.

```
go get pkg.amethysts.studio/renpy-graphviz
```

## Online demo

You can test this tool in the browser. If you really want to get `.png` files, please download. Note that I will not maintain this website, it is not guaranteed to represent the library 100%.

https://renpy.amethysts.studio

## Tags

I made a tag system to enforce some behaviours. For example

```renpy
label chapter_1: #renpy-graphviz: TITLE
```

Before tags, you must write `renpy-graphviz` in a comment to ensure there are no collision with existing words in your VN.

Here are the tags available

- `TITLE`: style for chapters
- `BREAK`: breaks the current flow, for parallel labels for example
- `IGNORE`: ignores the current label. Jumps to this label still exist
- `GAMEOVER`: style for endings
- `SKIPLINK`: avoid long arrows by creating a "shortcut" - read the doc below before using

Case, spaces and separators are handled very loosely, don't worry about it.

### TITLE & GAMEOVER

Set some styles

![Example of GAMEOVER & TITLE tags](./data/example-title-gameover.png)

### BREAK

Cancels any "guessed link".

```renpy
label one:
  "blah blah"

label two:
    "bla bla"

# renpy-graphviz: BREAK
label three:
    "the end"

```

|            without tag             |             with tag              |
| :--------------------------------: | :-------------------------------: |
| ![](data/example-break-before.png) | ![](data/example-break-after.png) |

### IGNORE

Ignore the current line. If this is a jump to a label that isn't ignored, the label will still appear on the graph but not the arrow that should go towards it.

### SKIPLINK

Avoids long arrows by creating another label with the same name. Beware, the label can't have any children and is marked by an asterix to show it is a copy.

```renpy
label one:
    if condition:
        jump six # renpy-graphviz: SKIPLINK
    else:
        pass

label two:
label three:
label four:
label five:
label six:
```

|              without tag              |               with tag               |
| :-----------------------------------: | :----------------------------------: |
| ![](data/example-skiplink-before.png) | ![](data/example-skiplink-after.png) |

## Limitations

This require your VN to be structured in a certain way, so it's possible that it isn't perfect for you. Feel free to raise an issue [here](https://github.com/EwenQuim/renpy-graphviz/issues), or to change your VN structure.

The program works only for scripts that do not stack call statement, i.e. the program expects a `break` statement before any other `label`/`call` if you used call to get there.

Works:

```renpy
label start:
  eileen "hello"
  call second
  eileen "I'm back"

# renpy-graphviz: BREAK <- recommended here but not mandatory, see Tags section
label second:
  eileen "inside a CALL statement"
  break # <- works !!!
```

Does NOT work:

```renpy
label start:
  eileen "hello"
  call second
  eileen "I'm back"

# renpy-graphviz: BREAK <- recommended here but not mandatory, see Tags section
label second:
  eileen "inside a CALL statement"
  call / jump third_label # <- Isn't taken into account !!!!

label third_label
```

## LICENSE

This program is free and under the [AGPLv3 license](https://www.gnu.org/licenses/agpl-3.0.en.html).

Beware, if you use this program, you must **credit it somewhere on your game**.

> Used Renpy Graph Vizualiser from EwenQuim

Enjoy!
