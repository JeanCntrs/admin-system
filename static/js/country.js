window.onload = () => {
    createMenu();
    listCountries();
}

const listCountries = () => {
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

const getEntityById = countryId => {
    fetch(`/countries/list/id/${countryId}`)
        .then(response => response.json())
        .then(response => {
            document.getElementById('inp_country_id').value = response.CountryId;
            document.getElementById('inp_country_name').value = response.Name;
            document.getElementById('inp_country_capital').value = response.Capital;
        })
}

const create = () => {
    const countryId = document.getElementById('inp_country_id').value;
    const countryName = document.getElementById('inp_country_name').value;
    const countryCapital = document.getElementById('inp_country_capital').value;

    if (countryName.trim().length === 0) {
        alert('Name field is required', '', 'error');
        return;
    }

    if (countryCapital.trim().length === 0) {
        alert('Capital field is required', '', 'error');
        return;
    }

    const country = {
        countryId: countryId == '' ? 0 : parseInt(countryId),
        name: countryName,
        capital: countryCapital
    }

    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch('countries/create', {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST',
                body: JSON.stringify(country)
            })
                .then(response => response.text())
                .then(response => {
                    if (response != '1') {
                        alert('An error has occurred');

                        return;
                    }

                    document.getElementById('btnCloseModal').click();
                    listCountries();
                    alert();

                    return;
                })
        }
    })
}

const deleteEntity = (id) => {
    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch(`countries/delete/${id}`)
                .then(response => response.text())
                .then(response => {
                    if (response != '1') {
                        alert('An error has occurred', '');

                        return;
                    }

                    listCountries();
                    alert('Success', 'Your data has been deleted');

                    return;
                })
        }
    })
}