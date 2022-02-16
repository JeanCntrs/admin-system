const socket = new WebSocket('ws://localhost:8000/socket');

socket.onopen = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Connected';
}

socket.onclose = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Disconnected';
}

socket.onmessage = (event) => {
    const data = event.data;

    if (data == 'createProvider') {
        const tableId = 'table';
        const currentPageIndex = getCurrentPageIndex(tableId);

        buildTable(() => {
            getCurrentPage(tableId, currentPageIndex);
        });
    } else if (data == 'createCountry') {
        buildSelectSearch();
        buildSelect();
    }
}

window.onload = () => {
    createMenu();
    buildTable(() => { });
    buildSelectSearch();
    buildSelect();
}

const buildTable = (callback) => {
    const url = '/providers/list';
    const tableHeaders = ['Provider ID', 'Name', 'Phone', 'Country Name'];
    const fields = ['ProviderId', 'Name', 'Phone', 'CountryName'];
    const elementId = 'provider_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'ProviderId';
    const isPopup = true;
    const isChecked = false;
    const isCallback = true;

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName, undefined, isPopup, isChecked, isCallback, () => {
        callback();
    });
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
    const providerId = document.getElementById('inp_provider_id').value;
    const name = document.getElementById('inp_provider_name').value;
    const address = document.getElementById('inp_address').value;
    const phone = document.getElementById('inp_phone').value;
    const email = document.getElementById('inp_email').value;
    const legalRepresentative = document.getElementById('inp_legal_representative').value;
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

    if (legalRepresentative.trim().length === 0) {
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

    const provider = {
        providerId: providerId == '' ? 0 : parseInt(providerId),
        name,
        address,
        phone,
        email,
        legalRepresentative,
        cellPhone,
        countryId: parseInt(country),
        ruc
    }

    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch('providers/create', {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST',
                body: JSON.stringify(provider)
            })
                .then(response => response.text())
                .then(response => {
                    if (response != 1) {
                        alert('An error has occurred');

                        return;
                    }

                    socket.send('createProvider');

                    document.getElementById('btnCloseModal').click();
                    buildTable();
                    alert();

                    return;
                })
        }
    })
}

const deleteEntity = (id) => {
    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch(`providers/delete/${id}`)
                .then(response => response.text())
                .then(response => {
                    if (response != '1') {
                        alert('An error has occurred', '');

                        return;
                    }

                    buildTable();
                    alert('Success', 'Your data has been deleted');

                    return;
                })
        }
    })
}