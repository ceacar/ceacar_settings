nnoremap cc i<space><esc>
"noremap <Up> <NOP>
"noremap <Down> <NOP>
"noremap <Left> <NOP>
"noremap <Right> <NOP>

"noremap k j
"noremap j k

"syntax on
"highlight Comment ctermfg=darkgreen
"highlight BadWhitespace ctermbg=red guibg=darkred
"highlight highlight_id_cursor=darkgreen
" size of a hard tabstop
"set tabstop=4

" size of an "indent"
set shiftwidth=2

" a combination of spaces and tabs are used to simulate tab stops at a width
" other than the (hard)tabstop
set softtabstop=2

" make "tab" insert indents instead of tabs at the beginning of a line
set smarttab

"another color scheme
"Plugin 'jnurmine/Zenburn'
"Plugin 'altercation/vim-colors-solarized'



" always uses spaces instead of tab characters
"set expandtab


" better set to utf 8 while working with python
" set encoding=utf-8
" above command triiger error e518 on vim so disabled


"vundle setup
set nocompatible              " required
filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

" alternatively, pass a path where Vundle should install plugins
"call vundle#begin('~/some/path/here')


" let Vundle manage Vundle, required
Plugin 'gmarik/Vundle.vim'

" Add all your plugins here (note older versions of Vundle used Bundle instead of Plugin)

" navigate any char on the screen
Plugin 'easymotion/vim-easymotion'

" enable python indentation plugin
Plugin 'vim-scripts/indentpython.vim'

" ctrl + p will search on all files in current directory
Plugin 'kien/ctrlp.vim'

" nerd tree the gui acting as a file browser in vim
"Plugin 'scrooloose/nerdtree'

" ner tree with tab, use tab to invoke nerd tree
"Plugin 'jistr/vim-nerdtree-tabs'

" PEP8 code check
"Plugin 'nvie/vim-flake8'

" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on    " required


" Enable folding
set foldmethod=indent
set foldlevel=99
" hi Folded guibg=#A0A0A0
hi Folded ctermbg=242
" Enable folding with the spacebar
nnoremap <space> za

" F2 to toggle the paste and inv paste
nnoremap <F2> :set invpaste paste?<CR>
set pastetoggle=<F2>
set showmode

" F3 to toggle relative line number on and off
function! NumberToggle()
  "echo &relativenumber
  "echo &number
  if(&relativenumber == 1)
    set number
  else
    set relativenumber
  endif
endfunc



nnoremap <F3> :call NumberToggle()<cr>


set invnumber nornu

nnoremap <F4> :set invnumber nornu<cr>


" enable the folding plugin
"Plugin 'tmhedberg/SimpylFold'

" enable python indentation
au BufNewFile,BufRead *.py
    \ set tabstop=4 |
    \ set softtabstop=4 |
    \ set shiftwidth=4 |
    \ set textwidth=200 |
    \ set expandtab |
    \ set autoindent |
    \ set fileformat=unix |

au BufNewFile,BufRead *.js,*.html,*.css,*.yml
    \ set tabstop=2 |
    \ set softtabstop=2 |
    \ set shiftwidth=2 |

"au BufNewFile,BufRead *.c,*.cpp
"    \ set cindent |
"    \ set shiftwidth=2 |
"    \ set expandtab |

" flag excess space 
au BufRead,BufNewFile *.py,*.pyw,*.c,*.h match BadWhitespace /\s\+$/

" vim auto completion
Bundle 'Valloric/YouCompleteMe'
let g:ycm_autoclose_preview_window_after_completion=1
map <leader>g  :YcmCompleter GoToDefinitionElseDeclaration<CR>
"<leader> is a the primary key: if it's space it will be 'space-g' will go into definition

"code highlight
" disabled for conflict with you complete me   :w will pop out a whitebar on
" left side
" Plugin 'scrooloose/syntastic'

" python code pretty
let python_highlight_all=1

"hide pyc from nerd tree
"let NERDTreeIgnore=['\.pyc$', '\~$'] "ignore files in NERDTree


"let nerd tree start itself
"autocmd StdinReadPre * let s:std_in=1
"autocmd VimEnter * if argc() == 0 && !exists("s:std_in") | NERDTree | endif

" make backspaces more powerfull
set backspace=indent,eol,start

syntax on
colorscheme xedit 
highlight Comment ctermfg=darkgreen

"highlight Comment guifg=darkgreen
highlight BadWhitespace ctermbg=red guibg=darkred


vnoremap <silent> # :s#^#\##<cr>:noh<cr>
vnoremap <silent> -# :s#^\###<cr>:noh<cr>

" fast navigation to char

nmap s <Plug>(easymotion-s2)
nmap t <Plug>(easymotion-t2)

" prevents from wrapping long code
"set formatoptions-=t


"set ycm extra conf
"let g:ycm_global_ycm_extra_conf = '~/.vim/bundle/YouCompleteMe/third_party/ycmd/cpp/ycm/.ycm_extra_conf.py'
" let nerd tree tab open up on vim startup
"let g:nerdtree_tabs_open_on_console_startup=1

