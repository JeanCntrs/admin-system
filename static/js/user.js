window.onload = () => {
    buildTable();
    buildSelectRoleType();
    buildSelectPerson();
}

const buildTable = () => {
    const url = '/users/list';
    const tableHeaders = ['User ID', 'Username', 'Fullname', 'Role Type'];
    const fields = ['UserId', 'Username', 'Fullname', 'RoleTypeName'];
    const elementId = 'user_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'UserId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}

const buildSelectRoleType = () => {
    const url = '/role-types/list';
    const value = 'RoleTypeId';
    const name = 'Name';
    const elementId = 'slcRoleType';

    getDataSelect(url, value, name, elementId);
}

const buildSelectPerson = () => {
    const url = '/persons/list/without-user';
    const value = 'PersonId';
    const name = 'Fullname';
    const elementId = 'slcPerson';

    getDataSelect(url, value, name, elementId);
}

const getEntityById = userId => {
    document.getElementById('dv_pwd').style.display = 'none';
    document.getElementById('inp_password').value = 'pwd';
    
    fetch(`/users/list/id/${userId}`)
        .then(response => response.json())
        .then(response => {
            document.getElementById('inp_user_id').value = response.UserId;
            document.getElementById('inp_user_name').value = response.Username;
            //document.getElementById('slcPerson').value = response.PersonId;
            document.getElementById('slcRoleType').value = response.RoleTypeId;
        })
}

const showPwd = () => {
    document.getElementById('dv_pwd').style.display = 'block';
}

const create = () => {
    const userId = document.getElementById('inp_user_id').value;
    const username = document.getElementById('inp_user_name').value;
    const password = document.getElementById('inp_password').value;
    const personId = document.getElementById('slcPerson').value;
    const roleTypeId = document.getElementById('slcRoleType').value;

    if (username.trim().length === 0) {
        alert('Username field is required', '', 'error');
        return;
    }

    if (password.trim().length === 0) {
        alert('Password field is required', '', 'error');
        return;
    }

    if (personId.trim().length === 0) {
        alert('Person field is required', '', 'error');
        return;
    }

    if (roleTypeId.trim().length === 0) {
        alert('Role type field is required', '', 'error');
        return;
    }

    const user = {
        userId: userId == '' ? 0 : parseInt(userId),
        username,
        password,
        personId: parseInt(personId),
        roleTypeId: parseInt(roleTypeId)
    }

    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch('users/create', {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST',
                body: JSON.stringify(user)
            })
                .then(response => response.text())
                .then(response => {
                    if (response != '1') {
                        alert('An error has occurred');

                        return;
                    }

                    document.getElementById('btnCloseModal').click();
                    buildTable();
                    buildSelectPerson();
                    alert();

                    return;
                })
        }
    })
}