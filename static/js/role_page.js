const socket = new WebSocket('ws://localhost:8000/socket');

socket.onopen = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Connected';
}

socket.onclose = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Disconnected';
}

socket.onmessage = (event) => {
    const data = event.data;

    if (data == 'createRoleType' || data == 'editRoleType') {
        const tableId = 'table';
        const currentPageIndex = getCurrentPageIndex(tableId);

        buildTable(() => {
            getCurrentPage(tableId, currentPageIndex);
        });
    }
}

window.onload = () => {
    createMenu();

    buildTable(() => { });
}

const buildTable = (callback) => {
    const url = '/role-types/list';
    const tableHeaders = ['Role Type ID', 'Name', 'Description'];
    const fields = ['RoleTypeId', 'Name', 'Description'];
    const elementId = 'role_page_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'RoleTypeId';
    const isPopup = false;
    const isChecked = false;
    const isCallback = true;

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName, undefined, isPopup, isChecked, isCallback, () => {
        callback();
    });
}

const getEntityById = roleTypeId => {
    window.location.href = `/role-page/edit/${roleTypeId}`;
}