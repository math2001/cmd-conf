# -*- encoding: utf-8 -*-

import subprocess
import os
import colorama

colorama.init(autoreset=True)


def get_current_git_branch():
    if os.path.exists('.git'):
        # from http://stackoverflow.com/a/12142066/6164984
        cmd = subprocess.Popen(['git', 'rev-parse', '--abbrev-ref', 'HEAD'], stdout=subprocess.PIPE,
                                 stderr=subprocess.PIPE)
        branch = cmd.stdout.read().decode().strip()
        if not branch:
            return
        cmd = subprocess.Popen(['git', 'status', '--short'], stdout=subprocess.PIPE)
        if cmd.stdout.read().decode() != '':
            return colorama.Fore.RED + colorama.Style.BRIGHT + branch
        return colorama.Fore.GREEN + colorama.Style.BRIGHT + branch


def get_pwd():
    pwd = os.getcwd().replace(os.path.expanduser('~'), '~').replace(os.path.sep, '/') \
    .replace('C:/', '/')
    return pwd + ('/' if pwd == '~' else '')


def break_line(line, visible_string_length):
    if visible_string_length > 20:
        line += '\n'
    else:
        line += ' '
    return line


pwd = get_pwd()
line = colorama.Fore.MAGENTA + colorama.Style.BRIGHT + pwd + colorama.Style.RESET_ALL
visible_string_length = len(pwd)
branch = get_current_git_branch()
if branch:
    line += ' on ' +  branch + colorama.Style.RESET_ALL
    visible_string_length += 4 + len(branch)


# CSW: ignore
print(break_line(line, visible_string_length) + '$ ', end='')
