var values = [];

$('.multi').click(function(e) {
    e.stopPropagation();
    var $this = $(this);
    var v = $this.attr('data-value');
    var i = values.indexOf(v)
    var $icn = $this.find('i')
    if (i < 0) {
        values.push(v);
        $this.addClass('btn-success');
        $this.removeClass('btn-default');
        $icn.addClass('fa-check');
        $icn.removeClass('fa-plus');
        if (values.length > 0) {
            $('.multi-submit').removeClass('disabled');
        }
        return
    }
    values.splice(i, 1);
    $this.addClass('btn-default');
    $this.removeClass('btn-success');
    $icn.removeClass('fa-check');
    $icn.addClass('fa-plus');
    if (values.length < 1) {
        $('.multi-submit').addClass('disabled');
    }
});

$('.multi-submit').click(function() {
    $('form.multi-form input.multi-input').val(values.join(','));
    $('form.multi-form').submit();
    // console.log($('form.multi-form input.multi-input').val());
});
