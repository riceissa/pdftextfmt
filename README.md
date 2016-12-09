# pdftextfmt

Format text copied from PDF files.

    $ go build pdftextfmt.go
    $ cat <<'EOF' | ./pdftextfmt
    > This text has  uneven spacing,
    > hyphen-
    > ation, and trailing whitespace.  
    > EOF
    This text has uneven spacing, hyphenation, and trailing whitespace.

Mappings for Vim:

    nnoremap Q Vip:!pdftextfmt<CR>gqq
    vnoremap Q :!pdftextfmt<CR>gqq
