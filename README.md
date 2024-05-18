# GPT Terminal
This CLI application (`gpterm`) allows you to interact with [Chat GPT](https://chatgpt.com/) via 
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
> brew install gpterm

### linux
Simply copy the binary out of our releases section on this repository and add it to your `$PATH`.

## Usage
Simply access the built-in help menu to get instructions.
> gpterm --help

At first run, make sure you execute `gpterm configure` and provide your [ChatGPT API key](https://platform.openai.com/api-keys).
If you ever need to edit it, it will be saved at `~/.gpterm/config`.

The simplest possible query is.
```text
$ gpterm "What is the meaning of life?"
The question "What is the meaning of life?" has intrigued philosophers, theologians, scientists, 
and thinkers throughout history, and it continues to be a central inquiry in the study of 
philosophy and existentialism.
...
```
## Features

### Patterns
We leverage Daniel Miessler's [fabric](https://github.com/danielmiessler/fabric) project to provide
some ready to use prompts that enhance the response quality you get out of the AI.

### Auto-copy to clipboard
Your last query response is automatically copied to the local clipboard, so that you're able to easily
paste it elsewhere (i.e.: your favorite note taking app).

### History
All your history (prompts and responses) are automatically saved to `~/.gpterm/history/`.

## Why
Why did I create this?  
ANS.: I wanted a project that could propel me to learn to code and structure applications in Golang and at the same time produced something that I could use on my daily workflow.
Being a heavy user of `Chat GPT`, it just made sense.
