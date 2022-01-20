# FujiSimuRecipesGen
A Go program to generate Fujifilm simulations from recipes of such simulations in CSV

A lot of persons write their recipes of Fujifilm film simulation within a spreadsheet kept up-to-date. Usually those come from 
[Fuji X Weekly](https://fujixweekly.com/). Then, they set themselves the properties of the recipe either directly into their own Fujifilm camera or
via the Fujifilm X-Raw Studio tool to create a custom film simulation in the camera.

The `FujiSimuRecipesGen` program aims to automatically generate Fujifilm simulations in FP1 format from such a spreadsheet. These FP1 files then can be 
copied or moved into the location of the film simulations for Fujifilm X-Raw Studio so that you can transfer them directly to your camera with the Fujifilm tool.

## Information

You can find more information about the context of this program, the recipes in a spreadsheet, and the FP1 film simulation with Fujifilm X-Raw Studio in the website below (it is in French):

https://thomashammoudi.notion.site/Documentation-preset-Fujifilm-a68e0d6ce170416987ad458475cca9ec

For information how to create custom film simulations with Fujifilm X-Raw Studio and to transfer them into your camera, you can read the tutorial below:

https://photography.tutsplus.com/tutorials/profile-raw-photos-in-xraw--cms-37087

## Requirements

* Create a settings file in YAML in which you specify the properties of both your camera and of Fujifilm X-Raw Studio. 
To know the required properties, please look at `example/settings.yaml` file.
* Have your spreadsheet in CSV format following the syntax as indicated with the `example/DemoPreset.csv` file.

## How to install

Just download the executable from the `build` directory that matches your operating system. Currently, the following are
made:

* For windows
* For MacOSX
* For GNU/Linux

## How to build

### Build for your system

* First download a [Go distribution](https://go.dev/dl/) and install it on your machine by following the instructions in the web site.
* Then fetch the code of this project by clicking on the `Code` button and then `Download ZIP`
* Extract the code from the ZIP archive in a location of your choice
* Open a terminal and go to the directory in which is located the code, and then tape the following command to construct the program:

```
  go build
```

* that will generate the program `FujiSimuRecipesGen`. You can then launch it in the terminal like this:

```
  FujiSimuRecipesGen -s mySettings.yaml -csv mySpreadsheetsWithTheRecipes.csv
```

* the FP1 film simulations will be generated in the folder from which the program has been executed.

### Build for all supported systems

First ensure you have a Makefile distribution installed in your system.

Then, follow the premisses presented in the section above. Instead of executing the go program, just do:

```
  make
```

The limitation is the `Makefile` here is written for Unix system.
