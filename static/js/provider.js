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

const create = () => {
    const name = document.getElementById('inp_provider_name').value;
    const address = document.getElementById('inp_address').value;
    const legarRepresentative = document.getElementById('inp_legal_representative').value;
    const cellPhone = document.getElementById('inp_cell_phone').value;
    const country = document.getElementById('slcCountry').value;
    const ruc = document.getElementById('inp_ruc').value;

    if (name.trim().length === 0) {
        alert('Name field is required', '', 'error');
        return;
    }

    if (address.trim().length === 0) {
        alert('Address field is required', '', 'error');
        return;
    }

    if (legarRepresentative.trim().length === 0) {
        alert('Legal representative field is required', '', 'error');
        return;
    }

    if (cellPhone.trim().length === 0) {
        alert('Cell phone field is required', '', 'error');
        return;
    }

    if (country.trim().length === 0) {
        alert('Country field is required', '', 'error');
        return;
    }

    if (ruc.trim().length === 0) {
        alert('RUC field is required', '', 'error');
        return;
    }

    console.log('create provider success');
    // confirmation().then((result) => {
    //     if (result.isConfirmed) {
    //         document.getElementById('frmCreateCategory').submit();
    //     }
    // })
}