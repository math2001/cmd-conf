@echo off

doskey gl=git log --oneline --all --graph --decorate -10 $*
doskey gll=git log --oneline --all --graph --decorate $*
doskey glo=git log --oneline $*
doskey gs=git status --short $*
doskey st=cd "%APPDATA%\Sublime Text 3\Packages"$*
doskey stu=cd "%APPDATA%\Sublime Text 3\Packages\User"$*
doskey sdf=subl %0
doskey ls="C:\Program Files\Git\usr\bin\ls.exe" -F -X --color --ignore "NTUSER.DAT{*" $*
doskey PackageChecker=python "%APPDATA%\Sublime Text 3\Packages\PackageChecker\PackageChecker.py" $*
doskey gst=python "C:\python\tools\git-semver-tag\git-semver-tag.py" -s $*

title Terminal
cls
echo.
:: loop: ask for a command, execute this command
:cmd
python "C:/users/math/cmd-conf/echo-prompt.py"
set cmd=
set /p cmd=
%cmd%
echo.
goto cmd
