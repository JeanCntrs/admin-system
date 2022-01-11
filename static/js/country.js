window.onload = () => {
    const url = '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const fields = ['CountryId', 'Name', 'Capital'];
    const elementId = 'country_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'CountryId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}

const getCountriesByName = () => {
    const name = document.getElementById("inp_search_category_name").value;

    const url = name ? `/countries/list/name/${name}` : '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const fields = ['CountryId', 'Name', 'Capital'];
    const elementId = 'country_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'CountryId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}

const getCountryById = countryId => {
    fetch(`/countries/list/id/${countryId}`)
        .then(response => response.json())
        .then(response => {
            document.getElementById('inp_country_id').value = response.CountryId;
            document.getElementById('inp_country_name').value = response.Name;
            document.getElementById('inp_country_capital').value = response.Capital;
        })
}