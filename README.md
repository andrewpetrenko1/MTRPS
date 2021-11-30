# MTRPS
[![Run on Repl.it](https://repl.it/badge/github/andrewpetrenko1/MTRPS)](https://replit.com/@andrewpetrenko1/MTRPS)

“CSV Sorter” is a CLI application that allows sorting of its input presented as CSV-text.

## Technical details

Required features:
1. The application runs as a CLI application.
2. It reads STDIN line by line. The end of the input is an empty line.
3. Each line is a list of comma-separated values (CSV). Each value is considered as a piece of text. The
number of values is the same in each line.
4. The application sorts all lines alphabetically by the first value in each line using the Tree Sort
algorithm.
5. The application prints the result immediately, when the user ends to enter input text (presses <Enter> at a new line).
#### The application supports options:
    -i file-name Use a file with the name file-name as an input.
    -o file-name Use a file with the name file-name as an output.
    -h The first line is a header that must be ignored during sorting but included in the output.
    -f N Sort input lines by value number N.
    -r Sort input lines in reverse order.
