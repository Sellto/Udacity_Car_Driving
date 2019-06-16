package main

// Server communicating with autonomous car simulator client
// task - Connect to the client (socketio)
// Task - Receiving camera image and speed value
// Task - Sending throttle and steer command

import (
	"log"
	 "github.com/urfave/cli"
	 "./action"
	 "os"
	 "errors"
	 "github.com/ewalker544/libsvm-go"
	 "strconv"
	 "fmt"
	)

func indexOf(element string, data []string) (int) {
	   for k, v := range data {
	       if element == v {
	           return k
	       }
	   }
	   return -1    //not found.
	}



func main() {
	var config,model,outputfile,inputfile,sample,testing,sidenum,nonenum string
	var svmtype,svmkerneltype,svmdegree,svmgamma,svmcoef,svmcost,svmnu,svmepsilonp,svmepsilon,svmcache,svmprob,svmverbose,svmcpu,svmif,svmof string
	app := cli.NewApp()
	app.Name = "SVM Project base on the Udacity Car Driving"
	app.Usage= ""
	app.Email = ""
	app.Author = "Emile Albert - Tom Selleslagh"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		 {
			 Name:    "run",
			 Aliases: []string{"r"},
			 Usage:   "",
			 Action:  func(c *cli.Context) error {
				 mode := indexOf(config,action.Available_config)
				 if mode > -1 {
					 action.RunServer(mode,model)
					 return nil
				 }
				 return errors.New("Choosen configuration not find")
			 },
			 Flags: []cli.Flag {
				 cli.StringFlag{
					Name: "config, c",
					Value: "hog32",
					Usage: "Configuration used to process the pictures from the car.\n\tThe availables configurations mode are : hog64,hog32,hogALL,gs,customgs,bw",
					Destination: &config,
				},
				cli.StringFlag{
				 Name: "model, m",
				 Value: "1",
				 Usage: "Path to the model which will be used by the autonomous car",
				 Destination: &model,
			 },
			},
		},
		{
			Name:    "extract",
			Aliases: []string{"e"},
			Usage:   "",
			Action:  func(c *cli.Context) error {
				mode := indexOf(config,action.Available_config)
				s,err := strconv.Atoi(sample)
				if err != nil {
					return errors.New("Bad sample pourcent format. Need integer value")
					}
				sn,err := strconv.Atoi(sidenum)
				if err != nil {
					return errors.New("Bad format for the quantity of side image. Need integer value")
					}
				nn,err := strconv.Atoi(nonenum)
				if err != nil {
					return errors.New("Bad format for the quantity of none image. Need integer value")
				}
				t,err := strconv.ParseFloat(testing,64)
				if err != nil {
					return errors.New("Bad testing pourcent format. Need float value")
				}
				switch mode {
				case 0 :
					action.GetSideHOGFeatures(action.HOG64,inputfile,outputfile,t,sn,nn)
					fmt.Println("bla")
				case 1 :
					action.GetSideHOGFeatures(action.HOG32,inputfile,outputfile,t,sn,nn)
				case 2 :
					action.GetHOGFeature(action.HOGALL,inputfile,s,t,outputfile)
				case 3:
					action.GetGSFeature(inputfile,s,t,outputfile)
				case 4:
					action.GetCustGSFeature(inputfile,s,t,outputfile)
				case 5:
					action.GetBWFeature(inputfile,s,t,outputfile)
				}

				return nil
			},
			Flags: []cli.Flag {
				cli.StringFlag{
				 Name: "config, c",
				 Value: "hogALL",
				 Usage: "Configuration used to process the pictures.\n\tThe availables configurations mode are : hog64,hog32,hogALL,gs,customgs,bw",
				 Destination: &config,
			 },
			 cli.StringFlag{
				Name: "outputFile, o",
				Value: "newDataSet",
				Usage: "Filename for the training and testing dataset",
				Destination: &outputfile,
			},
				cli.StringFlag{
				 Name: "inputFile, i",
				 Value: "driving_log.csv",
				 Usage: "csv file for the hogALL,gs,customgs,bw mode \n\t folder where the file is for the hog64,hog32 : \n\t\t- the class +1  is attribued to the side-[num].png image \n\t\t- the class -1 is attribued to the none-[num].png image",
				 Destination: &inputfile,
			 },
				 cli.StringFlag{
					Name: "sample, s",
					Value: "20",
					Usage: "Pourcentage from the all data which are used for the learning",
					Destination: &sample,
			},
				cli.StringFlag{
				 Name: "testing, p",
				 Value: "20.0",
				 Usage: "Pourcentage from the sample data that goes to the testing dataSet",
				 Destination: &testing,
			 },
				 cli.StringFlag{
					Name: "sidenum, n",
					Value: "1",
					Usage: "quantity of side images (hog64 and hog32 mode)",
					Destination: &sidenum,
			},
					cli.StringFlag{
					 Name: "nonenum, t",
					 Value: "1",
					 Usage: "quantity of none images (hog64 and hog32 mode)",
					 Destination: &nonenum,
				 },
		 },
	 },
		{
			Name:    "learn",
			Aliases: []string{"l"},
			Usage:   "",
			Action:  func(c *cli.Context) error {
				param := libSvm.NewParameter()
				convert, err := strconv.Atoi(svmtype)
				if err != nil {
					return errors.New("Bad svm Type")
				}
				param.SvmType = convert
				convert, err = strconv.Atoi(svmkerneltype)
				if err != nil {
					return errors.New("Bad svm Kernel Type")
				}
				param.KernelType = convert
				convert, err = strconv.Atoi(svmdegree)
				if err != nil {
					return errors.New("Bad svm Degree")
				}
				param.Degree = convert
				convert, err = strconv.Atoi(svmdegree)
				if err != nil {
					return errors.New("Bad svm Degree")
				}
				param.Degree = convert
				convert, err = strconv.Atoi(svmcache)
				if err != nil {
					return errors.New("Bad cache size value. Need integer value")
				}
				param.CacheSize = convert
				convert, err = strconv.Atoi(svmcpu)
				if err != nil {
					return errors.New("Bad cache size value. Need integer value")
				}
				param.NumCPU = convert
				convertf, err := strconv.ParseFloat(svmgamma,64)
				if err != nil {
					return errors.New("Bad svm Gamma")
				}
				param.Gamma = convertf
				convertf, err = strconv.ParseFloat(svmcoef,64)
				if err != nil {
					return errors.New("Bad svm Coef0")
				}
				param.Coef0 = convertf
				convertf, err = strconv.ParseFloat(svmepsilon,64)
				if err != nil {
					return errors.New("Bad svm Epsilon")
				}
				param.Eps = convertf
				convertf, err = strconv.ParseFloat(svmcost,64)
				if err != nil {
					return errors.New("Bad svm C")
				}
				param.C = convertf
				convertf, err = strconv.ParseFloat(svmnu,64)
				if err != nil {
					return errors.New("Bad svm Nu")
				}
				param.Nu = convertf
				convertf, err = strconv.ParseFloat(svmepsilonp,64)
				if err != nil {
					return errors.New("Bad svm Epsilon P value")
				}
				param.P = convertf
				convertb, err := strconv.ParseBool(svmprob)
				if err != nil {
					return errors.New("Bad svm Probability estimates value. Need true or false")
				}
				param.Probability = convertb
				convertb, err = strconv.ParseBool(svmverbose)
				if err != nil {
					return errors.New("Bad svm QuietMode value. Need true or false")
				}
				param.QuietMode = convertb
				action.Learning(svmif,svmof,param)
				return nil
			},
			Flags: []cli.Flag {
				cli.StringFlag{
				 Name: "svm_type, s",
				 Value: "0",
				 Usage: "set type of SVM :\n\t\t0 -- C-SVC\n\t\t1 -- nu-SVC\n\t\t2 -- one-class SVM\n\t\t3 -- epsilon-SVR\n\t\t4 -- nu-SVR\n\t\t",
				 Destination: &svmtype,
			 },
				 cli.StringFlag{
					Name: "kernel_type, k",
					Value: "2",
					Usage: "set type of kernel function : \n\t\t0 -- linear\n\t\t1 -- polynomial: (gamma*u'*v + coef0)^degree\n\t\t2 -- radial basis function: exp(-gamma*|u-v|^2)\n\t\t3 -- sigmoid: tanh(gamma*u'*v + coef0)\n\t\t",
					Destination: &svmkerneltype,
				},
					cli.StringFlag{
					 Name: "degree), d",
					 Value: "3",
					 Usage: "set degree in kernel function",
					 Destination: &svmdegree,
		 	 	},
				 cli.StringFlag{
					Name: "gamma : , g",
					Value: "0.0",
					Usage: "set gamma in kernel function",
					Destination: &svmgamma,
				},
				cli.StringFlag{
				 Name: "coef0 : , r",
				 Value: "0.0",
				 Usage: "set coef0 in kernel function",
				 Destination: &svmcoef,
				 },
				cli.StringFlag{
					Name: "cost : , c",
					Value: "1.0",
					Usage: "set the parameter C of C-SVC, epsilon-SVR, and nu-SVR",
					Destination: &svmcost,
				},
				cli.StringFlag{
				 Name: "nu : , n",
				 Value: "0.5",
				 Usage: "set the parameter nu of nu-SVC, one-class SVM, and nu-SVR",
				 Destination: &svmnu,
			 },
			 	cli.StringFlag{
					Name: "epsilon p: , p",
					Value: "0.1",
					Usage: "set the epsilon in loss function of epsilon-SVR",
					Destination: &svmepsilonp,
				},
				cli.StringFlag{
					Name: "cachesize : , m",
					Value: "100",
					Usage: "set cache memory size in MB",
					Destination: &svmcache,
				},
				cli.StringFlag{
					Name: "epsilon : , e",
					Value: "0.001",
					Usage: "set tolerance of termination criterion",
					Destination: &svmepsilon,
				},
				cli.StringFlag{
					Name: "probability_estimates : , b",
					Value: "false",
					Usage: "whether to train a SVC or SVR model for probability estimates, true or false ",
					Destination: &svmprob,
				},
				cli.StringFlag{
					Name: "QuietMode : , q",
					Value: "false",
					Usage: "Set verbose QuietMode on off, true or false",
					Destination: &svmverbose,
				},
				cli.StringFlag{
					Name: "cpu : , cp",
					Value: "-1",
					Usage: "Number of CPUs to use",
					Destination: &svmcpu,
				},
				cli.StringFlag{
					Name: "Dataset Filename : , i",
					Value: "",
					Usage: "Path to the dataset file",
					Destination: &svmif,
				},
				cli.StringFlag{
					Name: "Model Filename : , o",
					Value: "new.model",
					Usage: "Filename where the model will be dumped",
					Destination: &svmof,
				},
		 },
	 },
	 {
		 Name:    "test",
		 Aliases: []string{"t"},
		 Usage:   "",
		 Action:  func(c *cli.Context) error {
			 if config == "mean" {
				 mean,sd := action.GetMeanSD(inputfile,model)
				 fmt.Println("Mean :",mean,"- SD :",sd)
			 } else if config == "pourcent" {
				 pourcent := action.GetPourcent(inputfile,model)
				 fmt.Println("Efficiency : ",pourcent)
			 } else {
				 return errors.New("Unknow choosen mode")
			 }
			 return nil
		 },
		 Flags: []cli.Flag {
			 cli.StringFlag{
				Name: "mode, o",
				Value: "",
				Usage: "output mode (mean or pourcent) ",
				Destination: &config,
			},
			cli.StringFlag{
			 Name: "testingfile, t",
			 Value: "",
			 Usage: "Path to the dataset file",
			 Destination: &inputfile,
		 },
			cli.StringFlag{
			 Name: "model, m",
			 Value: "1",
			 Usage: "Path to the model which will be tested",
			 Destination: &model,
		 },
		},
	},
	 }
 err := app.Run(os.Args)
 if err != nil {
	 log.Fatal(err)
 }
}
