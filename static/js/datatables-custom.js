
var TRCLICK = function(tr) {
    if (tr.getAttribute('data-target') === '_blank') {
        window.open(tr.getAttribute('data-url'));
        return
    }
    window.location.href = tr.getAttribute('data-url');
}

$(document).on('click', 'tr.clickable', function() {
    TRCLICK(this);
});

$('#filter').on( 'keyup', function () {
    if (table !== undefined && !$.isEmptyObject(table)) {
        table.search( this.value ).draw();
    }
});

$('tr#search th').each(function () {
    var title = $(this).text();
    if (title != '') {
        $(this).html('<input id="columnSearch" data-column="' + title + '" type="text"/>');
    }
});

$('input#columnSearch').keyup(function() {
    var name = $(this).attr('data-column');
    var column = table.column(name + ':name');
    if (column.search() !== this.value ) {
        column.search(this.value).draw();
    }
});
