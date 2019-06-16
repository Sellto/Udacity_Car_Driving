# Artificial Intelligence project <br/> ***Self driving car***
@ ECAM Brussels | June 2019 

The objective of this scholar project is to implement an algorithm solution for self car driving. The environment is 
simulated with the Udacity simulator ["self-driving-car-sim"](https://github.com/udacity/self-driving-car-sim).

## Libraries 

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
## Usage

```bash

```

## License
[MIT](https://choosealicense.com/licenses/mit/)




