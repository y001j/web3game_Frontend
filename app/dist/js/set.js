    function dingwei(){
        top1=$('.i1').position().top+10+"px";
        $('.wz1').css({"top":top1,"left":0});
        top2=$('.i2').position().top+10+"px";
        $('.wz2').css({"top":top2,"left":0});
        top3=$('.i3').position().top+10+"px";
        $('.wz3').css({"top":top3,"left":0});


        top4=$('.i4').position().top+10+"px";
        $('.wz4').css({"top":top4,"right":0});
        top5=$('.i5').position().top+10+"px";
        $('.wz5').css({"top":top5,"right":0});
        top6=$('.i6').position().top+10+"px";
        $('.wz6').css({"top":top6,"right":0});

    }



	function setfont(){
		var winW = window.innerWidth;
		if (winW <= 1600 && winW > 800) {
			size = Math.round(winW / 16);
			$('html').css({'font-size': size/1.2 + 'px'})
		} else if (winW <= 750) {
				document.getElementsByTagName('html')[0].style.fontSize = document.documentElement.clientWidth/7.5 + 'px'
		} 

	}
	
    window.onresize = function() {
        dingwei();

		setfont();
   
    }
	

setfont();

