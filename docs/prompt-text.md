The prompt text is the text that is shown before you can type a command. By default, it is
`path>`

This is really bad, right? You get into a nested folder, and you don't have any room left to type
your command (you get this line break, which is really annoying)

But, were you aware that you could change it? Yes, and that's really easy!

```
prompt [text]
```

The command `prompt` changes it for you (all it does is change the `prompt` variable).


    ```
    prompt /?
    ```

So, for example:

```
D:\iscuting\prompt\text>prompt Your command^>

Your command> echo Cool!
Cool!

Your command>
```

!!! note
    There is a `^` before the `>` because you need to escape it. And, in windows batch, you the
    "escaper" is the `^` symbol.

But, this prompt is almost worst, right? But here's when it gets pretty cool:

```
C:\Users\math>prompt $P$_$$

C:\Users\math
$ cd Desktop

C:\Users\math\Desktop
$
```

This one uses the "variables". Here is the list

| Variable Name |                 Value                 |
|---------------|---------------------------------------|
| `$A`          | `&` (Ampersand)                       |
| `$B`          | `&#124;` (pipe)                       |
| `$C`          | `(` (Left parenthesis)                |
| `$D`          | Current date                          |
| `$E`          | Escape code (ASCII code 27)           |
| `$F`          | `)` (Right parenthesis)               |
| `$G`          | `>` (greater-than sign)               |
| `$H`          | Backspace (erases previous character) |
| `$L`          | `<` (less-than sign)                  |
| `$N`          | Current drive                         |
| `$P`          | Current drive and path                |
| `$Q`          | `=` (equal sign)                      |
| `$S`          | ` ` (space)                           |
| `$T`          | Current time                          |
| `$V`          | Windows version number                |
| `$_`          | Carriage return and linefeed          |
| `$$`          | `$` (dollar sign)                     |

!!! tip
    To display the full help, you can use `prompt /?`

But, it still can be improved.

Say, you want to display the the current git branch in it? With this solution, you're stuck.

So, here's the workaround: We're going to create a loop in our runner file (see the
[aliases page](aliases.md) if you don't know how to create it) that is going to

- display the prompt text *dynamically*
- ask for a command to run
- run that command
- go back up

So, in fact, the "default" (so, even the one you're going to set with the `prompt ....` command)
prompt is never going to be rendered.

And here's how it looks in batch

```dosbatch
:: your aliases
doskey alias=command
doskey alias=command
doskey alias=command

:: create the label "cmd"
:cmd
:: render your custom prompt
echo My prompt
:: reset the command to run
set cmd=
:: ask the command to run
set /p cmd=
:: run the command
%cmd%
:: print a new line
echo.
:: go back up (to the label "cmd")
goto cmd
```


So, in this case, it'd be worst, because the prompt text would be `My prompt`. But, the awesome part
is that you can call an external program, which is going to print the prompt!

!!! note
    This solution requires you to know how to program in at least one programming language. Because
    I love Python, this is the language I'm going to use, but you can use anything.

So, in my case, I've replaced the `echo My Prompt` by
`python "C:/users/math/cmd-conf/echo-prompt.py"`

!!! warning
    The path to your file that is going to render the prompt text *must* be absolute.

And, here is the content of my file: https://github.com/math2001/cmd-conf/blob/master/echo-prompt.py

This python script uses [`colorama`][colorama] which is a cross-platform python library for
outputting text in color in your terminal.

Note that you might want to optimise this file, because it's going to be called *each* time you're
going to type a new command!

So, build your own script to fit your needs, you're completely free now!

[colorama]: https://pypi.python.org/pypi/colorama
