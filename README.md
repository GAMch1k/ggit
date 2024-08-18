# GGIT - GAMch1k's git

This is just a git copy written on golang for me to gain experience


## Some Explanation

- .ggit folder is same as .git, but there is also .ggitignore file, for now there is only global one 

- .ggitignore comments starts with ```#```

- .ggitignore cheatsheet:
    - ```dir/``` - ignores full directory
    - ```*.ext``` - ignores file extension
    - ```*``` - can also be used to ignore specific literals in the filename, example: ```*test``` will ignore everything with ```test``` word in it
    - ```.keep``` - file used to keep folder

- Other explanations will be added when I will add something new

## Available Commands

```ggit init``` - initialize empty repository, nothing new

```ggit add``` - for now adding all code in the project

```ggit commit``` - nothing here for now

```ggit -h``` - list of commands