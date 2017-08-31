/*
	confirm prompt example
    <a data-message="Are you sure you would like todo something?" data-color="#ff0000" data-url="/endpoint" data-type="warning">Delete</a>
*/

/*
	conrirm action example
	confirm.yes = function(btn) {
        $('<form method="post" action="' + btn.attr('data-url') + '"></form>').submit();
    }
*/

function Confirm() {
	this.registerDisplay();
}
Confirm.prototype = {
	color: '',
	yes: function(btn) {
		alert('success');
	},
	registerDisplay: function() {

		$('.confirm-action').click(function(e) {
            e.stopPropagation();
			var btn = $(this);
            swal({
                title: '',
                text: btn.attr('data-message'),
                type: btn.attr('data-type'),
                showCancelButton: true,
                confirmButtonColor: btn.attr('data-color'),
                confirmButtonText: "Yes",
                closeOnConfirm: false
            }, function(){
				confirm.yes(btn);
            });
		});
	}
};

var confirm = new Confirm();
