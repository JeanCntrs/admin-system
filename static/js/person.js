window.onload = () => {
    buildTable();
}

const buildTable = () => {
    const url = '/persons/list';
    const tableHeaders = ['Person ID', 'Fullname', 'Name Type Person', 'Birthday'];
    const fields = ['PersonId', 'Fullname', 'NameTypePerson', 'Birthday'];
    const elementId = 'person_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'PersonId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}