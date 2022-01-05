window.onload = () => {
    const url = '/providers/list';
    const tableHeaders = ['Provider ID', 'Name', 'Phone', 'Country Name'];
    const fields = ['ProviderId', 'Name', 'Phone', 'CountryName'];
    const elementId = 'provider_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'ProviderId';

    getData(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}