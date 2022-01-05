window.onload = () => {
    const url = '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const fields = ['CountryId', 'Name', 'Capital'];
    const elementId = 'country_table';

    getData(url, tableHeaders, fields, elementId);
}