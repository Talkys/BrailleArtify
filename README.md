# BrailleArtify

This project aims to provide a simple way to create text art by converting black and white images into braille patterns. It is very efficient, but some improvements like parallelizing the character generation are possible.

## How to use

For images use ./braile -i image.png (it only works with jpegs and pngs)

For videos use ./braile -v frame_folder -f fps (the folder must contain only images and they must be in order. The fps value can handle non integers like 29.97)

## How to compile

The mod file can handle the compilation, so type go build . braile

## Results

To test with a simple image, we will use this Komi-san image, already optimized for generation. The program will work with any image with any size, but pre processing them makes the result better.

<img src="./img/komi.png" alt="Komi-san" style="width:100%;">

Now we can see the result after running the program:

[Click here to see the text file](./img/komi.txt)

The image will be inverted if you use light mode
