This part is the really awesome in my opinion. We're going to include almost every commands that
are available in Unix right in your *Windows* command prompt!

These days, almost everybody has `git` on it's system. If it's not the case, then

1. you should learn how to use it
2. and... you won't be able to have all the unix commands

But if you do, you'll be amazed!

If you open the `git bash`, you'll see that you have access to almost every Unix command working.
`ls`, `find`, `curl`, `vim`, and many many more.

Well, this means that, those commands, there somewhere on your computer right?

So, you can tell to your command prompt to use them!

!!! warning
    You'll need to have administrator right

They're in two different folders, but both are where you installed `git`.

!!! note
    In my case, it's in `C:/Program Files/Git`

So, if you have a look in `C:/Program Files/Git/usr/bin/`, there you have it! Paradise!! Plenty
of Unix commands. You should find `ls.exe` in here. So, if it's the case, it's the right folder,
**copy its path** (click on the address bar in explorer and <kbd>ctrl+c</kbd>).

So, you need to **add this folder to your path environment variable**.

- open the control panel
- get in `System and Security`
- `System`
- on the right, click on `Advanced System Settings`
- click on `Environment Variables`
- in the **bottom** part (*System* variable), find the `Path` variable, and click on edit.

At the *begginning* of the value, paste your path, and add a semi-colon `;` if there is None

!!! note
    All your paths in this `Path` variable MUST be separated with semi-colons.

!!! warning
    On windows 10, you'll get a nice interface with rows. In this case, you don't need to worry
    about the semi-colons. Just click on an empty row and paste (or click on a button `Add`, I'm
    not sure about that one).

    You need to make sure that your path is the first one (at least before `C:/windows/system32`,
    see bellow to know why)

Got here? Great! You just added the commands `ls`, `find`, `vim`, etc!

!!! note
    You'll need to click OK on each windows opened a restart the command prompt to *actually* have
    them.

Now adding the other folder should be a piece of cake! It's the same thing, just with an other
folder. This folder should be `C:\Program Files\Git\mingw64\bin`. You should find the `curl`
command in there!

So, follow the above steps again to add this path!

Now, restart your command prompt, and try to type `ls`. It works! Much better that `dir` right?

To learn how it some of those awesome commands, I recommend watching
[those tutos][unix-commands-tuto]

!!! note
    Usually, when you want to add files to your path, you edit the `user` variables (the *top*
    section). Here, we need to make the Unix commands override the default windows' ones (such
    as `find`). And the file that is executed is the first found.

    For example, the `find` command is in the `C:/Windows/System32` folder, and because it was in
    the `System` variable, you need to specify your path *before* it.

    If you didn't understood, it doesn't really matter, you'll understand it when you'll need it,
    you'll see :wink:

[unix-commands-tuto]: https://www.youtube.com/watch?v=j6vKLJxAKfw&list=PL-osiE80TeTvGhHkpvfmKWOiIPF8UVy6c
