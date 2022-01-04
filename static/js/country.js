window.onload = () => {
    const url = '/countries/list';
    const tableHeaders = ['Country ID', 'Name', 'Capital'];
    const elementId = 'country_table'

    getData(url, tableHeaders, elementId);
}

const getData = (url, tableHeaders, elementId) => {
    let table = '<table id="table" class="table">';

    fetch(url)
        .then(response => response.json())
        .then(response => {
            table += '<thead class="table-dark">';
            table += '<tr>';

            tableHeaders.forEach(header => {
                table += `<th>${header}</th>`;
            });

            table += '</tr>';
            table += '</thead>';
            table += '</table>';
            document.getElementById(elementId).innerHTML = table;
        })
}