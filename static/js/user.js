window.onload = () => {
    buildTable();
}

const buildTable = () => {
    const url = '/users/list';
    const tableHeaders = ['User ID', 'Username', 'Fullname', 'Role Type'];
    const fields = ['UserId', 'Username', 'Fullname', 'RoleTypeName'];
    const elementId = 'user_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'UserId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}