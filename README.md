# pdftextfmt

Use this script like this:

```bash
nvim -S pdftextfmt.vim
```

You can then bind this command to a keybinding using your OS, so that you can
easily run this command after copying text from a PDF.

Using Vim instead of Neovim also works, but make sure the specific version of
Vim you use has clipboard access. Try `vim --version | grep clipboard` and make
sure it says `+clipboard` or `+wayland_clipboard` or `+xterm_clipboard` or
something, and _doesn't_ says `-clipboard`.

## Instructions for installing on Windows 10/11

1. Download [`pdftextfmt.vim`](https://raw.githubusercontent.com/riceissa/pdftextfmt/master/pdftextfmt.vim) and save it on your computer somewhere. Make sure the extension is `.vim`, **not** `.vim.txt`. You might be able to do this in File Explorer by naming the file `pdftextfmt.vim.` with the trailing `.`.
2. [Install Neovim](https://github.com/neovim/neovim/blob/master/INSTALL.md) if you don't have it. Find the location of the `nvim.exe` (not `nvim-qt.exe`) and remember it. e.g. for me it's at `"C:\Program Files\Neovim\bin\nvim.exe"`.
3. Make a file called `pdftextfmt.cmd`. In it should be exactly one line that looks like:

   ```
   "C:\Program Files\Neovim\bin\nvim.exe" -S C:\Users\Issa\pdftextfmt.vim
   ```

   The first part should be your `nvim.exe` path. The second part (after the `-S`) should be the path to your `pdftextfmt.vim` file from step (1). (You'll need to surround it in quotes if there are spaces in the path.)

4. In File Explorer, right click on `pdftextfmt.cmd` and select "Create shortcut".
5. Right click on the shortcut (**not** the `.cmd` file), and select "Properties".
6. In the Properties window, in the Shortcut tab, there should be a text box called "Target". Keep whatever text is there, but add the following to the beginning of the box: `cmd.exe /C ` (there is a space after the `C`). So in the end, the box should say something like `cmd.exe /C "C:\Users\Issa\Desktop\pdftextfmt.cmd"`.
7. Now you can drag the shortcut to the taskbar and it will be pinned. From now on, you can just click the pinned shortcut on the taskbar to filter the contents of your clipboard. So your workflow will look something like: copy text from PDF using Ctrl-c → click the `pdftextfmt.cmd` shortcut icon on the taskbar → paste into whatever application you wanted to paste to using Ctrl-v (you don't even need to do Ctrl-Shift-v).
8. (Optional) You can right click on the icon on the taskbar, then right click on "pdftextfmt.cmd - Shortcut", then click Properties. In the Shortcut tab, there should be a dropdown that says "Run". You can select "Minimized" and then click Apply/OK. This will prevent a black command prompt window from flashing on the screen every time you use the program (unfortunately, it also seems to lose focus of the current window, so you have to copy, use pdftextfmt, then select the spot to paste in, rather than doing the last two steps in either order, because if you click the spot to paste in first, then you will have to click it again before you paste).
