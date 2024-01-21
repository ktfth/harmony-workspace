# Harmony

## Description

Harmony project is a simples application microservice, to register prompts of LLM models.

## Usage

The first thing you need to do is to use `just` to use recipes, when you are done, start like that:

```sh
cd harmony-engine
```

Start the drums...

```sh
just weaver_generate update run
```

After update the project with needed packages, if running the application in a free context.

```sh
just weaver_generate run
```

This complete the process to get up and running the application.

### Testing the application

Just made this simple:

```sh
just weaver_generate update build test
```

Or

```sh
just weaver_generate build test
```
