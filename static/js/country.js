window.onload = () => {
    const url = '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const fields = ['CountryId', 'Name', 'Capital'];
    const elementId = 'country_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'CountryId';

    getData(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}

const getCountriesByName = () => {
    const name = document.getElementById("inp_search_category_name").value;

    const url = name ? `/countries/list/${name}` : '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const fields = ['CountryId', 'Name', 'Capital'];
    const elementId = 'country_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'CountryId';

    getData(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}