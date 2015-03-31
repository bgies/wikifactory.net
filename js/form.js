var email;
// post email to datastore
$('.submit').click(function(e){
    $.post( "/signup", {email:$('.access').val()}, function( data ) {
        $('.thankyou').addClass('visible');
    });
    e.preventDefault();
    e.stopPropagation();
});
// post contact form values to data store
$('.tiny-submit').click(function(e){
    $.post( "/contact", {name:$('.contactname').val(), email:$('.contactemail').val(), subject:$('.contactsubject').val(), message:$('.contactmessage').val(), emailto:email}, function( data ) {
        //set contact form to invisible
        $('.contactform').removeClass('visible');
        // empty the input fields.
        $('.contactname').val("");
        $('.contactemail').val("");
        $('.contactsubject').val("");
        $('.contactmessage').val("");
    });
    e.preventDefault();
    e.stopPropagation();
});
// open contact form when clicking on these buttons
$('.large-btn, .small-btn, .contact, .new').click(function(e){
    email = $(e.target).parent().attr("id");
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').addClass('visible');
});
// close contact form when clicking on these divs
$('.projects-wrapper, .network-nodes, .cancel').click(function(e){
    e.preventDefault();
    e.stopPropagation();
    $('.contactform').removeClass('visible');
});
