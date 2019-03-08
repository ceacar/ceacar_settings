case $- in
    *i*) ;;
      *) return;;
esac

HISTCONTROL=ignoreboth

shopt -s histappend

HISTSIZE=1000
HISTFILESIZE=2000

shopt -s checkwinsize

[ -x /usr/bin/lesspipe ] && eval "$(SHELL=/bin/sh lesspipe)"

if [ -z "${debian_chroot:-}" ] && [ -r /etc/debian_chroot ]; then
    debian_chroot=$(cat /etc/debian_chroot)
fi

case "$TERM" in
    xterm-color) color_prompt=yes;;
esac


if [ -n "$force_color_prompt" ]; then
    if [ -x /usr/bin/tput ] && tput setaf 1 >&/dev/null; then
        color_prompt=yes
    else
        color_prompt=
    fi
fi

if [ "$color_prompt" = yes ]; then
    PS1='${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '
else
    PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w\$ '
fi
unset color_prompt force_color_prompt

case "$TERM" in
xterm*|rxvt*)
    PS1="\[\e]0;${debian_chroot:+($debian_chroot)}\u@\h: \w\a\]$PS1"
    ;;
*)
    ;;
esac

# enable color support of ls and also add handy aliases
if [ -x /usr/bin/dircolors ]; then
    test -r ~/.dircolors && eval "$(dircolors -b ~/.dircolors)" || eval "$(dircolors -b)"
    alias ls='ls --color=auto'
    #alias dir='dir --color=auto'
    #alias vdir='vdir --color=auto'

    alias grep='grep --color=auto'
    alias fgrep='fgrep --color=auto'
    alias egrep='egrep --color=auto'
fi

# Alias definitions.
if [ -f ~/.bash_aliases ]; then
    . ~/.bash_aliases
fi
# custom script
if [ -f ~/.bash_home_brew_script ]; then
    . ~/.bash_home_brew_script
fi
# enable programmable completion features 
if ! shopt -oq posix; then
  if [ -f /usr/share/bash-completion/bash_completion ]; then
    . /usr/share/bash-completion/bash_completion
  elif [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
  fi
fi


# directory color display
LS_COLORS=$LS_COLORS:'di=36:' ; export LS_COLORS

UTILITY_PATH=/home/ceacar/projects/ceacar_settings/utility
PATH=$PATH:~/Documents/code
PATH=$PATH:/usr/local/spark/bin
PATH=$PATH:/usr/local/go/bin
#add support for golang, install please check go.cheatsheet
PATH=$PATH:/usr/local/go/bin
PATH=$PATH:$UTILITY_PATH/bash_lib
PATH=$PATH:/opt/share/jvm/current_java_jdk/bin
export PATH
export PYTHONPATH=/home/ceacar/projects/ceacar_settings/excalibur:$PYTHONPATH
export JAVA_HOME=/opt/share/jvm/current_java_jdk
if $(uname -a | grep -q "Darwin" && echo "true" || echo "false");then
  export ceacar_setting='/Users/ceacar/projects/ceacar_settings'
else
  export ceacar_setting='/home/ceacar/projects/ceacar_settings'
fi
export notes=$ceacar_setting/notes

alias sshmycellphone='ssh xiazi@192.168.86.244 -p 8022'
# added by Anaconda3 2018.12 installer
# >>> conda init >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$(CONDA_REPORT_ERRORS=false '/home/ceacar/anaconda3/bin/conda' shell.bash hook 2> /dev/null)"
if [ $? -eq 0 ]; then
    \eval "$__conda_setup"
else
    if [ -f "/home/ceacar/anaconda3/etc/profile.d/conda.sh" ]; then
        . "/home/ceacar/anaconda3/etc/profile.d/conda.sh"
        CONDA_CHANGEPS1=false conda activate base
    else
        \export PATH="/home/ceacar/anaconda3/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda init <<<

# added by Anaconda3 installer
export PATH="/home/ceacar/anaconda3/bin:$PATH"
. /home/ceacar/anaconda3/etc/profile.d/conda.sh
