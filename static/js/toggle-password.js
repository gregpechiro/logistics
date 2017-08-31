$('button.toggle-pass').click(function() {
	var btn = $(this);
	var input = btn.closest('form').find('input.toggle-pass');
	var i = btn.find('i');
	if (input.attr('type') === 'password') {
		i.removeClass("fa-eye-slash");
		i.addClass("fa-eye");
		input.attr('type', 'text');
	} else {
		i.removeClass("fa-eye");
		i.addClass("fa-eye-slash");
		input.attr('type', 'password');
	}
});
