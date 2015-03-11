package main

var dataPath string = `/var/lib/noteddata/`
var helpStr string = `notice:
	noted add "exampleKey" "exampleContent" : add a note;
	noted del "exampleKey" :del a note;
	noted get "exampleKey" :get a note;
	noted append "exampleKey" :append;
	
	Warning : when first using this cmd, u should mkdir "/var/lib/noteddata/" and change permission.`
