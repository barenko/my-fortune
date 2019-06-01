# Fortune quotes installer

This project install a customized fortune file.

### Intructions
Edit the `my-fortune` text file with your quotes. Note that the lines with a single '%' char are the quote separators.

After edition completes, run the `fortune-apply.sh` script:

   $ ./fortune-apply.sh

This will install the `my-fortune` file into the fortune folder. Now the fortune will add the custom quotes into their database. If you wish to use only the custom ones, call the fortune with my-fortune as parameter:

   fortune my-fortune