window.onload = () => {
    $("#inp_birthday").datepicker({ dateFormat: 'dd/mm/yy' });
    buildTable();
    buildSelect();
}

const buildTable = () => {
    const url = '/persons/list';
    const tableHeaders = ['Person ID', 'Fullname', 'Name Type Person', 'Birthday'];
    const fields = ['PersonId', 'Fullname', 'NameTypePerson', 'FormattedBirthday'];
    const elementId = 'person_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'PersonId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}

const buildSelect = () => {
    const url = '/persons/list/type';
    const value = 'TypePersonId';
    const name = 'Name';
    const elementId = 'slcTypePerson';

    getDataSelect(url, value, name, elementId);
}

const getPersonsByFullname = () => {
    const fullname = document.getElementById("inp_search_person_fullname").value;

    const url = fullname ? `/persons/list/name/${fullname}` : '/persons/list';
    const tableHeaders = ['Person ID', 'Fullname', 'Name Type Person', 'Birthday'];
    const fields = ['PersonId', 'Fullname', 'NameTypePerson', 'FormattedBirthday'];
    const elementId = 'person_table';
    const showBtnEdit = true;
    const showBtnDelete = true;
    const propertyName = 'PersonId';

    getDataTable(url, tableHeaders, fields, elementId, showBtnEdit, showBtnDelete, propertyName);
}

const getEntityById = personId => {
    fetch(`/persons/list/id/${personId}`)
        .then(response => response.json())
        .then(response => {
            document.getElementById('inp_person_id').value = response.PersonId;
            document.getElementById('inp_person_name').value = response.Name;
            document.getElementById('inp_father_last_name').value = response.FatherLastName;
            document.getElementById('inp_mother_last_name').value = response.MotherLastName;
            document.getElementById('slcTypePerson').value = response.TypePersonId;
            document.getElementById('inp_birthday').value = response.FormattedBirthday;
        })
}

const create = () => {
    const personId = document.getElementById('inp_person_id').value;
    const personName = document.getElementById('inp_person_name').value;
    const fatherLastName = document.getElementById('inp_father_last_name').value;
    const motherLastName = document.getElementById('inp_mother_last_name').value;
    const typePerson = document.getElementById('slcTypePerson').value;
    const birthday = document.getElementById('inp_birthday').value;

    if (personName.trim().length === 0) {
        alert('Name field is required', '', 'error');
        return;
    }

    if (fatherLastName.trim().length === 0) {
        alert('Father last name field is required', '', 'error');
        return;
    }

    if (motherLastName.trim().length === 0) {
        alert('Mother last name field is required', '', 'error');
        return;
    }

    if (typePerson.trim().length === 0) {
        alert('Type person field is required', '', 'error');
        return;
    }

    if (birthday.trim().length === 0) {
        alert('Birthday field is required', '', 'error');
        return;
    }

    const person = {
        personId: personId == '' ? 0 : parseInt(personId),
        name: personName,
        fatherLastName,
        motherLastName,
        typePersonId: parseInt(typePerson),
        birthday: $('#inp_birthday').datepicker('getDate')
    }

    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch('persons/create', {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST',
                body: JSON.stringify(person)
            })
                .then(response => response.text())
                .then(response => {
                    if (response != '1') {
                        alert('An error has occurred');

                        return;
                    }

                    document.getElementById('btnCloseModal').click();
                    buildTable();
                    alert();

                    return;
                })
        }
    })
}

const deleteEntity = (id) => {
    confirmation().then((result) => {
        if (result.isConfirmed) {
            fetch(`persons/delete/${id}`)
                .then(response => response.text())
                .then(response => {
                    if (response != '1') {
                        alert('An error has occurred', '');

                        return;
                    }

                    buildTable();
                    alert('Success', 'Your data has been deleted');

                    return;
                })
        }
    })
}