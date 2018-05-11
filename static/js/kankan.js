var kkDapCtrl = null;
function kkGetDapCtrl()
{
	if(null == kkDapCtrl) {
	  try{
	  	if (window.ActiveXObject) {
	  	//if (navigator.userAgent.indexOf('MSIE') != -1) {
				kkDapCtrl = new ActiveXObject("DapCtrl.DapCtrl");  		
	  	}	else {
				var browserPlugins = navigator.plugins;
				for (var bpi=0; bpi<browserPlugins.length; bpi++) {
					try {
						if (browserPlugins[bpi].name.indexOf('Thunder DapCtrl') != -1) {
							var e = document.createElement("object");   
							e.id = "dapctrl_history";   
							e.type = "application/x-thunder-dapctrl"; 
							e.width = 0;   
							e.height = 0;
							document.body.appendChild(e);
							break;
						}
					} catch (e) {}
				}
				kkDapCtrl = document.getElementById('dapctrl_history');
	  	}
	  } catch(e) {}
	}
	return kkDapCtrl;
}

function start(url){
  var dapCtrl=kkGetDapCtrl();  
  try {
		var ver = dapCtrl.GetThunderVer("KANKAN", "INSTALL");
		var type = dapCtrl.Get("IXMPPACKAGETYPE");
		if(ver && type && ver >= 672 && type >= 2401)
		{
			dapCtrl.Put("sXmp4Arg", '"'+url+'"'+' /sstartfrom _web_xunbo /sopenfrom web_xunbo');			
		}	else {
			//alert('请先更新迅雷看看播放器,然后刷新本页面！');
				if(window.confirm("请先更新迅雷看看播放器\n\n点击“确定”下载并安装更新看看播放器\n\n安装后请刷新本页面\n\n如果不需要下载请点击“取消”")){
	window.open("http://xmp.down.sandai.net/kankan/XMPSetup_4.9.14.2052-www.exe");
	}
		}
	} catch(e) {
  	//alert('请先安装迅雷看看播放器,然后刷新本页面！');
	if(window.confirm("您未安装迅雷看看播放器\n\n点击“确定”下载并安装看看播放器\n\n安装后请刷新本页面\n\n如果不需要下载请点击“取消”")){
	window.open("http://xmp.down.sandai.net/kankan/XMPSetup_4.9.14.2052-www.exe");
	}
	}
}