#!/bin/bash
alias ll='ls -alF'
alias la='ls -A'
alias l='ls -CF'
alias lsize='ls -Salr'
alias less='less -R'
alias sl='ls'
alias scratch_notes='vim /tmp/temp_scratch.txt'
alias longestlines="awk 'length > max_length { max_length = length; longest_line = \$0;print max_length,\$0}'"
alias secondsSinceMidNightToTime='python ~/Documents/code/secondsSinceMidNightToTime.py'
alias reverseSecondsSinceMidNightToTime='python ~/Documents/code/reverseSecondsSinceMidNightToTime.py'
alias taqplace="cd /scratch/xiazi/taq"
alias taqscript="cd ~/taq/taqv22/ingest_taq"
alias taqsrc="cd ~/taq/taqv21/src"
alias stockEndTime="echo 57600000000"
alias stockStartTime="echo 34200000000"
#alias pypy="/opt/share/pypy/current/pypy-5.0-linux_x86_64-portable/bin/pypy"
alias lstime="ls -ltr"
alias phistory="cat ~/.persistent_history"
alias lssize="du -sh *"
alias lscolor="ls --color"
alias peekrawtaq="pypy ~/Documents/code/raw_taq_peek.py"
alias lsdir="ls -d */"
alias refreshbashrc='source ~/.bashrc'
alias dockerkillall='docker kill $(docker ps -q)'
alias vimnoplugin='vim -u NONE'
alias vimwithoutycm="vim --cmd 'set rtp^='~/.vimrd''"
uname_output="$(uname -s)"
alias vim="/usr/local/bin/vim"
alias ide="vim -u ~/.vimrc_ide"
#case "$uname_output" in
#  Linux*) alias vim="vim.nox" #only change vim to vim.nox in linux
#  #Darwin*) 
#esac
