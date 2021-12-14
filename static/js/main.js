const paginate = tableId => {
    $(`#${tableId}`).DataTable();
}

const alert = (title = 'Success', message = 'Your data has been saved', icon = 'success') => {
    Swal.fire({
        title,
        message,
        icon,
        allowOutsideClick: false
    })
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
    })
}