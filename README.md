# doit

I created this tool to help me find what notes I need to clean before saving for future reading.
This works with Emacs, so its possible to jump between files


## Install

```console
go install github.com/BrunoTeixeira1996/doit@latest
```

## Use

```console
help [shows help]
list [lists all todos of a dir recursively being .md the default extension] [-type][-help]
```

## Example

- List only `.md` files

```console
doit list
```

- List only `.tex` files

```console
doit list -type .tex
```
- List `.tex` and `.md` files

```console
doit list -type .tex,.md
```

![image](https://user-images.githubusercontent.com/12052283/148767490-3518e67b-d645-4144-b119-951162c6b0eb.png)


## Plan to do

- Integrate with Asana []
- Integrate with Notion []
