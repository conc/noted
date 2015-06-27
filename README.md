noted
==== 
#How to use:

noted add "exampleKey" "exampleContent" : add a note;  
noted del "exampleKey" :del a note;  
noted get "exampleKey" :get note;  
noted append "exampleKey" :append;  
	
Warning : when first using this cmd, u should mkdir "/var/lib/noteddata/" and change permission.  

#What is this

一个随笔记录程序，linux下命令行使用。

#Examples
noted add onename "zhang3"；  
noted get onename；  
noted append onename "333"；  
noted del onename；  

