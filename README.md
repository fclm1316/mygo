checktargetmd5
	
	背景：需要知道远程机器文件(jar文件)是否更新过。
	方案：1 通过ssh远程库连接远程机器算文件md5，文件创建时间大小，
	     保存在数据库中，并记录时间批次号
		 2 通过批次号对文件进行比对。注意文件的增加删除
	     3 使用rustfull 风格，get post delete put对接口增删改查。
		   
	接口：sysinfo 系统信息相关接口，ip地址所属环境，名称
	     sysenv  系统环境信息接口，验证测试压测生产全链路
		 sysname 系统名称接口
		 filelist 远程机器文件路径接口，可精准文件名也可模糊匹配
		 fileinfo 远程文件信息接口
		 sysuserpasswd 远程机器的用户名密码接口
		 sysbatch  批次接口
		 fileresult 比对结果接口
		
	
	
	
getnetworksource
    
    在chaosblade混沌工程中,需要对架构感知,系统调用之间的关系
    类似需要skywalking的调用关系图.利用gopacket(libpcap)抓取
    网卡流量并分析.....
    
    
    rebuildtcp 服务端
    		接受perception_http请求数据
    
    
    getpacket 
    		抓取服务端的包
    perception_http  客户端
    		从网卡中获得tcp包，组装后发给服务端，参考gor，增加uuid提高重组率
			需要改成发送kafka中
    	    

英语背单词

    english orm、洗牌算法
        欢迎来到背单词系统
        1 单元单词
        2 章节单词
        3 遗忘单词
        4 全部单词
        5 退出系统
        请输入需要的序列号: 4
        
        n:下一个, p:跳过, q:返回
        mankind   

    
小工具

    myojb
        centipede  模仿py脚本，类ansible
        pgsql   压力工具
        scanport 扫描21端口工具
        tidb 连接数据tidb
        tools 工具
        yq  爬虫提交表单
第二本书

    sbook    
第一本书

    fbook
    
其他书

    my