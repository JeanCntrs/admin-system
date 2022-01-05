window.onload = () => {
    const url = '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const fields = ['CountryId', 'Name', 'Capital'];
    const elementId = 'country_table';

    getData(url, tableHeaders, fields, elementId);
}

const getCountriesByName = () => {
    const name = document.getElementById("inp_search_category_name").value;

    const url = name ? `/countries/list/${name}` : '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const fields = ['CountryId', 'Name', 'Capital'];
    const elementId = 'country_table';

    getData(url, tableHeaders, fields, elementId);
}