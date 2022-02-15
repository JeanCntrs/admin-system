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
        const element = document.getElementsByClassName('paginate_button current')[0];
        const currentPage = element.innerHTML;
        
        buildTable(() => {
            while (true) {
                const selectorAll = document.querySelectorAll('#table_paginate :not(#table_previous) a');
                let link, bucleIndex
                let found = false

                for (let i = 0; i < selectorAll.length; i++) {
                    link = selectorAll[i];
                    bucleIndex = document.getElementsByClassName('paginate_button current')[0].innerHTML;

                    if (bucleIndex == currentPage) {
                        found = true;

                        break;
                    } else {
                        document.getElementById('table_next').click();
                    }
                }

                if (found) break;
            }
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