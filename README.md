# Fortune quotes installer

This project install a customized fortune file.


### Instalation TL;DR

Just run it:

```bash
    curl -fsSL https://raw.githubusercontent.com/barenko/my-fortune/master/fortune-install.sh | sh
```

### Instalation, the hard way

Edit the `my-fortune` text file with your quotes. Note that the lines with a single '%' char are the quote separators.

After edition completes, run the `fortune-apply.sh` script:

```bash
    $ ./fortune-apply.sh
```

This will install the `my-fortune` file into the fortune folder. Now the fortune will add the custom quotes into their database. If you wish to use only the custom ones, call the fortune with my-fortune as parameter:

```bash
    $ fortune my-fortune
```

