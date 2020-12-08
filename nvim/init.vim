set nocompatible
filetype off

" set the runtime path to include Vundle and initialize it
set rtp+=~/.vim/bundle/vundle/
"set rtp+=~/.vim/bundle/powerline/powerline/bindings/vim
set rtp+=/usr/share/vim/vimfiles/
call vundle#begin()

" let Vundle manage Vundle.
Plugin 'VundleVim/Vundle.vim'
Plugin 'fatih/vim-go'
Plugin 'tpope/vim-fugitive'
Plugin 'airblade/vim-gitgutter'
Plugin 'vim-airline/vim-airline'
Plugin 'vim-airline/vim-airline-themes'
Plugin 'ctrlpvim/ctrlp.vim'
Plugin 'cespare/vim-toml'
Plugin 'Shougo/deoplete.nvim'
Plugin 'zchee/deoplete-go'
Plugin 'dmix/elvish.vim'

call vundle#end()
filetype plugin indent on

syntax on

" colorscheme Tomorrow-Night
" ------------------------------------------------------------------
" Solarized Colorscheme Config
" ------------------------------------------------------------------
""let g:solarized_termtrans=1    "default value is 0
""let g:solarized_termcolors=256    "default value is 16
""syntax enable
""set background=dark
""colorscheme solarized
" ------------------------------------------------------------------

" The following items are available options, but do not need to be
" included in your .vimrc as they are currently set to their defaults.

" let g:solarized_degrade=0
" let g:solarized_bold=1
" let g:solarized_underline=1
" let g:solarized_italic=1
" let g:solarized_contrast="normal"
" let g:solarized_visibility="normal"
" let g:solarized_diffmode="normal"
" let g:solarized_hitrail=0
" let g:solarized_menu=1

" --- airline ---
set encoding=utf-8
let g:airline_theme='powerlineish'
let g:airline#extensions#hunks#enabled=1
let g:airline#extensions#branch#enabled=1
let g:airline_powerline_fonts = 1
if !exists('g:airline_symbols')
  let g:airline_symbols = {}
endif

" powerline symbols
let g:airline_left_sep = ''
let g:airline_left_alt_sep = ''
let g:airline_right_sep = ''
let g:airline_right_alt_sep = ''
let g:airline_symbols.branch = ''
let g:airline_symbols.readonly = ''
let g:airline_symbols.linenr = '☰'
let g:airline_symbols.maxlinenr = ''


" --- base16 ---

" --- lucius
""colorscheme lucius

" --- hybrid ---
colorscheme hybrid


" --- settings ---
"filetype plugin indent on

" disable bell/beep
set noerrorbells visualbell t_vb=

" highlight current line
set cursorline

"" support for troff/groff
au BufNewFile,BufRead *.groff set filetype=groff

au FileType mail colorscheme peachpuff
au FileType mail set syntax=off

"" support for Go
""
" au FileType go setlocal formatoptions=cqrot1 tw=80 ai nofoldenable
au FileType go setlocal formatoptions=cqrot1 ai nofoldenable
let g:go_template_autocreate = 0
let g:go_fmt_command = "goimports"
let g:go_def_mode = "gopls"
let g:go_info_mode = "gopls"
let g:go_rename_command = "gopls"

" insert mode completion
set completeopt=menu,longest

" size of a hard tabstop
set tabstop=4

" size of an "indent"
set shiftwidth=4

" maximum with of text that is being inserted.
" set textwidth=80
set textwidth=0

" enable wrap:
set wrap linebreak nolist
" disable wrap:
"set nowrap

" a combination of spaces and tabs are used to simulate tab stops at a width
" other than the (hard)tabstop
set softtabstop=4

set hlsearch        " When there is a previous search pattern, highlight all
                    " its matches.
  
set incsearch       " While typing a search command,
					"  show immediately where the
                    " so far typed pattern matches.

set ignorecase      " Ignore case in search patterns.

set smartcase       " Override the 'ignorecase' option if the search pattern
					" contains upper case characters.

set ruler " Show the line and column number of the cursor position, separated by a comma.

"" enable the use of the mouse
set mouse=a

"" line numbers
set number

set showcmd       " show the command we're typing
set showfulltag   " show full completion tags
set showmode      " show the mode all the time

set splitright
set splitbelow

set showmatch     " show matching bracket

set linebreak  "" wrap at 'breakat' instead of last char
set magic      "" enable 'magic'

set scrolloff=50  "" keep at least X lines above/below
set sidescroll=2
set sidescrolloff=5

set nospell

"" accelerate scrolling
set ttyfast
set lazyredraw

" always show statusline (useful for powerline)
set laststatus=2

" set nolist
" set listchars="tab:\ \ ,trail:•,extends:#,nbsp:."
" call clearmatches()

" menu-completion
set wildmode=longest,list,full
set wildmenu

    " Clipboard {
    set clipboard+=unnamedplus
    " }

	set guioptions=aAimrLT

" bind C-Space to completion
"inoremap <Nul> <C-x><C-o>

" Use deoplete.
let g:deoplete#enable_at_startup = 1
let g:deoplete#enable_ignore_case = 0
"let g:deoplete#disable_auto_complete = 1
"let g:deoplete#complete_method = "omnifunc"
"if !exists('g:deoplete#omni#input_patterns')
"  let g:deoplete#omni#input_patterns = {}
"endif
" deoplete tab-complete
inoremap <expr><tab> pumvisible() ? "\<c-n>" : "\<tab>"
" inoremap <silent><C-n> <silent>my_complete_func()<CR>

" make sure we get the proper Ctrl+Left-arrow and +Right escape codes
"set <C-Left>=Od
"set <C-Right>=Oc

"" cycle thru buffers
:map <S-Right> :bnext<CR>
:map <S-Left> :bprevious<CR>

"" change the backup directory for backup and swap filesfiles (.swp)
"" the double // is here so the complete path is encoded into the file name
set backupdir=$HOME/.config/vim/vim/backup//,/tmp
set directory=$HOME/.config/vim/vim/backup//,/tmp

""" HTML
autocmd FileType html setlocal shiftwidth=2 tabstop=2

""" XML
autocmd FileType xml setlocal shiftwidth=2 tabstop=2
