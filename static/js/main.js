const paginate = tableId => {
    $(`#${tableId}`).DataTable();
}

const alert = (title = 'Success', text = 'Your data has been saved', icon = 'success') => {
    Swal.fire({
        title,
        text,
        icon,
        allowOutsideClick: false
    });
}

const confirmation = (title = 'Are you sure?', text = 'If you are sure confirm the action') => {
    return Swal.fire({
        title,
        text,
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes, confirm it!',
        allowOutsideClick: false
    });
}

const getData = (url, tableHeaders, fields, elementId) => {
    let table = '<table id="table" class="table">';

    fetch(url)
        .then(response => response.json())
        .then(response => {
            table += '<thead class="table-dark">';
            table += '<tr>';

            tableHeaders.forEach(header => {
                table += `<th>${header}</th>`;
            });

            table += '</tr>';
            table += '</thead>';
            table += '<tbody>';

            response.forEach(element => {
                table += '<tr>';

                fields.forEach(field => {
                    table += '<td>';
                    table += element[field]
                    table += '</td>';
                });

                table += '</tr>';
            });

            table += '</tbody>';
            table += '</table>';
            document.getElementById(elementId).innerHTML = table;
        })
}