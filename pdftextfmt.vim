" Paste text
normal! "+P

" Remove beginning whitespace
:%s/^\s\+//e

" Remove trailing whitespace
:%s/\s\+$//e

" Join hyphens
:%s/-$\n//e

" Remove inconsistent spacing
:%s/\s\+/ /ge

" Add a blank line at the end of the buffer so the next command doesn't fail
normal! Go

" Join paragraphs. From :help edit-paragraph-join
:g/\S/,/^\s*$/join

" Delete blank lines
:g/^$/delete

" Create paragraph breaks
:%s/\n/\r\r/e

" Two extra blank lines will be created by the previous command. The following
" removes them.
:normal! Gdddd

" Copy text to clipboard
normal! ggvGg_"+y

:quit!
