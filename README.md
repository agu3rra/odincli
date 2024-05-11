# Perplexity CLI
This CLI application allows you to interact with [Perplexity.ai](https://www.perplexity.ai/) via 
the a terminal window.

## Table of contents
<!-- TOC depthfrom:2 depthto:2 -->

- [Table of contents](#table-of-contents)
- [Install](#install)
- [Usage](#usage)
- [Features](#features)
- [Why](#why)

<!-- /TOC -->

## Install

### macOS
> brew install pplcli

### linux
Simply copy the binary out of our releases section on this repository and add it to your `$PATH`.

## Usage
Simply access the built-in help menu to get instructions.
> ppl --help

At first run, make sure you execute `ppl configure` and provide your [Perplexity API key](https://docs.perplexity.ai/docs/getting-started).
If you ever need to edit it, it will be saved at `~/.pplcli/config`.

The simplest possible query is.
```bash
$ ppl "What is the meaning of life?"
The question "What is the meaning of life?" has intrigued philosophers, theologians, scientists, and thinkers throughout history, and it continues to be a central inquiry in the study of philosophy and existentialism.
```
## Features

### Patterns
We leverage Daniel Miessler's [fabric](https://github.com/danielmiessler/fabric) project to provide
some ready to use prompts that enhance the response quality you get out of Perplexity.
It works just like regular `fabric`, but targets `Perplexity` as its AI provider.

### Auto-copy to clipboard
Your last query response is automatically copied to the local clipboard, so that you're able to easily
paste it elsewhere (i.e.: your favorite note taking app).

### History
All your history (prompts and responses) are automatically saved to `~/.pplcli/history/`.

## Why
Why did I create this?  
ANS.: I wanted a project that could propel me to learn to code and structure applications in Golang and at the same time produced something that I could use on my daily workflow.
Being a heavy user of `Perplexity.ai`, it just made sense.
