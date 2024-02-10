### Build Your Own wc Tool

This challenge is to build your own version of the Unix command line tool wc!

The Unix command line tools are a great metaphor for good software engineering and they follow the Unix Philosophies of:

- Writing simple parts connected by clean interfaces - each tool does just one thing and provides a simple CLI that handles text input from either files or file streams.
- Design programs to be connected to other programs - each tool can be easily connected to other tools to create incredibly powerful compositions.

Following these philosophies has made the simple unix command line tools some of the most widely used software engineering tools - allowing us to create very complex text data processing pipelines from simple command line tools. There’s even a [Coursera course on Linux and Bash for Data Engineering](https://gb.coursera.org/learn/linux-and-bash-for-data-engineering-duke).

You can read more about the Unix Philosophy in the excellent book The [Art of Unix Programming](http://www.catb.org/~esr/writings/taoup/html).

### Usage

```bash
# build the project
make build
```

```bash
bin/wc -l <file_name>
# or
cat <file_name> | bin/wc -l
```
