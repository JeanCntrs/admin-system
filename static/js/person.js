window.onload = () => {
    $("#inp_birthday").datepicker();
    buildTable();
}

const buildTable = () => {
    const url = '/persons/list';
    const tableHeaders = ['Person ID', 'Fullname', 'Name Type Person', 'Birthday'];
    const fields = ['PersonId', 'Fullname', 'NameTypePerson', 'FormattedBirthday'];
    const elementId = 'person_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'PersonId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}

const getPersonsByFullname = () => {
    const fullname = document.getElementById("inp_search_person_fullname").value;

    const url = fullname ? `/persons/list/name/${fullname}` : '/persons/list';
    const tableHeaders = ['Person ID', 'Fullname', 'Name Type Person', 'Birthday'];
    const fields = ['PersonId', 'Fullname', 'NameTypePerson', 'FormattedBirthday'];
    const elementId = 'person_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'PersonId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}