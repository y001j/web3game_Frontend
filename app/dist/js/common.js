var navItem=0;
jQuery(function ($) {

	jQuery(".menu-handler").click(function() {
	
				if(navItem == 0) {
					jQuery(this).addClass("active");
					jQuery("html").addClass("menuOpen");
					$(".mb").show();
					navItem = 1;
				} else {
					jQuery(this).removeClass("active");
					jQuery("html").removeClass("menuOpen");
					$(".mb").hide();
					navItem = 0;
				}
			});
	
	


});


	
  function navFixed() {
            var mainTop;
       
            if ($(window).scrollTop() >$('#pro').offset().top-100) {
                $("header").addClass("on");
            } else {
                $("header").removeClass("on");
            
            }

 }



// 搜索下拉
$('.searchBox .search-block').click(function(e){
    e.stopPropagation();
    $(".searchBox .box").stop().slideToggle();
});
$(".searchBox").click(function(e){
    e.stopPropagation();
})
$(".mb").click(function(e){
					$('.menu-handler').removeClass("active");
					jQuery("html").removeClass("menuOpen");
					$(".mb").hide();
					navItem = 0;
})



$('body').click(function(){
    $(".searchBox .box").stop().slideUp();
})

	

	

  $('.back').click(function(){
		$('html, body').animate({scrollTop: 0},300);
  });


	
window.onscroll = function () {

         
}


		






new WOW().init(0);   

				
		
				
				
      
				

$(".menu .item").hover(function(){
			
				$(this).find('.secNav').addClass("active").siblings().removeClass("active");
		  
				$(this).addClass("active").siblings().removeClass("active");
				
				},function(){
					  $(this).find('.secNav').removeClass("active");
					  $(this).removeClass("active");

				});	

		
			$(".menu li.item i").click(function(){
						
						if($(this).hasClass("active")){
							$(this).parent().parent().find(".secNav").stop(false, false).slideUp();
							$(this).removeClass("active");
							
						}else{
					
							$(this).parent().parent().find(".secNav").stop(false, false).slideDown();
							 $(this).addClass("active");
						}

			});
 
 

	


 //获取url中的参数
 function getUrlParam(name) {
            var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
            var r = window.location.search.substr(1).match(reg);  //匹配目标参数
            if (r != null) return unescape(r[2]); return null; //返回参数值
}
 
            


	
