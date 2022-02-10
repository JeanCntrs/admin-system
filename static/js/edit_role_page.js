window.onload = () => {
    getEntityById();
}

const getEntityById = () => {
    const roleTypeId = document.getElementById('inp_hdn_role_type_id').value;

    fetch(`/role-types/list/id/${roleTypeId}`)
        .then(response => response.json())
        .then(response => {
            document.getElementById('inp_role_type_id').value = response.RoleTypeId;
            document.getElementById('inp_name_role_type').value = response.Name;
            document.getElementById('inp_description').value = response.Description;
        })

    buildTable();
}

const buildTable = () => {
    const url = '/pages/list';
    const tableHeaders = ['Page ID', 'Message', 'Route'];
    const fields = ['PageId', 'Message', 'Route'];
    const elementId = 'edit_role_page_table';
    const showBtnEdit = false;
    const showBtnDelete = false;
    const propertyName = 'PageId';
    const isPopup = false;
    const isChecked = true;
    const isCallback = true
    const callback = () => {
        getSelectedCheckboxByRoleType()
    }

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName, undefined, isPopup, isChecked, isCallback, callback);
}

const getSelectedCheckboxByRoleType = () => {
    const roleTypeId = document.getElementById('inp_hdn_role_type_id').value;

    fetch(`/role-page/list/id/${roleTypeId}`)
        .then(response => response.json())
        .then(response => {
            for (let index = 0; index < response.length; index++) {
                const data = response[index];
                
                document.getElementById(`chk${data.PageId}`).checked = true;
            }
        })
}

const showAlert = () => {
    const roleTypeId = document.getElementById('inp_role_type_id').value;
    const nameRoleType = document.getElementById('inp_name_role_type').value;
    const description = document.getElementById('inp_description').value;
    const selectedCheckbox = getSelectedCheckbox();
    
    if (nameRoleType.trim().length === 0) {
        alert('Name role type field is required', '', 'error');
        return;
    }

    if (description.trim().length === 0) {
        alert('Description field is required', '', 'error');
        return;
    }

    const roleType = {
        roleTypeId: roleTypeId == '' ? 0 : parseInt(roleTypeId),
        name: nameRoleType,
        description,
        pagesId: selectedCheckbox
    }

    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch('/role-types/create', {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST',
                body: JSON.stringify(roleType)
            })
                .then(response => response.text())
                .then(response => {
                    if (response != '1') {
                        alert('An error has occurred');

                        return;
                    }

                    document.location.href = '/role-page';

                    return;
                })
        }
    })
}