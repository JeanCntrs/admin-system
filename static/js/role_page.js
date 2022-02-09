window.onload = () => {
    buildTable();
}

const buildTable = () => {
    const url = '/role-types/list';
    const tableHeaders = ['Role Type ID', 'Name', 'Description'];
    const fields = ['RoleTypeId', 'Name', 'Description'];
    const elementId = 'role_page_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'RoleTypeId';
    const isPopup = false;
    const isChecked = false;

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName, undefined, isPopup, isChecked);
}

const getEntityById = roleTypeId => {
    window.location.href = `/role-page/edit/${roleTypeId}`;
}