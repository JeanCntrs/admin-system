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