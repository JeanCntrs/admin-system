window.onload = () => {
    const url = '/providers/list';
    const tableHeaders = ['Provider ID', 'Name', 'Phone', 'Country Name'];
    const fields = ['ProviderId', 'Name', 'Phone', 'CountryName'];
    const elementId = 'provider_table';

    getData(url, tableHeaders, fields, elementId);
}