window.onload = () => {
    const url = '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const elementId = 'country_table'
    const fields = ['CountryId', 'Name', 'Capital'];

    getData(url, tableHeaders, fields, elementId);
}