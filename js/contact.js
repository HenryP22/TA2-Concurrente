$(document).ready(function(){

	$("form input[name='date']").datepicker();

	$.validate({
		lang: 'es',
		errorMessagePosition: "top",
		scrollToTopOnError: true
	});

});