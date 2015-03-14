// Dynamic url update when scrolling into anchor tags
$(function () {
    var currentHash = ""
    $(document).scroll(function () {
        $('.tag').each(function () {
            var top = window.pageYOffset;
            var distance = top - $(this).offset().top;
            var hash = $(this).attr('id');
            // 30 is an arbitrary padding choice,
            // if you want a precise check then use distance===0
            if (distance < 80 && distance > -10 && currentHash != hash) {
                window.location.hash = (hash);
                currentHash = hash;
            }
        });
    });
});
