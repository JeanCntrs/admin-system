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

const getEntityById = providerId => {
    fetch(`/providers/${providerId}`)
        .then(response => response.json())
        .then(response => {
            document.getElementById('inp_provider_id').value = response.ProviderId;
            document.getElementById('inp_provider_name').value = response.Name;
            document.getElementById('inp_address').value = response.Address;
            document.getElementById('inp_phone').value = response.Phone;
            document.getElementById('inp_email').value = response.Email;
            document.getElementById('inp_legal_representative').value = response.LegalRepresentative;
            document.getElementById('inp_cell_phone').value = response.CellPhone;
            document.getElementById('slcCountry').value = response.CountryId;
            document.getElementById('inp_ruc').value = response.Ruc;
        })
}