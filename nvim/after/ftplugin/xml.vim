setlocal shiftwidth=2
setlocal tabstop=2

augroup autoindent
    au!
    autocmd BufWritePre *.xml :normal migg=G`i
augroup End
