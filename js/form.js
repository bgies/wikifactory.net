// get early access submit
$('.submit').click(function(e){
    $.post( "/signup", {email:$('.access').val()}, function( data ) {
        $('.thankyou').addClass('visible');
    });
    e.preventDefault();
    e.stopPropagation();
});
// contact form submit
$('.tiny-submit').click(function(e){
    $.post( "/contact", {name:$('.contactname').val(), email:$('.contactemail').val(), subject:$('.contactsubject').val(), message:$('.contactmessage').val()}, function( data ) {
        $('.contactform').removeClass('visible');
    });
    e.preventDefault();
    e.stopPropagation();
});
// copen general contact form when clicking on the following buttons
$('.large-btn').click(function(e){
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').addClass('visible');
});
$('.small-btn').click(function(e){
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').addClass('visible');
});
$('.contact').click(function(e){
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').addClass('visible');
});
$('.new').click(function(e){
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').addClass('visible');
});
// close general contact form when clicking on wrappers or cancel button
$('.projects-wrapper').click(function(e){
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').removeClass('visible');
});
$('.network-nodes').click(function(e){
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').removeClass('visible');
});
$('.cancel').click(function(e){
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').removeClass('visible');
});
