# Artificial Intelligence project <br/> ***Self driving car***
@ ECAM Brussels | June 2019

The objective of this scholar project is to implement an algorithm solution for self car driving. The environment is
simulated with the Udacity simulator ["self-driving-car-sim"](https://github.com/udacity/self-driving-car-sim).

## Library files
* image.go<br/>
Library developped for the project<br/>
**[Utility]** This file contains all required functions used for image processing (cropping, grayscale, resizing...)

* hog.go<br/>
Library developped for the project and based on [go-HoG](https://github.com/satojkovic/go-HoG-sample)<br/>
**[Utility]**

* parse.go<br/>
Library developped for the project<br/>
**[Utility]** This script defines the functions required for parsing CSV file returned
by the manual training phase in Udacity simulator and for parsing txt file obtained after the extraction and the testing dataset generation.

* svm.go<br/>
Library developped for the project and based on [libsvm-go](https://www.csie.ntu.edu.tw/~cjlin/libsvm/)<br/>
**[Utility]** This script contains all functions related to the use of SVM
classifier (predict, calculate prediction error, generate test entry)

* scan.go<br/>
Library developed for the project<br/>
**[Utility]** File defining the function required for line detection. Practically speaking the main objective is scanning the received image and predict line position
function of the computation of Histogram of Oriented Gradients (HoG).

* displayhog.go<br/>
Library developed for the project<br/>
**[Utility]** The file defines function for fancy displaying HoG computation results on the image.   



## Structure description
### *action* directory
The action directory contains different tools for different uses as explained below :

***drive.go*** contains all functions required for Udacity simulator (socket server) communication. It sends prediction
for steering wheel angle function of received images from centered camera and send a value for throttle (limited to
 15 mph). The processing of received images is initiated from this file and function of the parameters described in
***param.go***. These functions are consequently Udacity simulator specific.    

***extract.go*** extracts features from the raw dataset obtained after manual training laps in the Udacity simulator.
The extracted features are then organized in two dataset : training dataset and testing dataset. These functions are
consequently Udacity simulator specific.  

***learning.go*** generates the SVM model for prediction. The generated model is based on the given training dataset and
all the parameters for SVM classification.

***testing.go*** allows to test prediction model generated with ***learning.go*** on the testing dataset (see
***extract.go***) to compare prediction result and target value.


### *lib* directory   
The lib directory contains all files composing the project library. All files description as well as their use is
described in the "Library section" above.

### *main.go*
This last file is the main CLI application allowing to use all tools containing in action package (see *action*
directory section). As all tools parameters can be given in the launch command line, the main.go file parses in
the first instance all given inputs from the user to properly execute tools and function as requested.

 ## Usage with MAC/Linux 
A compiled CLI application, called svmtool, is downloadable on the release section. It offers 4 tools: extract, run, learn and test. For each one, a *help* section is available with -h parameter :
```bash
./svmtool run -h
```

Here are some use cases but don't hesitate to abuse *help* section to discover all available features:

### Launch Socket IO server
To launch the Socket IO Server. You have to use the **run** tool with two paramaters: the configuration and the model, respectively used for processing incoming pictures and SVM prediction. Here is an example :

```bash
./svmtool run -c hog64 -m HOG_C8x8_BS2_Or9_Grayscale_RightSide_CSVM_POLY3.model
```
### Extract Data
You can use the **extract** tool to create the  training and testing datasets from the Udacity simulator *csv* file.

```bash
./svmtool extract -o CGS_Dataset -c customgs -s 100 -t 10.0
```

### SVM tool

The **learn** and **test** tool are completely independent of the Udacity simulator. You can use it for create/learn and test any SVM model.
See help section for all different possibilities.

**Remark : use svmtool.exe for Windows usage**

## License
[MIT](https://choosealicense.com/licenses/mit/)
