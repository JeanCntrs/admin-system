window.onload = () => {
    buildTable();
    buildSelectSearch();
    buildSelect();
}

const buildTable = () => {
    const url = '/providers/list';
    const tableHeaders = ['Provider ID', 'Name', 'Phone', 'Country Name'];
    const fields = ['ProviderId', 'Name', 'Phone', 'CountryName'];
    const elementId = 'provider_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'ProviderId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}

const buildSelectSearch = () => {
    const url = '/countries/list';
    const value = 'CountryId';
    const name = 'Name';
    const elementId = 'slc_search_country';

    getDataSelect(url, value, name, elementId);
}

const buildSelect = () => {
    const url = '/countries/list';
    const value = 'CountryId';
    const name = 'Name';
    const elementId = 'slcCountry';

    getDataSelect(url, value, name, elementId);
}

const getProviderByCountryId = () => {
    const countryId = document.getElementById('slc_search_country').value;

    const url = countryId ? `/providers/list/id/${countryId}` : '/providers/list';
    const tableHeaders = ['Provider ID', 'Name', 'Phone', 'Country Name'];
    const fields = ['ProviderId', 'Name', 'Phone', 'CountryName'];
    const elementId = 'provider_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'ProviderId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}