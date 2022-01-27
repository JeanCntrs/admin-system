window.onload = () => {
    buildTable();
    buildSelectRoleType();
    buildSelectPerson();
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

const buildSelectRoleType = () => {
    const url = '/role-types/list';
    const value = 'RoleTypeId';
    const name = 'Name';
    const elementId = 'slcRoleType';

    getDataSelect(url, value, name, elementId);
}

const buildSelectPerson = () => {
    const url = '/persons/list';
    const value = 'PersonId';
    const name = 'Fullname';
    const elementId = 'slcPerson';

    getDataSelect(url, value, name, elementId);
}