# ReadMe for ASCII Art Colorization Program

## Description
This program takes user input from the command line to generate ASCII art in different styles and colors. It allows the user to choose between several banner styles (e.g., "standard", "shadow", "thinkertoy") and supports various ways of customizing text color using RGB, hexadecimal, or predefined named colors.

## Features
- **Text Colorization**: Supports text colorization using RGB values, hexadecimal color codes, or predefined color names.
- **Banner Styles**: Supports several banner styles, including "standard", "shadow", "thinkertoy", and "golden". You can also specify your own banner file.
- **Sub-string Highlighting**: Allows highlighting specific substrings in the text with custom colors.
- **Command-Line Arguments**: Flexible argument parsing to handle various combinations of options for color, banner style, and text input.

## Prerequisites
- Go (Golang) installed on your system.
- Command-line terminal to run the program.

## Usage
Run the program with the following command-line arguments:

```
go run main.go [OPTIONS] [TEXT] [BANNER]
```

### Arguments
- **[OPTIONS]** (optional):
  - `--color=rgb(r,g,b)`: Set the text color using RGB values (e.g., `--color=rgb(255,0,0)` for red).
  - `--color=#RRGGBB`: Set the text color using a hexadecimal code (e.g., `--color=#FF0000` for red).
  - `--color=[color_name]`: Use one of the predefined colors (e.g., `--color=red`, `--color=blue`, `--color=green`, etc.).
  
- **[TEXT]**: The text that you want to display, with optional newline characters (`\n`).
  
- **[BANNER]**: The ASCII banner style (e.g., "standard", "shadow", "thinkertoy", "golden"). You can also provide a custom banner file (must have a `.txt` extension).

### Examples
1. **Colorize text with RGB values**:
   ```bash
   go run main.go --color=rgb(255,0,0) "Hello, World!" standard
   ```
   This will print "Hello, World!" in red using the "standard" ASCII banner.

2. **Colorize text with a hex color code**:
   ```bash
   go run main.go --color=#FF0000 "Hello, World!" standard
   ```
   This will print "Hello, World!" in red using the "standard" ASCII banner.

3. **Using predefined color names**:
   ```bash
   go run main.go --color=blue "Hello, World!" shadow
   ```
   This will print "Hello, World!" in blue using the "shadow" ASCII banner.

4. **Highlight a substring**:
   ```bash
   go run main.go --color=green "Hello, World!" "This is a test" -- "World"
   ```
   This will print "Hello," in the default color and "World" in green, based on the matching substring "World."

## File Structure
- **main.go**: The main Go program containing the logic to process the command-line arguments and generate the ASCII art with the specified colors.
- **standard.txt**: Default ASCII banner file. This file contains a template for the "standard" banner. Other banner styles like "shadow", "thinkertoy", etc., should also have their respective text files.

## How It Works
1. **Argument Parsing**: The program reads the command-line arguments and determines the color, banner style, and text input. The color and banner style are applied to the text accordingly.
2. **Banner Display**: Based on the user's choice, the appropriate banner style is read from a `.txt` file and applied to the text.
3. **Text Colorization**: The program converts the specified color (either RGB, hexadecimal, or predefined) into an ANSI escape code, which is then used to print the text in the desired color.
4. **Substring Highlighting**: If a substring is specified, it is highlighted in the chosen color.

## Customization
You can add your own custom banner files by creating `.txt` files with the desired ASCII art. Just make sure to specify the file's name when running the program.

## Error Handling
- If an invalid number of arguments is provided, the program will display an error message and exit.
- If the specified banner file does not exist, the program will notify the user and exit.
- If an invalid color is specified, the program will notify the user and exit.

## License
This project is open-source and available under the MIT License.
