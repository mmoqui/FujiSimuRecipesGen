# FujiSimuRecipesGen
A Go program to generate Fujifilm simulations from recipes of such simulations in CSV

A lot of persons write their recipes of Fujifilm film simulation within a spreadsheet kept up-to-date. Usually those come from 
[Fuji X Weekly](https://fujixweekly.com/). Then, they set themselves the properties of the recipe either directly into their own Fujifilm camera or
via the Fujifilm X-Raw Studio tool to create a custom film simulation in the camera.

The `FujiSimuRecipesGen` program aims to automatically generate Fujifilm simulations in FP1 format from a such spreadsheet. These FP1 files then can be 
copied or moved into the location of the film simulations for Fujifilm X-Raw Studio so that you can transfer them directly to your camera with the Fujifilm tool.

## Requirements

* Create a settings file in YAML in which you specify the properties of both your camera and the version of Fujifilm X-Raw Studio (later is used to indicate as 
the film simulation was created by this tool). To know the required properties, please look at `example/settings.yaml` file.
* Have your spreadsheet in CSV format following the syntax as indicated with the `example/DemoPreset.csv` file.

## How to build

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
