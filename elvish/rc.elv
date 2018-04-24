
## enable elvish package manager
use epm

epm:-lib-dir = ~/.elvish/lib


use epm

epm:install &silent-if-installed=$true   \
  github.com/zzamboni/elvish-modules     \
  github.com/zzamboni/elvish-completions \
  github.com/zzamboni/elvish-themes      \
  github.com/xiaq/edit.elv               \
  github.com/muesli/elvish-libs

use re

use readline-binding

edit:insert:binding[Alt-Backspace] = $edit:kill-small-word-left~

edit:insert:binding[Alt-d] = { edit:move-dot-right-word; edit:kill-word-left }

use github.com/muesli/elvish-libs/git

## aliases

use github.com/zzamboni/elvish-modules/alias

alias:new ls e:ls --color=auto
alias:new ll e:ls -lh --color=auto
alias:new more less

## completions
use github.com/zzamboni/elvish-completions:git

use github.com/zzamboni/elvish-completions:cd
use github.com/zzamboni/elvish-modules/dir

alias:new cd &use=[github.com/zzamboni/elvish-modules/dir] dir:cd

use github.com/xiaq/edit.elv/compl/go
go:apply

E:LESS = "-i -R"

E:EDITOR = "vim"

E:LC_ALL = "en_US.UTF-8"

use github.com/zzamboni/elvish-modules/util

## customize prompt
#use github.com/muesli/elvish-libs/theme/powerline
use github.com/zzamboni/elvish-themes/chain
chain:segment-style[timestamp] = yellow
chain:prompt-segments = [timestamp su dir arrow]
chain:prompt-pwd-dir-length = 3
chain:bold-prompt = $true

-exports- = (alias:export)

## for direnv
use direnv
