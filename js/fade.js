// Fade arrow away when scrolling down.
$(window).scroll(function() {
    $(".arrow").css({
    'opacity' : 1-(($(this).scrollTop())/225)
    });
});
