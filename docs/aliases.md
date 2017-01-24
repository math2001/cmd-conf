Aliases are just awesome! And, it's pretty easy to create some. :wink: Here's how you can do it.

To create an alias, you need to use the `doskey` command:

```sh
doskey my-alias=my_command --to "run"
```

If you paste this in your command prompt, it'll work. But if you close it, and then reopen it, it
will be gone. Just an error.

So, you just need to tell to your command prompt to execute a bash script before opening. And this
bash script will set the aliases for you.

So, let's write that script!

!!! note
    If you've never wrote any windows batch, don't worry, this example is quiet easy to understand.

Create a file wherever you want (for my part, I created it in `C:/users/math/cmd-conf/runner.cmd`),
you just need to make sure you give it the `.cmd` extension.


```
:: this is a comment (a remark for the purist)
:: prevents every command of this script from being printed
@echo off

doskey gs=git status --short
doskey srf=subl %0
:: sublime runner file: opens this file in Sublime Text (%0 is a variable that contains the absolute
:: path to the current file)
```

So, now, you need to run this script each time you open the cmd. The `cmd` command has an option
`/k` which allows you to do this. It runs that script and opens the command prompt (with what the
script did taken into account, of course).

So, we're going to create a file (a shortcut) that is going to run `cmd /k <path_to_the_runner.cmd>`

- Find the `cmd.exe` To know where it is, just type in your command prompt, just type `where cmd`
your command prompt
- right click on it
- select `Send to → Desktop`
- (you can move that shortcut in a folder if you want to be organized)
- right click on that shortcut
- choose `Properties`
- get in the tab called `Shortcut`
- in the `Target` field, type `cmd /k ` plus the path to the file you created earlier
- you can set a shortcut to execute this file by focusing the field `Shortcut key` and pressing the
combination you want to use (<kbd>ctrl+shift+c</kbd> for example).
- hit `ok`

Now, double click on that file. It should open the command prompt, with the aliases you set working!

### Accepting new args

If you create an alias to script, and you try to give it other args from the command prompt, it
won't work. To make it do, you need to end your alias with `$*`

```bash
doskey hello=C:/my-python-scripts/say-hello.py $*
```

Now, this would works:

```
An\Ugly\prompt\text>hello --name matt
Hello Matt!
```

!!! tips
    You can add `cls` at the end of the file so that the first line in the command prompt isn't
    the first one.

    This command actually *clears* the screen. But welcome to the Windows command prompt:

    **The workaround world** :smile:
