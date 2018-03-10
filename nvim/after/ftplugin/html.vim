setlocal shiftwidth=2
setlocal tabstop=2

augroup autoindent
    au!
    autocmd BufWritePre *.html :normal migg=G`i
augroup End


