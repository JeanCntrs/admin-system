window.onload = () => {
    buildTable();
}

const buildTable = () => {
    const url = '/pages/list';
    const tableHeaders = ['Page ID', 'Message', 'Route'];
    const fields = ['PageId', 'Message', 'Route'];
    const elementId = 'create_role_page_table';
    const showBtnEdit = false;
    const showBtnDelete = false;
    const propertyName = 'PageId';
    const isPopup = false;
    const isChecked = true;

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName, undefined, isPopup, isChecked);
}