window.onload = () => {
    buildTable();
}

const buildTable = () => {
    const url = '/pages/list';
    const tableHeaders = ['Page ID', 'Message', 'Route'];
    const fields = ['PageId', 'Message', 'Route'];
    const elementId = 'page_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'PageId';
    const isPopup = false;

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName, undefined, isPopup);
}

const getEntityById = pageId => {
    fetch(`/pages/list/id/${pageId}`)
        .then(response => response.json())
        .then(response => {
            document.getElementById('inp_page_id').value = response.PageId;
            document.getElementById('inp_message').value = response.Message;
            document.getElementById('inp_route').value = response.Route;
        })
}

const create = () => {
    const pageId = document.getElementById('inp_page_id').value;
    const message = document.getElementById('inp_message').value;
    const route = document.getElementById('inp_route').value;

    if (message.trim().length === 0) {
        alert('Message field is required', '', 'error');
        return;
    }

    if (route.trim().length === 0) {
        alert('Route field is required', '', 'error');
        return;
    }

    const page = {
        pageId: pageId == '' ? 0 : parseInt(pageId),
        message,
        route
    }
    
    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch('pages/create', {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST',
                body: JSON.stringify(page)
            })
                .then(response => response.text())
                .then(response => {
                    if (response != '1') {
                        alert('An error has occurred');

                        return;
                    }

                    buildTable();
                    clearInputs();
                    alert();

                    return;
                })
        }
    })
}