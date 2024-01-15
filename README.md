# ASCII Art Generator
ASCII Art is a program that takes a string as input and generates a visually appealing graphical representation using ASCII characters. It offers a way to convert plain text into art by utilizing various ASCII characters to create patterns, shapes, and designs. The program allows users to customize the output by adjusting alignment, colors, and saving the result to a file.


# Flags
You can use these flags and arguments to customize the ASCII art output according to your preferences, including alignment, output file, colorization, and the reverse process.

* --align=<type>: Aligns the ASCII art to the specified position. The available alignment types are center, left, right, and justify. The default alignment is left. The program automatically adapts the graphical representation to fit the terminal size.

* --output=<fileName.txt>: Saves the ASCII art output to the specified file. The file name should be provided in the format <fileName.txt>. If this flag is not used, the output will be displayed in the terminal.

* --color=<color> <letters>: Colors the specified letters in the output. The <color> parameter defines the color for the letters, and <letters> specify the letter or letters to be colored. If <letters> is not specified, the entire string will be colored.

    - Color Notation
        The program supports different color notation systems, such as RGB 'rgb(x, y, z)', hexadecimal '#rrggbb', and ANSI colors. You can choose the notation system that suits your needs and specify the color using the appropriate format.

* --reverse=<fileName>: Reverses the ASCII art process and prints the original string from the specified file. The <fileName> parameter should be the name of the file containing the ASCII art representation.


# Usage
To use the ASCII Art program, follow these steps:

1. Install Go on your machine if you haven't already.
2. Clone the ASCII Art repository.
3. Run the program with the desired flags and arguments.

The program can be executed with a combination of the following arguments and flags:

```
$ go run main.go [OPTION] [STRING] [BANNER]
```

[OPTION]: Optional flags to customize the output. Use --align=<type> to specify the alignment type, --output=<fileName.txt> to save the output to a file, --color=<color> <letters> to color specific letters, and --reverse=<fileName> to reverse the ASCII art process and print the original string.
[STRING]: The string to be converted into ASCII art.
[BANNER]: Optional graphical template file to use for ASCII art representation.

* Note: The program accept multiple flags [OPTION] and/or [BANNER]. Additionally, the program can run with a single [STRING] argument. However, if the program receives any other format for the flags or arguments, it will display the following usage message.


- Example:
```
go run main.go --align=right --output=fileName.txt --color=red "something" shadow
```

## Usage Examples
Here are some examples of how to use the ASCII Art program:

1. Generate ASCII art with default settings and display it in the terminal:
```
$ go run main.go "Hello, World!"
```

2. Generate ASCII art and save the output to a file:
```
$ go run main.go --output=fileName.txt "Hello, World!"
```

3. Generate ASCII art aligned to the right and save the output to a file:
```
$ go run main.go --align=right --output=fileName.txt "Hello, World!"
```

4. Generate ASCII art aligned to the center with a specific banner:
```
$ go run main.go --align=center "Hello, World!" shadow
```

5. Generate ASCII art aligned to the left and save the output to a file:
```
$ go run main.go --align=left --output=fileName.txt "Hello, World!"
```

6. Generate ASCII art with a specific banner and save the output to a file:
```
$ go run main.go --align=justify --output=fileName.txt "Hello, World!" shadow
```

7. Generate ASCII art with colored letters:
```
$ go run main.go --color=red H --color=blue W "Hello, World!"
```

8. Generate ASCII art with colored letters:
```
$ go run main.go --color=red --align=right "Hello, World!"
```

9. Generate ASCII art with colored letters:
```
$ go run main.go --color=#663AE2 ! --align=center "Hello, World!"
```

10. Reverse the ASCII art process and print the original string:
```
$ go run main.go --reverse=fileName.txt
```

# Test File
There is a test file that test some of the program flags with banners, to run it use:
```
$ sh test.sh
```